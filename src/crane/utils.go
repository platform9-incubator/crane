package crane

import (
  "io/ioutil"
  "log"
  "net"
  "os/user"
  "os"
  "os/exec"
  "fmt"
  "strings"
  "time"
  "gopkg.in/urfave/cli.v2"
  "github.com/docker/containerd/api/grpc/types"
  "github.com/BurntSushi/toml"
  "google.golang.org/grpc"
  "google.golang.org/grpc/grpclog"
)

type exit struct {
  Code int
}

type CraneEnv struct {
    Cmd string
}


// cleanup is extremely unsafe
func cleanup(file_or_dir string) {
    exec.Command("rm", "-rf", file_or_dir).Run()
}

func dir_exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}


func extract_env_cmd(dir_path string) (string, error) {
    var craneEnv CraneEnv
    _, err := toml.DecodeFile(fmt.Sprintf("%s/.crane.env", dir_path), &craneEnv)
    if err != nil {
        return "", err
    }
    return craneEnv.Cmd, err
}


func clone_repo(url string, dest_path string, mkdir bool) (bool, error) {
  if _,err := dir_exists(dest_path); err != nil {
    if ! mkdir {
      return false, err
    }
    usr,err := user.Current()
    if err != nil {
      return false, err
    }
    homeDir := usr.HomeDir
    fileInfo,err := os.Lstat(homeDir)
    if err != nil {
      return false, err
    }
    if err = os.MkdirAll(dest_path, fileInfo.Mode()); err != nil {
      return false, err
    }
  }
  SafeRun(exec.Command("git", "clone", url, dest_path), "Git cloning failed")
  return true, nil
}


func cranetainer_path() (string, error) {
  fmt.Println("Finding cranetainers path")
  var dir string
  usr,err := user.Current()
  if err != nil {
    return "", err
  }
  homeDirs,err := ioutil.ReadDir(usr.HomeDir)
  if err != nil {
    return "", err
  }
  for _,file := range homeDirs {
    if strings.Contains(file.Name(), "cranetainer") {
      dir = usr.HomeDir + "/" + file.Name()
    }
  }
  return dir, nil
}

func errorExit(format string, a ...interface{}) {
  fmt.Printf(format, a)
  cli.OsExiter(2)
}


func GetClient(connTimeout int) types.APIClient {
  // Parse proto://address form addresses.
  // TODO: This is hardcoded right now to be the default containerd address
  bindSpec := "unix:///var/run/docker/libcontainerd/docker-containerd.sock"
  bindParts := strings.SplitN(bindSpec, "://", 2)
  if len(bindParts) != 2 {
    grpclog.Fatal(fmt.Sprintf("bad bind address format %s, expected proto://address", bindSpec), 1)
  }

  // reset the logger for grpc to log to dev/null so that it does not mess with our stdio
  grpclog.SetLogger(log.New(ioutil.Discard, "", log.LstdFlags))
  timeout := time.Duration(connTimeout) * time.Second
  dialOpts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithTimeout(timeout)}
  // Below 'WithDialer' call required in order to connect to Unix socket
  dialOpts = append(dialOpts,
    grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
      return net.DialTimeout(bindParts[0], bindParts[1], timeout)
    },
    ))
  conn, err := grpc.Dial(bindSpec, dialOpts...)
  if err != nil {
    grpclog.Fatal(err.Error(), 1)
  }
  return types.NewAPIClient(conn)
}

func Fatal(err string, code int) {
  fmt.Fprintf(os.Stderr, "[crane] %s\n", err)
  panic(exit{code})
}

func SafeRun(c *exec.Cmd, errMsg string) {
  err := c.Run()
  if err != nil {
    fmt.Println(err.Error())
    Fatal(errMsg, 2)
  }
}
