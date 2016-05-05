package mcore

import (
	"strconv"
)

type Int int
type Int32 int32
type Int64 int64

func (t Int) Itoa() string {
	return strconv.Itoa(int(t))
}
