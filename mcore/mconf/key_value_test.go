package mconf

import (
	"github.com/mabetle/mcore"
	"github.com/mabetle/mcore/mtest"
	"testing"
)

type TestLoader struct{}

func (l TestLoader) LoadKeyValue() mcore.StringKeyValueMap {
	c := mcore.NewStringKeyValueMap()
	c.Put("a", "a")
	c.Put("b", "0")
	c.Put("c", "5")
	return c
}

func TestKeyValueConfig(t *testing.T) {
	mtest.RegTest(t)
	c := NewConfig(&TestLoader{})

	mtest.AssertEqual(true, c.IsContain("a"))
	mtest.AssertEqual("a", c.GetString("a"))
	mtest.AssertEqual(false, c.GetBool("b"))
	mtest.AssertEqual(c.GetInt("c"), 5)
}
