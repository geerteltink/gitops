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
		RunE:  version,
		Use:   "version",
		Short: "Print the version number of GitOps",
		Long:  "All software has versions. This is GitOps's.",
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func version(cmd *cobra.Command, args []string) error {
	git := &Git{}
	fmt.Printf("GitOps %s compiled with %v on %v/%v\n", Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	gitVersion, err := git.Version()
	if err != nil {
		return err
	}
	fmt.Println(gitVersion)

	return nil
}
