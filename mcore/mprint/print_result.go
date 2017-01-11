package mprint

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func fprintHead(out io.Writer, head string) {
	if head != "" {
		fmt.Fprintf(out, "==== %s ====\n", head)
	}
}

func FprintSlice(out io.Writer, vs interface{}) {
	v := reflect.ValueOf(vs)
	// not a slice
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		n := v.Cap()
		if n == 0 {
			fmt.Fprintf(out, "No results.\n")
		}
		for i := 0; i < n; i++ {
			vv := v.Index(i).Interface()
			fmt.Fprintf(out, "%d.%v\n", i+1, vv)
		}
	default:
		fmt.Fprintf(out, "not a slice.\n")
	}
}

func PrintSlice(vs interface{}) {
	FprintSlice(os.Stdout, vs)
}

func FprintMap(out io.Writer, vs interface{}) {
	v := reflect.ValueOf(vs)
	switch v.Kind() {
	case reflect.Map:
		keys := v.MapKeys()
		for i, key := range keys {
			vv := v.MapIndex(key).Interface()
			fmt.Fprintf(out, "%d.%v:%v\n", i+1, key, vv)
		}
	default:
		fmt.Fprintf(out, "not a map.\n")
	}
}

func FprintStruct(out io.Writer, vs interface{}) {
	v := reflect.ValueOf(vs)
	switch v.Kind() {
	case reflect.Struct:
		fmt.Fprintf(out, "%+v", vs)
	default:
		fmt.Fprintf(out, "not a struct.\n")
	}
}

// Fprint out
func Fprint(out io.Writer, vs interface{}) {
	v := reflect.ValueOf(vs)
	switch v.Kind() {
	//case reflect.Array:
	case reflect.Slice, reflect.Array:
		FprintSlice(out, vs)
	case reflect.Map:
		FprintMap(out, vs)
	case reflect.Struct:
		FprintStruct(out, vs)
	default:
		fmt.Fprintf(out, "%v\n", vs)
	}
}

func Print(vs interface{}) {
	Fprint(os.Stdout, vs)
}

func HeadPrint(head string, vs interface{}) {
	fprintHead(os.Stdout, head)
	Print(vs)
}

func HeadFprint(out io.Writer, head string, vs interface{}) {
	fprintHead(os.Stdout, head)
	Fprint(out, vs)
}
