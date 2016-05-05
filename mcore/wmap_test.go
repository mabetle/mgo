package mcore

import (
	"github.com/mabetle/mcore/mtest"
	"testing"
)

var m KeyValueMap = make(map[string]interface{})

func InitMap() {
	m["a"] = "a"
	m["b"] = 3
	m["c"] = "c"
}

func TestIsMapHasKey(t *testing.T) {
	mtest.RegTest(t)
	InitMap()
	mtest.AssertTrue(IsMapHasKey(m, "a"))
	mtest.AssertTrue(!IsMapHasKey(m, "none"))
}

func TestGetMapKeys(t *testing.T) {
	InitMap()
	// order may not equal.
	mtest.AssertEqual(len(GetMapKeys(m)), 3)
}
