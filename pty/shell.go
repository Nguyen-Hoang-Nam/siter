package pty

import (
	"errors"
	"os"
	"runtime"
	"siter/config"

	"fyne.io/fyne/v2"
)

var (
	ERROR_UNSUPPORTED_OS  = errors.New("UNSUPPORTED_OS")
	ERROR_SHELL_NOT_FOUND = errors.New("SHELL_NOT_FOUND")
)

func GetShell(c config.Config) (process IPTY, err error) {
	goos := runtime.GOOS
	if goos == "windows" {
		return getShellWindow(c)
	} else if goos == "darwin" {
		return getShellLinux(c)
	} else if goos == "linux" {
		return getShellLinux(c)
	} else {
		return process, ERROR_UNSUPPORTED_OS
	}
}

func getShellWindow(c config.Config) (process IPTY, err error) {
	if c.Shell == "." {
		c.Shell = "cmd"
		if c.Shell == "" {
			return process, ERROR_SHELL_NOT_FOUND
		}
	}

	return startWindow(c), nil
}

func startWindow(c config.Config) (process IPTY) {
	process, err := NewPTYWindows(c)
	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	return
}

func getShellLinux(c config.Config) (process IPTY, err error) {
	if c.Shell == "." {
		c.Shell = os.Getenv("SHELL")
		if c.Shell == "" {
			return process, ERROR_SHELL_NOT_FOUND
		}
	}

	return startLinux(c), nil
}

func startLinux(c config.Config) (process IPTY) {
	process, err := NewPTYUnix(c)
	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	return
}
