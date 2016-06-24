package crane

import (
	"strings"
	"github.com/mitchellh/cli"
	"flag"
    "os/exec"
)

type ConvertCommand struct {
	Ui cli.Ui
}
// The three interface functions

func (c *ConvertCommand) Run(args []string) int {
    // Flags
	var container string
	var dest string


    ERR := c.Ui.Error
    INFO := c.Ui.Info
	tmp_dir := "/tmp/"
	tmp_tar := tmp_dir + "init.tar"

	cmdFlags := flag.NewFlagSet("convert", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help())}
	cmdFlags.StringVar(&container, "container", "", "Source file")
	cmdFlags.StringVar(&dest, "dest", "", "Destination repo")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if container == "" {
		ERR("Please specify src")
		return 1
	}

	if dest == "" {
		ERR("Please specify src")
		return 1
	}

	// It's not a Github URL and it isn't a valid path
	if !strings.HasPrefix(dest, "http") {
        INFO("am I making it in here?")
        found,err := dir_exists(dest)
        if !found && err == nil {
            cmd := exec.Command("mkdir", "-p", dest)
            err = cmd.Run()
            if err != nil {
                ERR("I tried")
                cleanup(tmp_tar)
                return 1
            }
        }
	}

	INFO("Exporting image")
	cmd := exec.Command("docker", "export", container, "--output=", "\\'"+dest+"\\'")
	err := cmd.Run()
	if err != nil {
		cleanup(dest)
        output,_ := cmd.CombinedOutput()
		ERR(string(output))
		return 1
	}

	INFO("Thanks done")
	return 0
}


func (*ConvertCommand) Help() string {
	helpText := `
	Usage: crane convert --src DOCKER-CONTAINER-NAME --dest DEST-GIT
	`
	return strings.TrimSpace(helpText)
}

func (c *ConvertCommand) Synopsis() string {
	return "convert docker image to crane-repo"
}