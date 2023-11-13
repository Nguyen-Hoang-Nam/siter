//go:build darwins || linux
// +build darwins linux

package pty

import (
	"os"
	"os/exec"
	"siter/config"

	"github.com/creack/pty"
)

func StartProcess(c *config.Config) (p PTYProcess, err error) {
	os.Setenv("TERM", "dumb")
	p, err = pty.Start(exec.Command(c.Shell.Command))

	// if c.Shell == "sh" {
	// 	preprocessCommand := "stty erase ^H\r"
	// 	preprocessCommandLen = len(preprocessCommand) + 3
	// 	process.Write([]byte(preprocessCommand))
	// }

	return
}
