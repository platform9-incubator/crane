package crane

import (
	"strings"
	"github.com/mitchellh/cli"
	"flag"
	"fmt"
)

type ConvertCommand struct {
	Ui cli.Ui
}
// The three interface functions

func (c *ConvertCommand) Run(args []string) int {
	var src string
	var dest string

	cmdFlags := flag.NewFlagSet("convert", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help())}
	cmdFlags.StringVar(&src, "src", "", "Source file")
	cmdFlags.StringVar(&dest, "dest", "", "Destination repo")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if src == "" {
		c.Ui.Error("Please specify src")
		return 1
	}

	if dest == "" {
		c.Ui.Error("Please specify src")
		return 1
	}

	c.Ui.Output(fmt.Sprintf("Thanks Done"))
	return 0
}

func (*ConvertCommand) Help() string {
	helpText := `
	Usage: crane convert --src DOCKER-IMAGE-NAME --dest DEST-GIT
	`
	return strings.TrimSpace(helpText)
}

func (c *ConvertCommand) Synopsis() string {
	return "convert docker image to crane-repo"
}