package mcore

import (
	"testing"
	"github.com/mabetle/mcore/mtest"
)

// DemoDemo  > demo_demo
func TestToTableName(t *testing.T){
	mtest.RegTest(t)
	mtest.AssertEqual("demo_demo", ToTableName("DemoDemo"))
}

func TestToCamel(t *testing.T){
	mtest.RegTest(t)
	mtest.AssertEqual("DemoDemo", ToCamel("demo_demo"))
}
