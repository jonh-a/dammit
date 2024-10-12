/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"go_dammit/pkg"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Analyze the output of a given command",
	Long: `Analyze the output of a given command. To retrieve the output,
you will be prompted to re-run the command.

Example:
go_dammit run "git push"`,
	Run: func(cmd *cobra.Command, args []string) {
		out := pkg.Ask(args[0])
		fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().String("command", "", "Command to troubleshoot")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
