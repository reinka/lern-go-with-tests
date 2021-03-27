package poker

import (
	"os"
)

type tape struct {
	file *os.File
}

// Write writes len(p) bytes from p to the tape file
func (t *tape) Write(p []byte) (n int, err error) {
	_ = t.file.Truncate(0)
	_, _ = t.file.Seek(0, 0)
	return t.file.Write(p)
}
