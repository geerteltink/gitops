package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Cmd struct {
	Name   string
	Args   []string
	Stdin  *os.File
	Stdout *os.File
	Stderr *os.File
}

func (cmd Cmd) String() string {
	return fmt.Sprintf("%s %s", cmd.Name, strings.Join(cmd.Args, " "))
}

func (cmd *Cmd) WithArg(arg string) *Cmd {
	cmd.Args = append(cmd.Args, arg)

	return cmd
}

func (cmd *Cmd) WithArgs(args ...string) *Cmd {
	for _, arg := range args {
		cmd.WithArg(arg)
	}

	return cmd
}

func (cmd *Cmd) Output() (string, error) {
	// verboseLog(cmd)
	c := exec.Command(cmd.Name, cmd.Args...)
	c.Stderr = cmd.Stderr
	output, err := c.Output()

	return string(output), err
}

func New(name string) *Cmd {
	return &Cmd{
		Name:   name,
		Args:   []string{},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}
