package mio

import (
	"bufio"
	"bytes"
	"io"
)

func NewBufferWriter() io.Writer {
	b := bytes.NewBuffer(make([]byte, 0))
	return bufio.NewWriter(b)
}

// Buf implents io.Reader.
func NewBufferReader(s string) *bufio.Reader {
	b := bytes.NewBufferString(s)
	return bufio.NewReader(b)
}
