package cmd

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// upstreamCmd represents the upstream command
var upstreamCmd = &cobra.Command{
	Use:   "upstream",
	Short: "set upstream remote",
	Long: `Setup the upstream remote to the original project location, which needs to be done only once. For example:

gitops upstream <uri-to-original-project>
gitops upstream git@github.com:<original_organization>/<project>.git`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 || args[0] == "" {
			return errors.New("Requires an uri argument")
		}

		// Detect upstream from package.json

		// Get remotes: git remote -v
		rs, err := git.Remotes()
		if err != nil {
			err = fmt.Errorf("Can't load git remote")
			return
		}

		out, err := exec.Command("git", "status", "-v").Output()
		if err != nil {
			return err
		}
		fmt.Printf("output\n", out)

		// Find upstream remote
		// hasUpstream := false

		// if !hasUpstream: git remote add upstream <uri>
		// else: git remote set-url upstream <uri>

		// Track upstream
		// git config branch.master.remote upstream

		uri := args[0]
		fmt.Println(uri)
		exec.Command("sh", "-c", "echo '1 2 3'")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(upstreamCmd)
}
