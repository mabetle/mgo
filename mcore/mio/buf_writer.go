package mio

import (
	"bufio"
	"fmt"
	"io"
)

// BufWriter wrap bufio.Writer and provide more functions.
type BufWriter struct {
	*bufio.Writer
}

func NewBufWriter(out io.Writer) *BufWriter {
	bout := bufio.NewWriter(out)
	return &BufWriter{Writer: bout}
}

// Writef
func (w BufWriter) Writef(f string, args ...interface{}) (int, error) {
	return w.WriteString(fmt.Sprintf(f, args...))
}

// WriteValues value can be any type
func (w BufWriter) WriteValues(values ...interface{}) (int, error) {
	s := fmt.Sprint(values...)
	return w.WriteString(s)
}
