package mcore

import (
	"strconv"
)

type Byte byte
type Bytes []byte

func (bs Bytes) AppendBool(b bool) Bytes {
	r := strconv.AppendBool([]byte(bs), b)
	return Bytes(r)
}

func (bs Bytes) AppendQuote(s string) Bytes {
	r := strconv.AppendQuote([]byte(bs), s)
	return Bytes(r)
}

func (bs Bytes) AppendQuoteRune(r rune) Bytes {
	t := strconv.AppendQuoteRune([]byte(bs), r)
	return Bytes(t)
}
