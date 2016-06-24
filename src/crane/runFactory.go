package crane

import (
	"strings"
	"github.com/mitchellh/cli"
	"flag"
	"os/exec"
)

type RunCommand struct {
	Ui cli.Ui
}


// The three interface functions

func (c *RunCommand) Run(args []string) int {
	var src string
	var tag string
	var portmap string
	var local string
	var name string

	tmp_tar := "/tmp/temp.tar"
	dest := "/tmp/test"

	cmdFlags := flag.NewFlagSet("convert", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help())}
	cmdFlags.StringVar(&src, "src", "", "Source Repo")
	cmdFlags.StringVar(&local, "local", "", "Local repo")
	cmdFlags.StringVar(&tag, "tag", "", "Optional tag to run from")
	cmdFlags.StringVar(&portmap, "port", "", "Port mapping")
	cmdFlags.StringVar(&name, "name", "", "Name of image and container")

	if err := cmdFlags.Parse(args); err != nil {
		c.Ui.Error("Can't parse command line")
		return 1
	}

	if src == "" && local == "" {
		c.Ui.Error("Please specify src or local")
		return 1
	}
	if name == "" {
		c.Ui.Error("Please specify a name")
		return 1
	}
	
	c.Ui.Output("Fetching the git repo")

	if src != "" {
        if _,err := clone_repo(src, dest); err != nil {
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
    envCmd,err := extract_env_cmd(dest)
    if err != nil {
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

	cmd = exec.Command("docker", "import", tmp_tar, name)
	err = cmd.Run()
	if err != nil {
		c.Ui.Error("Can't import docker file")
		return 1
	}

	c.Ui.Output("Running the container")
	if portmap != "" {
		cmd = exec.Command("docker", "run", "--name", name, "-p", portmap, name, envCmd)
	} else {
		cmd = exec.Command("docker", "run", "--name", name,  name, envCmd)
	}
	err = cmd.Run()
	if err != nil {
		c.Ui.Error("Can't run container")
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
