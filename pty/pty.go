package pty

import (
	"errors"
	"runtime"
	"siter/config"
)

type PTYProcess interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
}

var ErrUnsupportedOS = errors.New("UNSUPPORTED_OS")

func Start(c *config.Config) (process PTYProcess, err error) {
	goos := runtime.GOOS
	if goos != "windows" && goos != "linux" && goos != "darwin" {
		return process, ErrUnsupportedOS
	}

	process, err = StartProcess(c)
	if err != nil {
		return nil, err
	}

	return

}
