package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Git represents a collection of git related functions.
type Git struct{}

// Command returns a git command.
func (g *Git) Command(args ...string) *exec.Cmd {
	cmd := exec.Command("git", args...)
	log.Println("exec", cmd.Args)
	return cmd
}

// Exec returns the result of a git command.
func (g *Git) Exec(args ...string) (string, error) {
	cmd := g.Command(args...)
	out, err := cmd.Output()
	return strings.TrimSpace(string(out)), err
}

// Version returns the current git version.
func (g *Git) Version() (string, error) {
	out, err := g.Exec("version")
	if err != nil {
		return "", fmt.Errorf("error running git version: %s", err)
	}
	return firstLine(out), nil
}

// Remotes returns all git remotes.
func (g *Git) Remotes() ([]string, error) {
	out, err := g.Exec("remote", "-v")
	if err != nil {
		return nil, fmt.Errorf("error getting git remotes: %s", err)
	}
	return outputLines(out), err
}

// HasDevelopBranch returns whether the upstream develop branch exists.
func (g *Git) HasDevelopBranch() bool {
	// git ls-remote --exit-code --heads upstream develop
	out, err := g.Exec("ls-remote", "--exit-code", "--heads", "upstream", "develop")
	if err != nil || out == "" {
		return false
	}

	return true
}

func firstLine(output string) string {
	if i := strings.Index(output, "\n"); i >= 0 {
		return output[0:i]
	}

	return output
}

func outputLines(output string) []string {
	output = strings.TrimSuffix(output, "\n")
	if output == "" {
		return []string{}
	}

	return strings.Split(output, "\n")
}
