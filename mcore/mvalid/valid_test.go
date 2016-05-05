package mvalid

import (
	"github.com/mabetle/mgo/mcore/mtest"
	"testing"
)

type Model struct {
	Id    string `validator:"required"`
	Email string `validator:"email"`
}

func TestXXX(t *testing.T) {
	m := Model{}
	m.Id = "1"
	m.Email = "demo@demo.com"
	mtest.RegTest(t)
	mtest.AssertEqual(true, true)
	PrintValidate(m)
}
