package mterm

import (
	//"fmt"
	"os"
	"strings"
)

func IsXterm() bool {
	t := strings.ToLower(os.Getenv("TERM"))
	rt := strings.ToLower(os.Getenv("RUN_TERM"))
	if strings.Contains(t, "xterm") || 
		strings.Contains(t, "linux") || 
		strings.Contains(rt, "xterm") {
		return true
	} else {
		return false
	}
}
