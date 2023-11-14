package config

import (
	"os"
	"runtime"
)

type parsingShell struct{ Command string }

func (s *parsingShell) UnmarshalText(text []byte) error {
	command := string(text)
	if command == "." {
		goos := runtime.GOOS
		if goos == "windows" {
			command = "cmd"
		} else if goos == "linux" || goos == "darwin" {
			command = os.Getenv("SHELL")
			if command == "" {
				command = "sh"
			}
		}
	}

	s.Command = command

	return nil
}
