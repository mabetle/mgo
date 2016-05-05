package mrun

import (
	"fmt"
	"strings"
)

type Runner interface {
	Run()
}

var (
	runMap = make(map[string]Runner)
)

var runnerIndex = 0

// AddRunner
func AddRunner(v Runner, args ...string) {
	var key string
	if len(args) < 1 {
		key = fmt.Sprintf("%v\n", v)
	} else {
		key = args[0]
	}

	runMap[fmt.Sprintf("r%d%s", runnerIndex, key)] = v
	runnerIndex++
}

func RunAllRunners() {
	for key, _ := range runMap {
		RunKey(key)
	}
}

func RunKey(key string) {
	for k, v := range runMap {
		if strings.HasPrefix(k, key) {
			fmt.Printf("Run Runner, Key: %s\n", k)
			v.Run()
		}
	}
}

// Run a runner
func Run(v Runner, args ...interface{}) {
	v.Run()
}

// run runner until C-C stop apps.
func LoopRun(v Runner, args ...interface{}) {
	for {
		Run(v, args...)
	}
}

// run runer for specific times
func TimesRun(v Runner, times int, args ...interface{}) {
	for i := 0; i < times; i++ {
		Run(v, args...)
	}
}
