package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	cmdFeature = &cobra.Command{
		RunE:  feature,
		Use:   "feature",
		Short: "checkout feature",
		Long:  "Checkout feature from pull request.",
	}
)

func init() {
	cmdFeature.PersistentFlags().Int("pr", 0, "pull request number")
	cmdFeature.PersistentFlags().StringP("branch", "b", "develop", "base branch for the pr")
	cmdFeature.MarkFlagRequired("pr")

	rootCmd.AddCommand(cmdFeature)
}

func feature(cmd *cobra.Command, args []string) error {
	var git = &Git{}
	var pr int
	var branch string
	var err error

	pr, err = cmd.Flags().GetInt("pr")
	if err != nil || pr < 1 {
		return errors.New("The --pr argument must be set and must be a valid pr")
	}

	branch, err = cmd.Flags().GetString("branch")
	// Check if develop branch exists
	if branch == "develop" && git.HasDevelopBranch() == false {
		fmt.Println("Develop branch not found, using master branch as a base")
		branch = "master"
	}

	// git checkout -b <branch> [<start point>]
	_, err = git.Exec("checkout", "-b", fmt.Sprintf("feature/%d", pr), branch)
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
	fmt.Printf(" - keep-a-changelog entry:added --pr %d \"adds ...\"\n", pr)
	fmt.Printf(" - keep-a-changelog entry:changed --pr %d \"changes ...\"\n", pr)
	fmt.Printf(" - keep-a-changelog entry:deprecated --pr %d \"deprecated ...\"\n", pr)
	fmt.Printf(" - keep-a-changelog entry:removed --pr %d \"removes ...\"\n", pr)
	fmt.Printf("\nWhen ready merge the pull request:\n")
	fmt.Printf(" - gitops merge\n")

	return nil
}
