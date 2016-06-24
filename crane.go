package main

import (
	"log"
	"os"
	"github.com/mitchellh/cli"
	"crane"
)

func main() {

	c := cli.NewCLI("crane", "0.1.0")
	ui := &cli.BasicUi{Writer: os.Stdout}
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"convert" : func() (cli.Command, error) {
			return &crane.ConvertCommand{
				Ui: ui,
			}, nil
		},
		"init" : func() (cli.Command, error) {
			return &crane.InitCommand{
				Ui: ui,
			}, nil
		},
		"run" : func() (cli.Command, error) {
			return &crane.RunCommand{
				Ui: ui,
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
