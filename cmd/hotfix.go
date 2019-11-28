package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	cmdHotfix = &cobra.Command{
		RunE:  hotfix,
		Use:   "hotfix",
		Short: "checkout hotfix",
		Long:  "Checkout hotfix from pull request.",
	}
)

func init() {
	cmdHotfix.PersistentFlags().Int("pr", 0, "pull request number")
	cmdHotfix.PersistentFlags().StringP("branch", "b", "master", "base branch for the pr")
	cmdHotfix.MarkFlagRequired("pr")

	rootCmd.AddCommand(cmdHotfix)
}

func hotfix(cmd *cobra.Command, args []string) error {
	var git = &Git{}
	var pr int
	var branch string
	var err error

	pr, err = cmd.Flags().GetInt("pr")
	if err != nil || pr < 1 {
		return errors.New("The --pr argument must be set and must be a valid pr")
	}

	// git checkout -b <branch> [<start point>]
	branch, err = cmd.Flags().GetString("branch")
	_, err = git.Exec("checkout", "-b", fmt.Sprintf("hotfix/%d", pr), branch)
	if err != nil {
		return err
	}

	// git fetch upstream refs/pull/<pr>/head
	_, err = git.Exec("fetch", "upstream", fmt.Sprintf("refs/pull/%d/head", pr))
	if err != nil {
		return err
	}

	// git merge FETCH_HEAD --no-ff -m "chore: merge pull request #<pr>"
	_, err = git.Exec("merge", "FETCH_HEAD", "--no-ff", "-m", fmt.Sprintf("chore: merge pull request (#%d)", pr))
	if err != nil {
		return err
	}

	fmt.Printf("\nRun tests and tools:\n")
	fmt.Printf(" - composer check\n")
	fmt.Printf(" - keep-a-changelog entry:fixed --pr %d \"fixes ...\"\n", pr)
	fmt.Printf("\nWhen ready merge the pull request:\n")
	fmt.Printf(" - gitops merge\n")

	return nil
}
