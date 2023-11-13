//go:build windows
// +build windows

package pty

import (
	"siter/config"

	"github.com/UserExistsError/conpty"
)

func StartProcess(c *config.Config) (p PTYWindows, err error) {
	p, err := conpty.Start(c.Shell.Command)
	if err != nil {
		return p, err
	}

	return
}
