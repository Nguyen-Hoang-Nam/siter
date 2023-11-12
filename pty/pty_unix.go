//go:build darwins || linux
// +build darwins linux

package pty

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"siter/config"

	"github.com/creack/pty"
)

type PTYUnix struct {
	process       *os.File
	shellCommand  string
	maxBufferSize int
}

func NewPTYWindows(c config.Config) (IPTY, error) {
	return nil, errors.New("MISS_MATCH_OS")
}

func NewPTYUnix(c config.Config) (p PTYUnix, err error) {
	os.Setenv("TERM", "dumb")
	startCommand := exec.Command(c.Shell)
	process, err := pty.Start(startCommand)
	if err != nil {
		return p, err
	}

	// if c.Shell == "sh" {
	// 	preprocessCommand := "stty erase ^H\r"
	// 	preprocessCommandLen = len(preprocessCommand) + 3
	// 	process.Write([]byte(preprocessCommand))
	// }

	return PTYUnix{process: process, shellCommand: c.Shell, maxBufferSize: c.ScrollbackLines}, nil
}

func (p PTYUnix) Read() io.Reader {
	return p.process
}

func (p PTYUnix) Write(text []byte) (int, error) {
	return p.process.Write(text)
}

func (p PTYUnix) Close() {
	p.process.Close()
}
