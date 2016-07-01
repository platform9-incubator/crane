package crane

import (
	"gopkg.in/urfave/cli.v2" // imports as package "cli"
  "fmt"
  "os/exec"
  "io/ioutil"
  netcontext "golang.org/x/net/context"
  "github.com/docker/containerd/api/grpc/types"
)



var InitCommand = &cli.Command{
  Name:    "init",
  Aliases: []string{"i"},
  Usage:   "Initialize a crane-tainer",
  ArgsUsage: "--src <value> container-name",
  // Use Flags if you want that
  Flags: []cli.Flag {
    &cli.StringFlag{
      Name: "src",
      Value: "",
      Usage: "Github repo URL or path to local directory",
    },
    &cli.StringFlag{
      Name: "dest",
      Value: "/crane/<container-name>",
      Usage: "Where to initialize your crane-tainer",
    },
  },
  Action: initMain,
  // I feel like subcommands are a better experience
  //Subcommands: []*cli.Command{
  //  sub_add(),
  //},

}



func initMain(c *cli.Context) error {
  src := c.String("src")
  containerName := c.Args().First()
  if src == "" || containerName == "" {
    cli.ShowCommandHelp(c, "init")

  }
  client := GetClient(60)
  resp, err := client.GetServerVersion(netcontext.Background(), &types.GetServerVersionRequest{})
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println(resp.Major)
  crane_dir,err := cranetainer_path()
  if err != nil {
    errorExit("Something bad happened while trying to walk the path", "")
  }
  fmt.Println("Found a directory!")
  fmt.Println(crane_dir)
  return nil

  containers,err := ioutil.ReadDir(crane_dir)
  for _,container := range containers {
    if container.Name() == containerName {
      fmt.Printf("'%s' already exists. Container names must be unique\n", containerName)
      cli.OsExiter(2)
    }
  }

  success,err := clone_repo(src, fmt.Sprintf("%s/%s/%s", crane_dir, containerName, "rootfs"), true)
  if err != nil {
    errorExit("Error cloning repo", nil)
    //fmt.Printf("'%s' already exists. Container names must be unique", containerName)
    //cli.OsExiter(2)
  }
  if success {
    fmt.Println("Amazing!! Cloned successfully!!")
  }

  cmd := exec.Command("docker-runc", "spec")
  err = cmd.Run()
  if err != nil {
    fmt.Println("Can't find docker-runc")
    cli.OsExiter(2)
    return nil
  }

  cmd = exec.Command("mv", "config.json", fmt.Sprintf("%s/%s/%s", crane_dir, containerName, "config.json"))
  err = cmd.Run()
  if err != nil {
    errorExit("Can't move file into %s", containerName)
  }

  fmt.Println(src)
  fmt.Println(containerName)
  return nil
}

var sub_add = cli.Command {
  Name:    "add",
  Aliases: []string{"a"},
  Usage:   "Add stuff",
  Action:  func (c *cli.Context) error {
    fmt.Print(c.Args().Get(0))
    fmt.Print(c.Args().Get(1))
    fmt.Print(c.Args().Get(2))
    fmt.Println("Init add subcommand!!")
    return nil
  },
}
