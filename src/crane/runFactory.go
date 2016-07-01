package crane

import (
	"gopkg.in/urfave/cli.v2" // imports as package "cli"
  "fmt"
  netcontext "golang.org/x/net/context"
  "github.com/docker/containerd/api/grpc/types"
)

var RunCommand = &cli.Command {
  Name:    "run",
  Aliases: []string{"r"},
  Usage:   "Run a crane-tainer",
  ArgsUsage: "container-name",
  // Use Flags if you want that
  Flags: []cli.Flag {
    &cli.StringFlag{
      Name: "src",
      Value: "",
      Usage: "Github repo URL or path to local directory",
    },
  },
  // I feel like subcommands are a better experience
  Subcommands: []*cli.Command {
    runContainer,
    runStateCheck,
  },
}

var runContainer = &cli.Command {
  Name:    "container",
  Aliases: []string{"c"},
  Usage:   "Run an existing container",
  Action:  func (c *cli.Context) error {
    src := c.String("src")
    containerName := c.Args().First()
    if containerName == "" {
      cli.ShowSubcommandHelp(c)
      cli.OsExiter(2)
    }
    fmt.Println(src)
    fmt.Println(containerName)
    return nil
  },
}

var runStateCheck = &cli.Command {
  Name:    "state",
  Aliases: []string{"s"},
  Usage:   `Specify a container whose state you wish to examine.
            Shows state of all containers by default`,
  Action:  func (c *cli.Context) error {
    client := GetClient(60)
    resp,err := client.State(netcontext.Background(), &types.StateRequest{
      Id: c.Args().First(),
    })
    if err != nil {
      fmt.Println(err.Error())
      return err
    }
    fmt.Println(resp.String())
    return nil
  },
}
