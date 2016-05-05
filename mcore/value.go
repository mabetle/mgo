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
