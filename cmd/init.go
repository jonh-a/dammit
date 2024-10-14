package cmd

import (
	"github.com/spf13/cobra"
	"go_dammit/pkg"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create `dammit` shell alias",
	Long:  "Create `dammit` shell alias that executes go_dammit with a single command.",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Init()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
