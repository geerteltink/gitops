package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// upstreamCmd represents the upstream command
var upstreamCmd = &cobra.Command{
	Use:   "upstream",
	Short: "set upstream remote",
	Long: `
Set the upstream remote to the original project location, which needs to be done only once. For example:

gitops upstream <uri-to-original-project>
gitops upstream git@github.com:<original_organization>/<project>.git`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var git = &Git{}

		if len(args) == 0 || args[0] == "" {
			return errors.New("Requires an uri argument")
		}

		// TODO: Detect upstream from package.json or composer.json

		uri := args[0]

		// Get remotes: git remote -v
		remotes, err := git.Remotes()
		if err != nil {
			return err
		}

		// Find upstream remote
		hasUpstream := false
		for _, remote := range remotes {
			if strings.HasPrefix(remote, "upstream") {
				hasUpstream = true
			}
		}

		command := "add"
		if hasUpstream {
			command = "set-url"
			fmt.Println("Setting upstream remote:", uri)
		} else {
			fmt.Println("Adding upstream remote:", uri)
		}

		_, err = git.Exec("remote", command, "upstream", uri)
		if err != nil {
			return err
		}

		// Track upstream
		// git config branch.master.remote upstream
		_, err = git.Exec("config", "branch.master.remote", "upstream")
		if err != nil {
			return err
		}
		fmt.Println("Tracking upstream remote")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(upstreamCmd)
}
