package mcore

import (
	"testing"
	"github.com/mabetle/mgo/mcore/mtest"
)

func TestGetLocation(t *testing.T){
	mtest.RegTest(t)
	mtest.AssertEqual("abc/def.go", GetLocation("abc","def.go"))
	mtest.AssertEqual("abc/def.go", GetLocation("abc/","def.go"))
}
