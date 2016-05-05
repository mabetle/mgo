package mcon

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	NEW_LINE_BYTE = byte(10)
)

// ReadLine from os.Stdio
func ReadLine() (result string) {
	r := bufio.NewReader(os.Stdin)
	result, _ = r.ReadString(NEW_LINE_BYTE)
	if strings.HasSuffix(result, "\n") {
		result = strings.TrimSuffix(result, "\n")
	}
	return
}

// ReadLineWithMsg
func ReadLineWithMsg(msg string) string {
	fmt.Print(msg)
	return ReadLine()
}

func ReadNotBlankLine() (result string) {
	for {
		result = ReadLine()
		if result == "" {
			fmt.Println("input blank line, try again")
		} else {
			break
		}
	}
	return
}

func ReadNotBlankLineWithMsg(msg string) string {
	fmt.Print(msg)
	return ReadNotBlankLine()
}
