package crane

import (
  "io/ioutil"
  "os/user"
  "os"
  "os/exec"
  "fmt"
  "github.com/BurntSushi/toml"
  //"path/filepath"
  "strings"
  "gopkg.in/urfave/cli.v2"
)

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
    fileInfo,err := os.Lstat("/Users/josh/cranetainers")
    if err != nil {
      return false, err
    }
    if err = os.MkdirAll(dest_path, fileInfo.Mode()); err != nil {
      return false, err
    }
  }
  cmd := exec.Command("git", "clone", url, dest_path)
  err := cmd.Run()
  if err != nil {
    return false, err
  }
  return true, err
}


func cranetainer_path() (string, error) {
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
