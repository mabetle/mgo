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

func getTest() *testing.T {
	if T == nil {
		fmt.Printf("No T\n")
	}
	return T
}

func Equal(input, expect interface{}, msgs ...interface{}) {
	TestEqual(getTest(), input, expect, msgs...)
}

func True(input bool, msgs ...interface{}) {
	TestTrue(getTest(), input, msgs...)
}

func False(input bool, msgs ...interface{}) {
	TestFalse(getTest(), input, msgs...)
}

// AssertTrue
func AssertTrue(value bool, msgs ...interface{}) {
	AssertEqual(value, true, msgs...)
}

// AssertFalse
func AssertFalse(value bool, msgs ...interface{}) {
	AssertEqual(value, false, msgs...)
}

// AssertEqual
func AssertEqual(input interface{}, expect interface{}, msgs ...interface{}) {
	TestEqual(T, input, expect, msgs...)
}

func TestEqual(t *testing.T, input, expect interface{}, msgs ...interface{}) {
	sInput := fmt.Sprintf("%v", input)
	sExpect := fmt.Sprintf("%v", expect)
	if sInput == sExpect {
		return
	}
	t.Errorf("Error: input: %v , expect: %v\n", input, expect)
	if len(msgs) > 0 {
		if t != nil {
			t.Error(msgs...)
		} else {
			fmt.Printf(fmt.Sprint(msgs...))
		}
	}
}

func TestTrue(t *testing.T, value bool, msgs ...interface{}) {
	TestEqual(t, value, true, msgs...)
}

func TestFalse(t *testing.T, value bool, msgs ...interface{}) {
	TestEqual(t, value, false, msgs...)
}
