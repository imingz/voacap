package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var GitVersion = "" // gitVersion 是语义化的版本号.

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of voacap",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(GitVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
