package mcore

import (
	"github.com/mabetle/mgo/mcore/mtest"
	"testing"
)

func TestSepEnd(t *testing.T) {
	mtest.RegTest(t)
	mtest.AssertEqual("def", String("/aaa/abc.def").SepEnd("."))
	mtest.AssertEqual("def", String("/aaa/abc/def").SepEnd("/"))
	mtest.AssertEqual("def", String("/aaa/abc\\def").SepEnd("\\"))
	mtest.AssertEqual("def", String("/aaa/abc@def").SepEnd("@"))
	mtest.AssertEqual("", String("/aaa/abc@def").SepEnd("#"))
}

func TestTrimEndIndex(t *testing.T) {
	mtest.RegTest(t)

	mtest.AssertEqual(String("/a/b/c.txt").TrimEndIndex("/"), "/a/b")
}

func TestTrimBeginIndex(t *testing.T) {
	mtest.RegTest(t)

	mtest.AssertEqual(String("a/b/c.txt").TrimBeginIndex("/"), "b/c.txt")
}
