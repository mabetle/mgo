package mcore

import (
	"github.com/mabetle/mcore/mtest"
	"testing"
)

type Demo struct {
	Name     string
	Age      int
	RealName string
}

func (m Demo) Hello() string {
	return "hello"
}

var (
	demo = Demo{
		Name:     "demo",
		Age:      20,
		RealName: "Demo",
	}
)

func TestGetFields(t *testing.T) {
	mtest.RegTest(t)
	mtest.AssertTrue(true)
}

func TestIsHasMethod(t *testing.T) {
	mtest.RegTest(t)
	mtest.AssertTrue(IsHasMethod(demo, "Hello"))
	mtest.AssertFalse(IsHasMethod(demo, "None"))
}
