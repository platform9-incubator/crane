package main

import (
	"os"
	"gopkg.in/urfave/cli.v2" // imports as package "cli"
  "crane"
  "time"
)


func main() {
  app := &cli.App{
    Version: "v0.0.1",
    Compiled: time.Now(),
    Authors: []*cli.Author{
      &cli.Author{
        Name:  "Roopak Parikh",
        Email: "rparikh@platform9.com",
      },
      &cli.Author{
        Name:  "Joshua Hurt",
        Email: "josh@platform9.com",
      },
    },
    EnableBashCompletion: true,
    BashComplete: func(c *cli.Context) {
      cli.ShowCompletions(c)
    },
    Commands: []*cli.Command{
      crane.InitCommand(),
      crane.RunCommand(),
      crane.ConvertCommand(),
    },
  }

  app.Run(os.Args)
}
