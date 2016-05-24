package mcore

import (
	"fmt"
)

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
