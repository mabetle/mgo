package mrun

import (
	"fmt"
	"reflect"
	"strings"
)

var MethodMap = make(map[string]reflect.Value)

// AddStruct
func AddStruct(v interface{}) {
	typ := reflect.ValueOf(v)
	for i := 0; i < typ.NumMethod(); i++ {
		m := typ.Method(i)
		key := m.Type().Name()
		MethodMap[fmt.Sprintf("m%d%s", i, key)] = m
	}
}

func RunMethod(name string) {
	for k, m := range MethodMap {
		if strings.HasPrefix(k, name) {
			fmt.Printf("Run Method, Key: %s\n", k)
			m.Call([]reflect.Value{})
		}
	}
}

func RunAllMethods() {
	for _, v := range MethodMap {
		v.Call([]reflect.Value{})
	}
}
