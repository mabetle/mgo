package mrun

import (
	"fmt"
	"strings"
)

var funcMap = make(map[string]func())

var funcIndex = 0

// AddFunc
func AddFunc(fn func(), args ...string) {
	var key string
	if len(args) < 1 {
		key = fmt.Sprintf("%v\n", fn)
	} else {
		key = args[0]
	}
	key = fmt.Sprintf("f%d%s", funcIndex, key)
	funcMap[key] = fn
	funcIndex++
}

func RunFunc(fn func()) {
	fn()
}

func RunAllFuncs() {
	for key, _ := range funcMap {
		RunFuncWithKey(key)
	}
}

func RunFuncWithKey(key string) {
	for k, v := range funcMap {
		if strings.HasPrefix(k, key) {
			fmt.Printf("Run Func, Key: %s\n", k)
			v()
		}
	}
}

func LoopRunFunc(fn func()) {
	for {
		go fn()
	}
}

func TimesRunFunc(fn func(), times int) {
	for i := 0; i < times; i++ {
		go fn()
	}
}
