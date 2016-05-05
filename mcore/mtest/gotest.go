// mtest
package mtest

import (
	"fmt"
	"testing"
)

var T *testing.T

func RegTest(t *testing.T) {
	T = t
}

func Equal(input, expect interface{}, msg ...interface{}) {
	TestEqual(T, input, expect, msg...)
}

func True(input bool, msg ...interface{}) {
	TestTrue(T, input, msg...)
}

func False(input bool, msg ...interface{}) {
	TestFalse(T, input, msg...)
}

// AssertTrue
func AssertTrue(value bool, msg ...interface{}) {
	AssertEqual(value, true, msg...)
}

// AssertFalse
func AssertFalse(value bool, msg ...interface{}) {
	AssertEqual(value, false, msg...)
}

// AssertEqual
func AssertEqual(input interface{}, expect interface{}, msg ...interface{}) {
	TestEqual(T, input, expect, msg...)
}

func TestEqual(t *testing.T, input, expect interface{}, msg ...interface{}) {
	sInput := fmt.Sprintf("%v", input)
	sExpect := fmt.Sprintf("%v", expect)
	if sInput == sExpect {
		return
	}
	t.Errorf("Error: input: %v , expect: %v\n", input, expect)
	if len(msg) > 0 {
		t.Error(msg...)
	}
}

func TestTrue(t *testing.T, value bool, msg ...interface{}) {
	TestEqual(t, value, true, msg...)
}

func TestFalse(t *testing.T, value bool, msg ...interface{}) {
	TestEqual(t, value, false, msg...)
}
