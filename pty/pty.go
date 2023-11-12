package pty

import "io"

type IPTY interface {
	Read() io.Reader
	Write([]byte) (int, error)
	Close()
}
