package mcore

import (
	"fmt"
)

// print any type value
func Print(v ...interface{}) {
	fmt.Print(v...)
}

func Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

// after print, start a new line
func Println(v ...interface{}) {
	fmt.Println(v...)
}

// each item in array start a new line.
func PrintStringArray(vs []string) {
	for _, v := range vs {
		fmt.Printf("%v\n", v)
	}
}

// inpect model and print
// model should be a struct
func PrintModel(model interface{}) {
	fmt.Println()
	fmt.Printf("%+v", model)
	fmt.Println()
}

func PrintModels(models []interface{}) {
	for _, v := range models {
		PrintModel(v)
	}
}
