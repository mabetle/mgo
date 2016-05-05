package mcore

import (
	"strconv"
)

const (
	TRUE  = true
	FALSE = false
)

const (
	B_TRUE  = Bool(TRUE)
	B_FALSE = Bool(FALSE)
)

type Bool bool

func (b Bool) Format() string {
	return strconv.FormatBool(bool(b))
}
