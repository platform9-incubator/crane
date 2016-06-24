package crane

import (
	"strings"
	"github.com/mitchellh/cli"
	"flag"
)

type InitCommand struct {
	Ui cli.Ui
}
// The three interface functions

func (c *InitCommand) Run(args []string) int {
	var src string
	var tag string

	cmdFlags := flag.NewFlagSet("init", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help())}
	cmdFlags.StringVar(&src, "src", "", "Source Repo")
	cmdFlags.StringVar(&tag, "tag", "", "Optional tag to run from")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if src == "" {
		c.Ui.Error("Please specify src")
		return 1
	}
	c.Ui.Output("Fetching the git repo")

	c.Ui.Output("Running")
	return 0
}

func (*InitCommand) Help() string {
	helpText := `
	Usage: crane init --src <git-repo> --tag <tag-name>
	`
	return strings.TrimSpace(helpText)
}

func (c *InitCommand) Synopsis() string {
	return "Initializes your repo and container"
}

