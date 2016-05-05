package mcore

import (
	"github.com/mabetle/mgo/mcore/mtest"
	"testing"
)

var (
	skv StringKeyValueMap = make(map[string]string)
)

func init() {
	skv["a"] = "a"
	skv["b"] = "b"
	skv["c"] = "T"
}

func TestIsContain(t *testing.T) {
	mtest.RegTest(t)
	mtest.AssertEqual(true, skv.IsContain("a"))
	mtest.AssertEqual(true, skv.GetBool("c"))
}
