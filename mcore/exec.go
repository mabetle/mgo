package mcore

import (
	"os/exec"
)

// NewCommand return new exec cmd
// args[0] cmd name
// args[1:] cmd args
func NewCommand(args []string) *exec.Cmd {
	n := len(args)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return exec.Command(args[0])
	}
	return exec.Command(args[0], args[1:]...)
}

// NewCommandFromString create
func NewCommandFromString(cmd string) *exec.Cmd {
	return NewCommand(ParseStringToArgs(cmd))
}
