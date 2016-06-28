package crane

import (
	"gopkg.in/urfave/cli.v2" // imports as package "cli"
  "fmt"
)

func RunCommand() *cli.Command {
  // Flags
  const cmdName = "run"
  command := cli.Command{
    Name:    cmdName,
    Aliases: []string{"r"},
    Usage:   "Run a crane-tainer",
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
      container(),
    },
  }
  return &command
}

func container() *cli.Command {
  command := cli.Command{
    Name:    "container",
    Aliases: []string{"c"},
    Usage:   "Run an existing container",
    Action:  func (c *cli.Context) error {
      src := c.String("src")
      containerName := c.Args().First()
      if containerName == "" {
        cli.ShowSubcommandHelp(c)
        cli.OsExiter(2)
      }
      fmt.Println(src)
      fmt.Println(containerName)
      return nil
    },
  }
  return &command
}
// The three interface functions

func Run (c *cli.Context) error {
	return nil
	//var src string
	//var tag string
	//var portmap string
	//var local string
	//var name string
    //
    //ERR := c.Ui.Error
    //INFO := c.Ui.Info
	//tmp_tar := "/tmp/temp.tar"
	//dest := "/tmp/test"
    //
	//cmdFlags := flag.NewFlagSet("convert", flag.ContinueOnError)
	//cmdFlags.Usage = func() { c.Ui.Output(c.Help())}
	//cmdFlags.StringVar(&src, "src", "", "Source Repo")
	//cmdFlags.StringVar(&local, "local", "", "Local repo")
	//cmdFlags.StringVar(&tag, "tag", "", "Optional tag to run from")
	//cmdFlags.StringVar(&portmap, "port", "", "Port mapping")
	//cmdFlags.StringVar(&name, "name", "", "Name of image and container")
    //
	//if err := cmdFlags.Parse(args); err != nil {
	//	c.Ui.Error("Can't parse command line")
	//	return 1
	//}
    //
	//if src == "" && local == "" {
	//	c.Ui.Error(c.Help())
	//	return 1
	//}
	//if name == "" {
	//	c.Ui.Error(c.Help())
	//	return 1
	//}
    //if portmap == "" {
     //   portmap = "8000:8000"
    //}
	//
	//c.Ui.Output("Cloning repo into "+dest)
    //
	//if src != "" {
     //   if _,err := clone_repo(src, dest); err != nil {
     //       ERR("Cloning failed")
     //       return 1
     //   }
	//}
    //INFO("Success!\n")
    //
	//if local != "" {
	//	dest = local
	//}
    //
	//if tag != "" {
	//	cmd := exec.Command("cd", dest)
	//	err := cmd.Run()
	//	if err != nil {
	//		c.Ui.Error(err.Error())
	//		return 1
	//	}
    //
	//	cmd = exec.Command("git", "checkout", "tags/"+tag, "-b", tag)
	//	err = cmd.Run()
	//	if err != nil {
	//		c.Ui.Error(err.Error())
	//		return 1
	//	}
	//}
    //
	//c.Ui.Output("Analyzing .crane.env to extract cmd...")
    //envCmd,err := extract_env_cmd(dest)
    //if err != nil {
     //   return 1
    //}
    //INFO("Success!\nFound: "+envCmd)
    //
	//c.Ui.Output("Creating a temporary image...")
	//cmd := exec.Command("tar", "-C", dest, "-cvf", tmp_tar, "./")
	//err = cmd.Run()
	//if err != nil {
	//	c.Ui.Error("Tar of the git repo failed")
	//	return 1
	//}
    //INFO("Success!\n")
    //
	//c.Ui.Output("Importing image into docker...")
	//cmd = exec.Command("docker", "import", tmp_tar, name)
	//err = cmd.Run()
	//if err != nil {
	//	c.Ui.Error("Can't import docker file")
	//	return 1
	//}
    //INFO("Success!\n")
    //
	//c.Ui.Output("Running the container...")
	//if portmap != "" {
	//	cmd = exec.Command("docker", "run", "--name", name, "-p", portmap, name, envCmd)
	//} else {
	//	cmd = exec.Command("docker", "run", "--name", name,  name, envCmd)
	//}
	//err = cmd.Run()
	//if err != nil {
	//	c.Ui.Error("Can't run container")
	//	return 1
	//}
    //INFO("Success!\n")
	//return 0
}
