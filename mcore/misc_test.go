package mcore

import (
	"testing"
	"github.com/mabetle/mgo/mcore/mtest"
)

func TestCheckNullOrBlank(t *testing.T) {
	mtest.TestTrue(t,CheckNullOrBlank("",""))
	mtest.TestFalse(t,CheckNullOrBlank("1",""))
}


