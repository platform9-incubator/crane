package crane

import (
	"gopkg.in/urfave/cli.v2" // imports as package "cli"
  "fmt"
)

func ConvertCommand() *cli.Command {
  const cmdName = "convert"
  command := cli.Command{
    Name:    cmdName,
    Aliases: []string{"c"},
    Usage:   "Convert an existing image to an OCI bundle",
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
      convertImage(),
      convertContainer(),
    },
  }
  return &command
}

func convertImage() *cli.Command {
  command := cli.Command{
    Name:    "image",
    Aliases: []string{"i"},
    Usage:   "Convert an existing Docker image to a crane-tainer",
    Action:  func (c *cli.Context) error {
      imageName := c.Args().First()
      if imageName == "" {
        cli.ShowSubcommandHelp(c)
        cli.OsExiter(2)
      }
      fmt.Println(imageName)
      return nil
    },
  }
  return &command
}

func convertContainer() *cli.Command {
  command := cli.Command{
    Name:    "container",
    Aliases: []string{"c"},
    Usage:   "Convert an existing Docker container to a crane-tainer",
    Action:  func (c *cli.Context) error {
      containerName := c.Args().First()
      if containerName == "" {
        cli.ShowSubcommandHelp(c)
        cli.OsExiter(2)
      }
      fmt.Println(containerName)
      return nil
    },
  }
  return &command
}

func Convert (c *cli.Context) error {
	return nil
    //// Flags
	//var container string
	//var localDir string
    //var repo bool
    //
    //
    //ERR := c.Ui.Error
    //INFO := c.Ui.Info
    //const CRANE_INIT_MSG = "Repo instantiated by crane"
	//tmp_dir := "/tmp/"
	//tmp_tar := tmp_dir + "init.tar"
    //crane_unpack_dir := tmp_dir + "crane"
    //
	//cmdFlags := flag.NewFlagSet("convert", flag.ContinueOnError)
	//cmdFlags.Usage = func() { c.Ui.Output(c.Help())}
	//cmdFlags.StringVar(&container, "container", "", "Source file")
	//cmdFlags.StringVar(&localDir, "local", "", "Destination on disk")
    //cmdFlags.BoolVar(&repo, "repo", false, "Create a new Github Repo for this container's contents")
    //
	//if err := cmdFlags.Parse(args); err != nil {
	//	return 1
	//}
    //
	//if container == "" && (repo == false || localDir == "") {
	//	ERR(c.Help())
	//	return 1
	//}
    //
	//// It's not a Github URL so make sure the local path is okay
	//if ! repo {
     //   INFO("am I making it in here?")
     //   found, err := dir_exists(localDir)
     //   if !found && err == nil {
     //       cmd := exec.Command("mkdir", "-p", localDir, crane_unpack_dir)
     //       err = cmd.Run()
     //       if err != nil {
     //           ERR("I tried")
     //           //cleanup(tmp_tar)
     //           return 1
     //       }
     //   }
    //}
    //// We're having to instantiate a GitHub repo now
    //localDir = tmp_tar
    //INFO(localDir)
    //INFO("Moving tar'd container files to directory: "+ localDir)
    //INFO("This may take a while...")
    //cmd := exec.Command("docker", "export", "-o", localDir, container)
    //err := cmd.Run()
    //if err != nil {
     //   cleanup(localDir)
     //   ERR("Failed to move container files to directory")
     //   return 1
    //}
    //
    //
    //INFO("Exploding container into directory "+ crane_unpack_dir)
    //cmd = exec.Command("tar", "-xf", localDir, "-C", crane_unpack_dir)
    //err = cmd.Run()
    //if err != nil {
     //   cleanup(localDir)
     //   ERR("Something happened while untarring the contents of the container")
     //   return 1
    //}
    //INFO("Success!\n")
    //
    //cmd = exec.Command("cd", crane_unpack_dir)
    //err = cmd.Run()
    //if err != nil {
     //   ERR("Failed trying to CD into "+crane_unpack_dir)
     //   cleanup(localDir)
     //   cleanup(crane_unpack_dir)
     //   return 1
    //}
    //
    //
    //INFO("Creating git repo")
    //cmd = exec.Command("git", "init", crane_unpack_dir)
    //err = cmd.Run()
    //if err != nil {
     //   ERR("Failed trying to initialize git repo")
     //   //cleanup(localDir)
     //   //cleanup(crane_unpack_dir)
     //   return 1
    //}
    //cmd = exec.Command("export", "GIT_DIR="+crane_unpack_dir+"/.git", "GIT_WORK_TREE="+crane_unpack_dir)
    //err = cmd.Run()
    //if err != nil {
     //   ERR("Attempting to set Git environment variables")
     //   //cleanup(localDir)
     //   //cleanup(crane_unpack_dir)
     //   return 1
    //}
    //cmd = exec.Command("git", "commit", "-am", CRANE_INIT_MSG)
    //err = cmd.Run()
    //if err != nil {
     //   ERR("Failed trying to initialize git repo")
     //   //cleanup(localDir)
     //   //cleanup(crane_unpack_dir)
     //   return 1
    //}
    //
    //var userName, repoName string
    //fmt.Print("Enter github username: ")
    //fmt.Scan(&userName)
    //fmt.Print("Enter desired repo name: ")
    //fmt.Scan(&repoName)
    //
    //cmd = exec.Command("git", "remote", "add", "origin", "https://github.com/"+userName+"/"+repoName+".git")
    //_,err = cmd.StdinPipe()
    //cmd.Stdout = os.Stdout
    //cmd.Stderr = os.Stderr
    //if err := cmd.Start(); err != nil {
     //   ERR("Failed on start")
     //   return 1
    //}
    //if err := cmd.Wait(); err != nil {
     //   ERR("Failed on start")
     //   return 1
    //}
    //if err != nil {
     //   cleanup(localDir)
     //   return 1
    //}
    //
    //
    //cmd = exec.Command("git", "push", "origin", "master")
    //err = cmd.Run()
    //if err != nil {
     //   ERR("Failed trying to push the repo")
     //   //cleanup(localDir)
     //   //cleanup(crane_unpack_dir)
     //   return 1
    //}
    //INFO("Success!\n")
    //
	//INFO("Thanks done")
	//return 0
}
