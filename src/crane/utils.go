package crane

import (
    "os"
    "os/exec"
    "fmt"
    "github.com/BurntSushi/toml"
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

func clone_repo(url string, dest_path string) (bool, error) {
    if _,err := dir_exists(dest_path); err != nil {
        return false, err
    }
    cmd := exec.Command("git", "clone", url, dest_path)
    err := cmd.Run()
    if err != nil {
        return false, err
    }
    return true, err
}

