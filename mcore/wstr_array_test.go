package mcore

import (
	. "github.com/mabetle/mgo/mcore/mtest"
	"testing"
)

func TestStringArrayMerge(t *testing.T) {
	a := StringArray([]string{"Hello", "World"})
	m := a.Merge()
	exp := "Hello\nWorld\n"
	AssertEqual(m, exp)
}

func TestStringArrayRemoveBlank(t *testing.T) {
	sa := []string{"a", " ", "b"}
	a := StringArray(sa)
	AssertEqual(a.RemoveBlank().Merge(), "a\nb\n")

	AssertEqual(StringArray([]string{"a", " \t ", "b"}).RemoveBlank().Merge(), "a\nb\n")
}

func TestStringArrayRemoveComment(t *testing.T) {
	AssertEqual(
		StringArray([]string{"a", "b", "# c"}).RemoveComment("#").Merge(),
		"a\nb\n")
}

func TestStringArrayJoin(t *testing.T) {
	sa := NewStringArray("a", "b", "c")
	AssertEqual(sa.Join(","), "a,b,c")
}
