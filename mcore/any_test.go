package mcore

import (
	"testing"
	. "github.com/mabetle/mgo/mcore/mtest"
)

func TestIsValueEqual(t *testing.T) {
	AssertTrue(IsValueEqual(123,"123"))
	AssertTrue(IsValueEqual(123.05,"123.05"))
}

