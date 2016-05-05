package mcore

import (
	"strings"
)

// ReadLines
func ReadFileLines(file string) (lines []string, err error) {
	content, err := GetFileContent(file)
	if nil != err {
		return nil, err
	}
	return strings.Split(content, "\n"), nil
}


