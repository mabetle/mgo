package mcore_test

import (
	"github.com/mabetle/mcore"
	"testing"
)

func TestJoinPath(t *testing.T) {
	if "a/b/c.txt" != mcore.JoinPath("a", "b", "c.txt") {
		t.Error("join path failed")
	}
	a := mcore.JoinPath("a.txt")
	t.Log(a)
	if "a.txt" != a {
		t.Error(a)
	}
}
