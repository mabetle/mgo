package mcore

import (
	"os"
)

// Expand expand env
func Expand(v string) string {
	return os.ExpandEnv(v)
}
