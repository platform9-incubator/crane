package crane

import (
	"strings"
	"github.com/mitchellh/cli"
	"os/exec"
	"flag"
	"fmt"
)

type InitCommand struct {
	Ui cli.Ui
}

// The three interface functions

func (c *InitCommand) Run(args []string) int {
	var src string
	var img_name string
	var init_cmd string
	var tag string

	const PORTS = "8000:8000"
	ERR := c.Ui.Error
	INFO := c.Ui.Info
	var err error
	foundEnvCmd := false
	tmp_dir := "/tmp/"
	tmp_tar := tmp_dir + "init.tar"
	tmp_src := tmp_dir + "cloned_repo"


	cmdFlags := flag.NewFlagSet("init", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help())}
	cmdFlags.StringVar(&src, "src", "", "Source Repo")
	cmdFlags.StringVar(&img_name, "name", "", "Name for new Docker image")
	cmdFlags.StringVar(&init_cmd, "init", "", "Command to run once it's in a container")
	cmdFlags.StringVar(&tag, "tag", "", "Optional tag to run from")

	fmt.Println(args)

	err = cmdFlags.Parse(args)
	if err != nil {
		return 1
	}

	if src == "" {
		ERR("Please specify src")
		return 1
	}
	if img_name == "" {
		ERR("Please specify an image name")
		return 1
	}

	// Source is a github repo
	if strings.HasPrefix(src, "http") {
		// Clone repo
        if _,err := clone_repo(src, tmp_src); err != nil {
            return 1
        }
        src = tmp_src
	}

    INFO("Analyzing .crane.env")
    envCmd,err := extract_env_cmd(src)
    if err != nil {
        ERR("Couldn't find .crane.env in directory")
        return 1
    }


	INFO("Creating a temporary image")
	cmd := exec.Command("tar", "-C", tmp_dir, "-cvf", tmp_tar, src)
	err = cmd.Run()
	if err != nil {
		ERR("Tar of the git repo failed. Removing tmp_tar")
		cleanup(tmp_tar)
		return 1
	}


	INFO("Importing image")
	cmd = exec.Command("docker", "import", tmp_tar, img_name)
	err = cmd.Run()
	if err != nil {
		ERR("Can't import docker file")
		cleanup(tmp_tar)
		return 1
	}
	 // Clean up tmp_tar anyways
	cleanup(tmp_tar)

	if !foundEnvCmd {
		INFO("No command found. Image created but not container")
		return 1
	}

	// We can actually create the container and run it
	INFO("Creating the container")
	cmd = exec.Command("docker", "run", "-p", PORTS,
		img_name, envCmd)
	err = cmd.Run()
	if err != nil {
		ERR("Can't run container")
		return 1
	}
	INFO("Docker container running")
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

