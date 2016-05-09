package mstat

import (
	"testing"

	"github.com/mabetle/mgo/mcore/mtest"
)

func TestXXX(t *testing.T) {
	mtest.RegTest(t)
	values := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 20.0}
	mtest.AssertEqual(1.0, Min(values))
	mtest.AssertEqual(20.0, Max(values))
}
