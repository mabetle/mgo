package demo

import (
	"github.com/mabetle/mgo/mcore/mrun"
)

func AddFunc(fn func(), args ...string) {
	mrun.AddFunc(fn, args...)
}

func AddRunner(r mrun.Runner, args ...string) {
	mrun.AddRunner(r, args...)
}

func AddStruct(v interface{}) {
	mrun.AddStruct(v)
}

func Main() {
	mrun.Main()
}
