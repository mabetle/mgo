package mio

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

func NewStringReader(s string) io.Reader {
	return strings.NewReader(s)
}

func NewStringWriter() io.Writer {
	b := bytes.NewBuffer(make([]byte, 0))
	return bufio.NewWriter(b)
}
