//go:build windows
// +build windows

package pty

import (
	"siter/config"

	"github.com/UserExistsError/conpty"
)

func startProcess(c *config.Config) (p PTYProcess, err error) {
	p, err = conpty.Start(c.Shell.Command)
	if err != nil {
		return p, err
	}

	return
}
