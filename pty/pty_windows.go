//go:build windows
// +build windows

package pty

import (
	"errors"
	"os"

	"siter/config"

	"github.com/UserExistsError/conpty"
)

type PTYWindows struct {
	cpty          *conpty.ConPty
	shellCommand  string
	maxBufferSize int
}

func NewPTYUnix(c config.Config) (IPTY, error) {
	return nil, errors.New("MISS_MATCH_OS")
}

func NewPTYWindows(c config.Config) (p PTYWindows, err error) {
	cpty, err := conpty.Start(c.Shell)
	if err != nil {
		return p, err
	}

	return PTYWindows{cpty: cpty, shellCommand: c.Shell, maxBufferSize: c.ScrollbackLines}, nil
}

func (p PTYWindows) Read(buffer *[][]rune) {
	go func() {
		line := []byte{}
		// *buffer = append(*buffer, line)

		for {
			n, err := p.cpty.Read(line)
			if err != nil {
				// if err == io.EOF {
				// 	return
				// }
				os.Exit(0)
			}

			// (*buffer)[len(*buffer)-1] = line
			if n > 0 {
				if len(*buffer) > p.maxBufferSize {
					*buffer = (*buffer)[1:]
				}

				// line = []rune{}
				*buffer = append(*buffer, []rune(string(line[:n])))
			}
		}
	}()
}

func (p PTYWindows) Write(text []byte) (int, error) {
	return p.cpty.Write(text)
}

func (p PTYWindows) Close() {
	p.cpty.Close()
}
