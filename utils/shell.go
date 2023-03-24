package utils

import (
	"errors"
	"os"
	"runtime"

	"fyne.io/fyne/v2"
)

var (
	ERROR_UNSUPPORTED_OS  = errors.New("UNSUPPORTED_OS")
	ERROR_SHELL_NOT_FOUND = errors.New("SHELL_NOT_FOUND")
)

func GetShell(shellCommand string) (process IPTY, err error) {
	goos := runtime.GOOS
	if goos == "windows" {
		return getShellWindow(shellCommand)
	} else if goos == "darwin" {
		return getShellLinux(shellCommand)
	} else if goos == "linux" {
		return getShellLinux(shellCommand)
	} else {
		return process, ERROR_UNSUPPORTED_OS
	}
}

func getShellWindow(shellCommand string) (process IPTY, err error) {
	if shellCommand == "." {
		shellCommand = os.Getenv("SHELL")
		if shellCommand == "" {
			return process, ERROR_SHELL_NOT_FOUND
		}
	}

	return startWindow(shellCommand), nil
}

func startWindow(shellCommand string) (process IPTY) {
	process, err := NewPTYWindows(shellCommand)
	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	return
}

func getShellLinux(shellCommand string) (process IPTY, err error) {
	if shellCommand == "." {
		shellCommand = os.Getenv("SHELL")
		if shellCommand == "" {
			return process, ERROR_SHELL_NOT_FOUND
		}
	}

	return startLinux(shellCommand), nil
}

func startLinux(shellCommand string) (process IPTY) {
	process, err := NewPTYUnix(shellCommand)
	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	return
}
