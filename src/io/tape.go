package main

import "io"

type tape struct {
	file io.ReadWriteSeeker
}

// Write writes len(p) bytes from p to the tape file
func (t *tape) Write(p []byte) (n int, err error) {
	_, _ = t.file.Seek(0, 0)
	return t.file.Write(p)
}
