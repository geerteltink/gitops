package cmd

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	cmdMerge = &cobra.Command{
		RunE:  merge,
		Use:   "merge",
		Short: "merge hotfix/feature",
		Long: `
Merge hotfixes or features from pull request.

If the current branch starts with hotfix/, it is merged into master and develop.
If the current branch starts with feature/, it is merged into develop only.
The --branch argument takes a comma seperated list of target branches, which 
overrules the detected branches.
The order of the branches is important. The first one will have 'Close #1' added
to the commit message, others will have 'Forward port #1'.

Usage:

gitops merge
gitops merge --branch master
gitops merge --branch master,develop
`,
	}
)

func init() {
	cmdMerge.PersistentFlags().StringP("branch", "b", "", "base branch for the pr")

	rootCmd.AddCommand(cmdMerge)
}

func merge(cmd *cobra.Command, args []string) error {
	var git = &Git{}
	var branch string
	var err error
	var out string

	// Detect type and pr
	// git rev-parse --abbrev-ref HEAD
	var prBranch string
	prBranch, err = git.Exec("rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return err
	}

	result := strings.Split(prBranch, "/")
	prType := result[0]
	prNumber := result[1]
	log.Println("Detected branch:", prBranch)
	log.Println("Detected pr type:", prType)
	log.Println("Detected pr number:", prNumber)

	if prType != "hotfix" && prType != "feature" {
		return errors.New("The active branch must have a format like '[hotfix|feature]/pr_number'")
	}

	if _, err := strconv.Atoi(prNumber); err != nil {
		return errors.New("The active branch is missing a pr number")
	}

	// Get branches from argument
	branch, _ = cmd.Flags().GetString("branch")
	branches := strings.Split(branch, ",")
	hasDevelopBranch := git.HasDevelopBranch()
	if len(branches) == 0 || branches[0] == "" {
		branches = []string{"master"}
		if prType == "hotfix" && hasDevelopBranch == true {
			branches = []string{"master", "develop"}
		}
		if prType == "feature" && hasDevelopBranch == true {
			branches = []string{"develop"}
		}
	}

	fmt.Println("Target branches:", branches)

	for index, target := range branches {
		fmt.Printf("Merging %s #%s into %s branch\n", prBranch, prNumber, target)

		// git checkout <target_branch>
		out, err = git.Exec("checkout", target)
		if err != nil {
			return err
		}
		if out != "" {
			fmt.Println(out)
		}

		// git merge --no-ff hotfix/1 -m "chore: merge hotfix (#1)" -m "Forward port #1"
		title := fmt.Sprintf("chore: merge %s (#%s) into %s", prType, prNumber, target)
		footer := fmt.Sprintf("Forward port #%s", prNumber)
		if index == 0 {
			footer = fmt.Sprintf("Close #%s", prNumber)
		}
		out, err = git.Exec("merge", "--no-ff", prBranch, "-m", title, "-m", footer)
		if err != nil {
			return err
		}
		if out != "" {
			fmt.Println(out)
		}
	}

	// git checkout master
	if out, err = git.Exec("checkout", "master"); err != nil {
		return err
	}
	if out != "" {
		fmt.Println(out)
	}

	// git branch -D hotfix/1
	if out, err = git.Exec("branch", "-D", prBranch); err != nil {
		return err
	}
	if out != "" {
		fmt.Println(out)
	}

	return nil
}
