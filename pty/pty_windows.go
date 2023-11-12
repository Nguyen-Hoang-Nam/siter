//go:build windows
// +build windows

package pty

import (
	"errors"
	"io"

	"siter/config"
	// "siter/utils"

	"github.com/UserExistsError/conpty"
)

type PTYWindows struct {
	process       *conpty.ConPty
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

	return PTYWindows{process: cpty, shellCommand: c.Shell, maxBufferSize: c.ScrollbackLines}, nil
}

func (p PTYWindows) Read() io.Reader {
	return p.process
	// go func() {
	// 	// logger, _ := utils.NewLog("siter-log.txt")
	// 	// defer logger.Close()

	// 	line := make([]byte, 1000)

	// 	for {
	// 		n, err := p.cpty.Read(line)
	// 		if err != nil {
	// 			os.Exit(0)
	// 		}

	// 		if n > 0 {
	// 			if p.shellCommand == "cmd" {
	// 				if len(*buffer) > p.maxBufferSize {
	// 					*buffer = (*buffer)[1:]
	// 				}

	// 				// logger.Show(string(line[:n]))

	// 				*buffer = append(*buffer, []rune(string(line[:n])))
	// 			} else if p.shellCommand == "powershell" {
	// 				*buffer = append(*buffer, []rune("Not supported powershell yet."))
	// 			} else {
	// 				*buffer = append(*buffer, []rune("Not supported powershell yet."))
	// 			}

	// 		}
	// 	}
	// }()
}

func (p PTYWindows) Write(text []byte) (int, error) {
	return p.process.Write(text)
}

func (p PTYWindows) Close() {
	p.process.Close()
}
