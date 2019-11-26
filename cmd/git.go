package cmd

import (
	"strings"
)

func gitCmd(args ...string) *Cmd {
	cmd := New("git")

	/*
		for _, v := range GlobalFlags {
			cmd.WithArg(v)
		}
	*/

	for _, a := range args {
		cmd.WithArg(a)
	}

	return cmd
}

func gitRemotes() ([]string, error) {
	remoteCmd := gitCmd("remote", "-v")
	remoteCmd.Stderr = nil
	output, err := remoteCmd.Output()
	return outputLines(output), err
}

func outputLines(output string) []string {
	output = strings.TrimSuffix(output, "\n")
	if output == "" {
		return []string{}
	} else {
		return strings.Split(output, "\n")
	}
}
