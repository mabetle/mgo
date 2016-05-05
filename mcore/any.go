package mcore

import (
	"fmt"
)

// IsValueEqual judge by string value
func IsValueEqual(a, b interface{}) bool {
	var sa, sb string
	sa = fmt.Sprint("%v", a)
	sb = fmt.Sprint("%v", b)
	if sa == sb {
		return true
	}
	return false
}

// AppendBefore append array
func AppendBefore(old []interface{}, value interface{}) (r []interface{}) {
	r = append(r, value)
	for _, v := range old {
		r = append(r, v)
	}
	return r
}

// Append append array
func Append(old []interface{}, value interface{}) []interface{} {
	return append(old, value)
}
