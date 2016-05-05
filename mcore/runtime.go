package mcore

import (
	"runtime"
)

var O_NL = "\n"

func GetOS() string {
	return runtime.GOOS
}

func IsWindows() bool {
	return String(GetOS()).IsContainIgnoreCase("WINDOWS")
}

// FIXME not work in linux
func IsLinux() bool {
	return String(GetOS()).IsContainIgnoreCase("LINUX")
}

// FIXME not work in MacOS
func IsDarwin() bool {
	return String(GetOS()).IsContainIgnoreCase("DARWIN")
}
