//go:build darwins || linux
// +build darwins linux

package pty

import (
	"os"
	"os/exec"
	"siter/config"

	"github.com/creack/pty"
)

func startProcess(c *config.Config) (p PTYProcess, err error) {
	os.Setenv("TERM", "dumb")

	command := c.Shell.Command
	args := []string{"-c", "stty erase ^H; " + command}

	p, err = pty.Start(exec.Command(command, args...))

	return
}
