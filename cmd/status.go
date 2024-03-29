package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "recursively check status",
	Long: `Recursively checks the git status in the current and all sub directories. For example:

gitops status`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If .git dir exists we're in a single repo
		if checkDir(".git") {
			displayGitStatus(".")
			return nil
		}

		// Scan all sub directories
		fmt.Print("Scanning sub directories of . \n\n")
		files, err := ioutil.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			dir := file.Name()
			if checkDir(dir) {
				os.Chdir(dir)
				displayGitStatus(dir)
				os.Chdir("..")
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func checkDir(dir string) bool {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return false
	}

	if !fileInfo.IsDir() {
		return false
	}

	return true
}

func displayGitStatus(project string) {
	if !checkDir(".git") {
		// Not a git repository
		return
	}

	branch := getBranch()
	changes := getChanges()
	changedFiles := getChangedFiles()

	error := color.New(color.FgHiRed).SprintFunc()
	success := color.New(color.FgHiGreen).SprintFunc()

	if changes != "" || len(changedFiles) > 0 {
		project = error(project)
	} else {
		project = success(project)
	}

	if changes != "" {
		changes = strings.Replace(changes, "ahead ", "↑", -1)
		changes = strings.Replace(changes, "behind ", "↓", -1)
	}

	changedFilesStatus := ""
	if len(changedFiles) > 0 {
		changedFilesStatus = fmt.Sprintf("[+%d]", len(changedFiles))
	}

	fmt.Printf("%s/%s %s%s\n", project, branch, error(changes), error(changedFilesStatus))
}

func getBranch() string {
	// Get branch name
	cmdName := "git"
	cmdArgs := []string{"rev-parse", "--abbrev-ref", "HEAD"}
	branch, err := exec.Command(cmdName, cmdArgs...).Output()
	if err == nil {
		return strings.TrimSpace(string(branch))
	}

	// Might be a new repo, fallback to status
	cmdName = "git"
	cmdArgs = []string{"status", "-bs"}
	branch, err = exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return "------"
	}

	// Strip spaces and hastags and return the string
	return strings.Trim(strings.TrimSpace(string(branch)), "# ")
}

func getChanges() string {
	cmdName := "git"
	cmdArgs := []string{"for-each-ref", "--format=%(push:track)", "refs/heads"}
	changes, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(changes))
}

func getChangedFiles() []string {
	changedFiles := []string{}
	cmdName := "git"
	cmdArgs := []string{"status", "-s"}
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return changedFiles
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			changedFiles = append(changedFiles, scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return changedFiles
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return changedFiles
	}

	return changedFiles
}
