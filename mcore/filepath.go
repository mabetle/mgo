package mcore

import (
	"os"
	"strings"
)

// Home returns user home path
func Home() string {
	return GetEnv("HOME")
}

// ExpandPath expand ~,%VAR%,$VAR,$(VAR),${VAR) to real path
func ExpandPath(path string) string {
	// replace ~ to $HOME
	if strings.HasPrefix(path, "~") {
		path = "$HOME" + strings.TrimPrefix(path, "~")
	}
	return os.ExpandEnv(path)
}

// JoinPath join path
func JoinPath(paths ...string) string {
	sb := NewStringBuffer()
	for i, path := range paths {
		// expand it
		path = ExpandPath(path)
		// delete end /
		path = strings.TrimRight(path, "/")
		// not last one and  and not end with /
		if i != 0 && !strings.HasSuffix(path, "/") {
			path = "/" + path
		}
		sb.Append(path)
	}
	return sb.String()
}
