package utils

import (
	"log"
	"os"
)

type Log struct {
	file *os.File
}

func NewLog(filename string) (l Log, err error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	return Log{file: f}, nil
}

func (l Log) Show(text string) {
	log.SetOutput(l.file)
	log.Println(text)
}

func (l Log) Close() {
	l.file.Close()
}
