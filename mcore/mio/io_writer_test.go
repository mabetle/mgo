package mio

import (
	"github.com/mabetle/mgo/mcore/mtest"
	"os"
	"testing"
)

func TestWriteText(t *testing.T) {
	mtest.RegTest(t)
	WriteText(os.Stdout, "test write to os.Stdout\n")
	WriteText(os.Stderr, "test write to os.Stderr\n")
}

func TestWriteLines(t *testing.T) {
	mtest.RegTest(t)
	lines := []string{"line 1", "line 2", "line 3"}
	WriteLines(os.Stdout, lines)
}

func TestWriteStrings(t *testing.T) {
	mtest.RegTest(t)
	WriteStrings(os.Stdout, "hello,", "demo write strings.")
}
