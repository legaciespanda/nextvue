/*
Copyright © 2023 Ernest Obot <ernestobot.dev@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

var directory string //directory to scaffold respective framework into
var name string      //project name

func scaffoldApp(appDirectory string, appType string, projectname string) error {
	var command string

	//only run on windows
	if runtime.GOOS == "windows" {
		//check if nodejs is installed
		_, er := exec.LookPath("node")
		if er != nil {
			return fmt.Errorf("Nextvue Error -): %v", "didn't find 'NodeJS' on this machine. Kindly install Nodejs to run\n")
		}

		//change directory to where the app will be scafollded
		home, _ := os.UserHomeDir()
		err := os.Chdir(filepath.Join(home, appDirectory))
		if err != nil {
			//panic(err)
			return fmt.Errorf("Nextvue Error -): %v", err)
		}

		//checks
		if appType == "nextjs" {
			command = "npm create next-app " + projectname + " --yes"
		} else if appType == "nuxtjs" {
			command = "yarn create nuxt-app " + projectname
		} else {
			return fmt.Errorf("Nextvue Error -): Unsupported app. Kindly use nextjs or nuxtjs")
		}

		cmd := exec.Command("cmd", "/C", command)
		// The `Output` method executes the command and
		// collects the output, returning its value
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		errc := cmd.Run()
		if errc != nil {
			return fmt.Errorf("Nextvue Error -): Unable to scaffold "+appType+" app: %v", err)
		}
		// out, err := cmd.Output()
		// if err != nil {
		// 	// if there was any error, print it here
		// 	//fmt.Println("could not run command: ", err)
		// 	return fmt.Errorf("Nextvue Error -): could not run command: %v", err)
		// }

		// // otherwise, print the output from running the command
		// fmt.Println("Nextvue -): ", string(out))
		return nil
	} else {
		return fmt.Errorf("Nextvue Error -):  can only run on windows for now")
	}
}

//command for scaffolding nextjs application
var scaffoldCmd = &cobra.Command{
	Use:     "scaffold",
	Short:   "Scaffold new Next JS/Vue js application",
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// get the flag value, its default value is false
		//fstatus, _ := cmd.Flags().GetBool("float")
		//taskTitle, _ := cmd.Flags().GetString("title")
		fmt.Printf("Kindly take a cofee while your " + args[0] + " app is scaffolding...\n")
		//res := "Scaffolding nextjs" //pkg.scaffoldNext(version, directory)
		errorx := scaffoldApp(directory, args[0], name)
		fmt.Println(errorx)
		if errorx != nil {
			return
		} else {
			fmt.Printf("Congratulations, your " + args[0] + " scaffolded!\n")
		}

	},
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)
	scaffoldCmd.Flags().StringVarP(&name, "name", "n", "", "Project name")
	scaffoldCmd.Flags().StringVarP(&directory, "directory", "d", "", "Directory to scaffold into")

	//ensure that directory and project name is a reqired flag
	if err := scaffoldCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}

	if err := scaffoldCmd.MarkFlagRequired("directory"); err != nil {
		fmt.Println(err)
	}
}
