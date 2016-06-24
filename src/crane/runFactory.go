package crane

import (
	"strings"
	"github.com/mitchellh/cli"
	"flag"
	"os/exec"
	"fmt"
	"github.com/BurntSushi/toml"
)

type RunCommand struct {
	Ui cli.Ui
}

type CraneEnv struct {
	Cmd string
}

// The three interface functions

func (c *RunCommand) Run(args []string) int {
	var src string
	var tag string
	var portmap string
	var local string

	tmp_tar := "/tmp/temp.tar"
	dest := "/tmp/test"

	cmdFlags := flag.NewFlagSet("convert", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help())}
	cmdFlags.StringVar(&src, "src", "", "Source Repo")
	cmdFlags.StringVar(&local, "local", "", "Local repo")
	cmdFlags.StringVar(&tag, "tag", "", "Optional tag to run from")
	cmdFlags.StringVar(&portmap, "port", "", "Port mapping")

	if err := cmdFlags.Parse(args); err != nil {
		c.Ui.Error("Can't parse command line")
		return 1
	}

	if src == "" && local == "" {
		c.Ui.Error("Please specify src")
		return 1
	}
	c.Ui.Output("Fetching the git repo")

	if src != "" {
		cmd := exec.Command("git", "clone", src, dest)
		err := cmd.Run()
		if err != nil {
			c.Ui.Error(err.Error())
			return 1
		}
	}

	if local != "" {
		dest = local
	}

	if tag != "" {
		cmd := exec.Command("cd", dest)
		err := cmd.Run()
		if err != nil {
			c.Ui.Error(err.Error())
			return 1
		}

		cmd = exec.Command("git", "checkout", "tags/"+tag, "-b", tag)
		err = cmd.Run()
		if err != nil {
			c.Ui.Error(err.Error())
			return 1
		}
	}

	c.Ui.Output("Analyzing .crane.env")
	var craneEnv CraneEnv
	_, err := toml.DecodeFile(fmt.Sprintf("%s/.crane.env", dest), &craneEnv)
	if err != nil {
		c.Ui.Error(err.Error())
		c.Ui.Error("Can't parse the .crane.env in the git repository")
		return 1
	}

	c.Ui.Output("Creating a temporary image")
	cmd := exec.Command("tar", "-C", dest, "-cvf", tmp_tar, "./")
	err = cmd.Run()
	if err != nil {
		c.Ui.Error("Tar of the git repo failed")
		return 1
	}

	c.Ui.Output("Importing image")

	cmd = exec.Command("docker", "import", tmp_tar, "crane_image")
	err = cmd.Run()
	if err != nil {
		c.Ui.Error("Can't import docker file")
		return 1
	}

	c.Ui.Output("Running the container")
	if portmap != "" {
		cmd = exec.Command("docker", "run", "-p", portmap, "crane_image", craneEnv.Cmd)
	} else {
		cmd = exec.Command("docker", "run", "crane_image", craneEnv.Cmd)
	}
	err = cmd.Run()
	if err != nil {
		c.Ui.Error("Can't import docker file")
		return 1
	}
	c.Ui.Output("Docker container running")
	return 0
}

func (*RunCommand) Help() string {
	helpText := `
	Usage: crane run --src <git-repo> --tag <tag-name>
	`
	return strings.TrimSpace(helpText)
}

func (c *RunCommand) Synopsis() string {
	return "Runs a container from a git repo"
}
