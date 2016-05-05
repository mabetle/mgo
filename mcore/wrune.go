package mcore

import (
	"strconv"
)

type Rune rune

func (t Rune) IsPrint() bool {
	return strconv.IsPrint(rune(t))
}
