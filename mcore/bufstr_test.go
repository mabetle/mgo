package mcore

import (
	"testing"
	"github.com/mabetle/mgo/mcore/mtest"
)

func TestStrBuffer(t *testing.T){
	mtest.RegTest(t)

	sb:=NewStringBuffer()
	sb.Append("Hello")

	mtest.Equal("Hello", sb.String())
}

