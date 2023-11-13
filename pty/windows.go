//go:build windows
// +build windows

package pty

import (
	"siter/config"
	// "siter/utils"

	"github.com/UserExistsError/conpty"
)

func StartProcess(c config.Config) (p PTYWindows, err error) {
	cpty, err := conpty.Start(c.Shell.Command)
	if err != nil {
		return p, err
	}

	return
}

// func (p PTYWindows) Read() io.Reader {
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
// }
