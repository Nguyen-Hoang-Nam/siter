//go:build darwins || linux
// +build darwins linux

package utils

import (
	"bufio"
	"errors"
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

const MaxBufferSize = 16

type PTYUnix struct {
	process      *os.File
	shellCommand string
}

func NewPTYWindows(shellCommand string) (IPTY, error) {
	return nil, errors.New("MISS_MATCH_OS")
}

func NewPTYUnix(shellCommand string) (p PTYUnix, err error) {
	os.Setenv("TERM", "dumb")
	startCommand := exec.Command("/bin/bash")
	process, err := pty.Start(startCommand)
	if err != nil {
		return p, err
	}

	return PTYUnix{process: process, shellCommand: shellCommand}, nil
}

func (p PTYUnix) Read(buffer [][]rune) {
	reader := bufio.NewReader(p.process)

	go func() {
		line := []rune{}
		buffer = append(buffer, line)
		for {
			r, _, err := reader.ReadRune()
			if err != nil {
				if err == io.EOF {
					return
				}
				os.Exit(0)
			}

			line = append(line, r)
			buffer[len(buffer)-1] = line
			if r == '\n' {
				if len(buffer) > MaxBufferSize {
					buffer = buffer[1:]
				}

				line = []rune{}
				buffer = append(buffer, line)
			}
		}
	}()
}

func (p PTYUnix) Write(text []byte) (int, error) {
	return p.process.Write(text)
}

func (p PTYUnix) Close() {
	p.process.Close()
}
