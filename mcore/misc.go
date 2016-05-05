package mcore

import (
	"fmt"
	"strings"
)

// check value if null or blank
func CheckNullOrBlank(value interface{}, format string, args ...interface{}) bool {
	return CheckNullOrBlankWithErrMsg(value, fmt.Sprintf(format, args...))
}

func CheckNullOrBlankWithErrMsg(value interface{}, errMsg string) bool {
	if String(fmt.Sprintf("%v", value)).IsBlank() {
		fmt.Println(errMsg)
		return true
	}
	return false
}

// IsIncludeExcludeIn
func IsIncludeExcludeIn(field string, fields []string, include, exclude string) bool {
	includes := strings.Split(include, ",")
	excludes := strings.Split(exclude, ",")

	keyString := String(field)

	// exist include and not in include, false
	if include != "" && !keyString.IsInArrayIgnoreCase(includes) {
		return false
	}

	// if include == "", then only check exclude

	// in exclude, false
	if keyString.IsInArrayIgnoreCase(excludes) {
		return false
	}

	return true
}
