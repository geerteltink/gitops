package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// Version contains the build version
	Version string = "0.2.1"

	// versionCmd represents the version command
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of GitOps",
		Long:  "All software has versions. This is GitOps's.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("GitOps %s compiled with %v on %v/%v\n", Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
