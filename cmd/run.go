package cmd

import (
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
		pkg.Run(args[0])
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().String("command", "", "Command to troubleshoot")
}
