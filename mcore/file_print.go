package mcore

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// PrintFile
func PrintFile(location string) {
	if content, err := ioutil.ReadFile(location); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%v\n", string(content))
	}
}

// PrintFile2
func PrintFile2(location string) {
	f, err := os.OpenFile(location, os.O_RDONLY, 0660)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	cat(bufio.NewScanner(f))
}

// cat
func cat(scanner *bufio.Scanner) error {
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return scanner.Err()
}

// PrintSubFiles
// exts format: .xxx,yyy
func PrintSubFiles(dir string, r bool, exts string) {
	result := GetSubFiles(dir, r, exts, "", "")
	for _, v := range result {
		fmt.Println(v)
	}
}

// PrintFileWithLineNumber
func PrintFileWithLineNumber(file string) {
	data, err := ReadFileLines(file)
	if nil != err {
		fmt.Println(err)
		return
	}
	lineNums := len(data)
	width := len(string(lineNums)) + 1
	for i := 0; i < lineNums; i++ {
		fmt.Printf("%s %s\n", GetFixedWidthNum(i+1, width), data[i])
		//fmt.Printf("%d %s\n",i+1,data[i])
	}
}
