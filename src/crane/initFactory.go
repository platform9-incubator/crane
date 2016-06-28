package crane

import (
	"gopkg.in/urfave/cli.v2" // imports as package "cli"
  "fmt"
  //"github.com/google/go-github/github"
  "os/exec"
  "io/ioutil"
  //"regexp"
  //"strings"
  //"os"
)


func InitCommand() *cli.Command {
  const cmdName = "init"
  command := cli.Command{
    Name:    cmdName,
    Aliases: []string{"i"},
    Usage:   "Initialize a crane-tainer",
    ArgsUsage: "--src <value> container-name",
    // Use Flags if you want that
    Flags: []cli.Flag {
      &cli.StringFlag{
        Name: "src",
        Value: "",
        Usage: "Github repo URL or path to local directory",
      },
      &cli.StringFlag{
        Name: "dest",
        Value: "/crane/<container-name>",
        Usage: "Where to initialize your crane-tainer",
      },
    },
    Action: initMain,
    // I feel like subcommands are a better experience
    //Subcommands: []*cli.Command{
    //  sub_add(),
    //},
  }
  return &command
}



func initMain(c *cli.Context) error {
  src := c.String("src")
  containerName := c.Args().First()
  if src == "" || containerName == "" {
    cli.ShowCommandHelp(c, "init")

  }
  //cmd := exec.Command("ps", "-e", "-o", "command")
  //psResults,err := cmd.Output()
  //r, _ := regexp.Compile("docker-containerd -l \"(.+)\"")
  //socketForCtr := r.FindString(string(psResults))
  //fmt.Println(socketForCtr)

  crane_dir,err := cranetainer_path()
  if err != nil {
    errorExit("Something bad happened while trying to walk the path", "")
  }
  fmt.Println("Found a directory!")
  fmt.Println(crane_dir)
  containers,err := ioutil.ReadDir(crane_dir)
  for _,container := range containers {
    if container.Name() == containerName {
      fmt.Printf("'%s' already exists. Container names must be unique\n", containerName)
      cli.OsExiter(2)
    }
  }

  success,err := clone_repo(src, fmt.Sprintf("%s/%s/%s", crane_dir, containerName, "rootfs"), true)
  if err != nil {
    errorExit("Error cloning repo", nil)
    //fmt.Printf("'%s' already exists. Container names must be unique", containerName)
    //cli.OsExiter(2)
  }
  if success {
    fmt.Println("Amazing!! Cloned successfully!!")
  }


  cmd = exec.Command("docker-runc", "spec")
  err = cmd.Run()
  if err != nil {
    fmt.Println("Can't find docker-runc")
    cli.OsExiter(2)
    return nil
  }

  cmd := exec.Command("mv", "config.json", fmt.Sprintf("%s/%s/%s", crane_dir, containerName, "config.json"))
  err = cmd.Run()
  if err != nil {
    errorExit("Can't move file into %s", containerName)
  }


  fmt.Println(src)
  fmt.Println(containerName)
  return nil
}

func sub_add() *cli.Command {
  command := cli.Command{
    Name:    "add",
    Aliases: []string{"a"},
    Usage:   "Add stuff",
    Action:  func (c *cli.Context) error {
      fmt.Print(c.Args().Get(0))
      fmt.Print(c.Args().Get(1))
      fmt.Print(c.Args().Get(2))
      fmt.Println("Init add subcommand!!")
      return nil
    },
  }
  return &command
}

func Init (c *cli.Context) error {
  fmt.Print("INIT FUNCTION!!!")
	return nil
	//var src string
	//var img_name string
	//var init_cmd string
	//var tag string
    //
	//const PORTS = "8000:8000"
	//ERR := c.Ui.Error
	//INFO := c.Ui.Info
	//var err error
	//foundEnvCmd := false
	//tmp_dir := "/tmp/"
	//tmp_tar := tmp_dir + "init.tar"
	//tmp_src := tmp_dir + "cloned_repo"
    //
    //
	//cmdFlags := flag.NewFlagSet("init", flag.ContinueOnError)
	//cmdFlags.Usage = func() { c.Ui.Output(c.Help())}
	//cmdFlags.StringVar(&src, "src", "", "Source Repo")
	//cmdFlags.StringVar(&img_name, "name", "", "Name for new Docker image")
	//cmdFlags.StringVar(&init_cmd, "init", "", "Command to run once it's in a container")
	//cmdFlags.StringVar(&tag, "tag", "", "Optional tag to run from")
    //
	//err = cmdFlags.Parse(args)
	//if err != nil {
	//	return 1
	//}
    //
	//if src == "" {
	//	ERR(c.Help())
	//	return 1
	//}
	//if img_name == "" {
	//	ERR(c.Help())
	//	return 1
	//}
    //
	//// Source is a github repo
	//if strings.HasPrefix(src, "http") {
	//	// Clone repo
     //   if _,err := clone_repo(src, tmp_src); err != nil {
     //       return 1
     //   }
     //   src = tmp_src
	//}
	//INFO("Success!\n")
    //
    //INFO("Analyzing .crane.env")
    //envCmd,err := extract_env_cmd(src)
    //if err != nil {
     //   ERR("Couldn't find .crane.env in directory")
    //}
    //if foundEnvCmd {
     //   INFO("Success!\nFound: "+envCmd)
    //}
    //
    //
	//INFO("Creating a temporary image")
	//cmd := exec.Command("tar", "-C", tmp_dir, "-cvf", tmp_tar, src)
	//err = cmd.Run()
	//if err != nil {
	//	ERR("Tar of the git repo failed. Removing tmp_tar")
	//	cleanup(tmp_tar)
	//	return 1
	//}
    //
    //
	//INFO("Importing image")
	//cmd = exec.Command("docker", "import", tmp_tar, img_name)
	//err = cmd.Run()
	//if err != nil {
	//	ERR("Can't import docker file")
	//	cleanup(tmp_tar)
	//	return 1
	//}
	// // Clean up tmp_tar anyways
	//cleanup(tmp_tar)
    //
	//if !foundEnvCmd {
	//	INFO("No command found. Image created but not container")
	//	return 1
	//}
    //
	//// We can actually create the container and run it
	//INFO("Creating the container")
	//cmd = exec.Command("docker", "run", "-p", PORTS,
	//	img_name, envCmd)
	//err = cmd.Run()
	//if err != nil {
	//	ERR("Can't run container")
	//	return 1
	//}
	//INFO("Docker container running")
	//return 0
}
