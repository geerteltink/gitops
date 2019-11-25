package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of GitOps",
	Long:  "All software has versions. This is GitOps's.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GitOps v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
