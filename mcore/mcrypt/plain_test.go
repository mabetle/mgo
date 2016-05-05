package mcrypt

import (
	"github.com/mabetle/mcore/mtest"
	"testing"
)

var (
	plain = NewPlain()
)

func TestPlainEncode(t *testing.T) {
	mtest.RegTest(t)
	mtest.AssertEqual(plain.Encode("demo"), "demo")
}
