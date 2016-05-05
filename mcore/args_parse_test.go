package mcore

import (
	"testing"
	"github.com/mabetle/mcore/mtest"
)


func TestArgsA(t *testing.T){
	mtest.RegTest(t)

	s:="-d -f=demo.txt Hello Demo"
	a:=NewArgsFromString(s)

	mtest.Equal(4, a.NArgs())
}


