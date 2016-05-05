package demo

import (
	"github.com/mabetle/mcore/mrun"
)

// AddFunc add func
func AddFunc(fn func(), args ...string) {
	mrun.AddFunc(fn, args...)
}

// AddRunner add runner
func AddRunner(r mrun.Runner, args ...string) {
	mrun.AddRunner(r, args...)
}

// Main main
func Main() {
	mrun.Main()
}
