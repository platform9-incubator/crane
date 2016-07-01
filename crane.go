package main

import (
	"os"
	"gopkg.in/urfave/cli.v2" // imports as package "cli"
  "time"
	"crane"
)

type exit struct {
	Code int
}

func main() {
	// We want our defer functions to be run when calling fatal()
	defer func() {
		if e := recover(); e != nil {
			if ex, ok := e.(exit); ok == true {
				os.Exit(ex.Code)
			}
			panic(e)
		}
	}()
  app := &cli.App{
    Version: "v0.0.1",
    Compiled: time.Now(),
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "conn-timeout",
				Value: "60",
			},
		},
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
			crane.InitCommand,
			crane.RunCommand,
			crane.ConvertCommand,
    },
  }
  app.Run(os.Args)
}
