package wrobfig

import (
	"github.com/mabetle/mcore"
	"github.com/mabetle/mcore/mtest"
	"testing"
)

func TestReadConfigFromString(t *testing.T){
	mtest.RegTest(t)

	s:=`
hello=hello from string
name=demo
	`
	c,err:=ReadConfigFromString(s)
	mcore.CheckError(err)

	v, _:=c.String("","hello")
	mtest.Equal("hello from string", v)
}

