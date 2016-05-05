package mmath

import (
	"testing"
	"github.com/mabetle/mgo/mcore/mtest"
)


func TestRound(t *testing.T){
	mtest.RegTest(t)
	mtest.AssertTrue(RoundInt(3.14159)==3)
	mtest.AssertTrue(RoundInt(3.54159)==4)

	mtest.AssertTrue(Round(3.54159,1)==3.5)
	mtest.AssertTrue(Round(3.54159,2)==3.54)
	mtest.AssertTrue(Round(3.54159,3)==3.542)
	mtest.AssertTrue(Round(13.54159,-1)==10)

	mtest.AssertTrue(RoundUp(3.54159,1)==3.6)
	mtest.AssertTrue(RoundUp(3.54159,2)==3.55)
	mtest.AssertTrue(RoundUp(3.54159,3)==3.542)
	mtest.AssertTrue(RoundUp(13.54159,-1)==20)

	mtest.AssertTrue(RoundDown(3.54159,1)==3.5)
	mtest.AssertTrue(RoundDown(3.54159,2)==3.54)
	mtest.AssertTrue(RoundDown(3.54159,3)==3.541)
	mtest.AssertTrue(RoundDown(13.54159,-1)==10)

}



