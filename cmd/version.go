package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of GitOps",
	Long:  "All software has versions. This is GitOps's.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("GitOps %s compiled with %v on %v/%v\n", "0.2.0", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
