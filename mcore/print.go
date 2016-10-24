package mcore

import (
	"fmt"
	"reflect"
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

// PrintModels .
func PrintModels(models []interface{}) {
	for _, v := range models {
		PrintModel(v)
	}
}

// PrintRows .
func PrintRows(model interface{}, include, exclude string) {
	v := reflect.ValueOf(model)
	if v.Kind() != reflect.Slice {
		fmt.Printf("Not a slice\n")
		return
	}
	fs := GetUsedArrayFields(model, include, exclude)
	for _, f := range fs {
		fmt.Printf("%s\t", f)
	}
	fmt.Printf("\n")
	// if not a slice, return input arg
	for i := 0; i < v.Len(); i++ {
		row := v.Index(0).Interface()
		for _, f := range fs {
			fv := GetFieldValue(row, f)
			fmt.Printf("%v\t", fv)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
