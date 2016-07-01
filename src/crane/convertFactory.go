package crane

import (
	"gopkg.in/urfave/cli.v2" // imports as package "cli"
  "fmt"
)

var ConvertCommand = &cli.Command {
  Name:    "convert",
  Aliases: []string{"c"},
  Usage:   "Convert an existing image to an OCI bundle",
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
  Subcommands: []*cli.Command{
    convertImage,
    convertContainer,
  },
}

var convertImage = &cli.Command {
  Name:    "image",
  Aliases: []string{"i"},
  Usage:   "Convert an existing Docker image to a crane-tainer",
  Action:  func (c *cli.Context) error {
    imageName := c.Args().First()
    if imageName == "" {
      cli.ShowSubcommandHelp(c)
      cli.OsExiter(2)
    }
    fmt.Println(imageName)
    return nil
  },
}

var convertContainer = &cli.Command {
  Name:    "container",
  Aliases: []string{"c"},
  Usage:   "Convert an existing Docker container to a crane-tainer",
  Action:  func (c *cli.Context) error {
    containerName := c.Args().First()
    if containerName == "" {
      cli.ShowSubcommandHelp(c)
      cli.OsExiter(2)
    }
    fmt.Println(containerName)
    return nil
  },
}
