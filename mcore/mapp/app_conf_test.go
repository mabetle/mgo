package mapp

import (
	"github.com/mabetle/mgo/mcore/mtest"
	"testing"
)

var (
	dbcConf = NewDefaultAppConf("dbc")
)

func init() {
	dbcConf.Init()
}

func TestAppConf(t *testing.T) {
	mtest.RegTest(t)
	mtest.AssertEqual(false, NewDefaultAppConf("none").IsExist())
	mtest.AssertEqual(dbcConf.GetVendorName(), "Mabetle")
}
