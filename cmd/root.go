/*
Copyright © 2023 Ernest Obot <ernestobot.dev@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var appVersion = "0.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "nextvue",
	Version: appVersion,
	Short:   "Scafold Nextjs and NuxtJs Application",
	Long:    `NextVue CLI application allows you to easily scaffold Next Js and Nuxt Js javascript framework with s single command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to nextvue")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
