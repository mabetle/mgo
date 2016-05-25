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

// GetString
func GetString(value interface{}) string {
	if v, ok := value.(string); ok {
		return v
	}

	if v, ok := value.([]byte); ok {
		return string(v)
	}

	return fmt.Sprintf("%v", value)
}

// SepJoin join any value
func SepJoin(sep string, values ...interface{}) string {
	sb := NewStringBuffer()
	for i, v := range values {
		sb.Append(fmt.Sprint(v))
		if i != len(values)-1 {
			sb.Append(sep)
		}
	}
	return sb.String()
}

// Join join value with blank
func Join(values ...interface{}) string {
	return SepJoin("", values...)
}
