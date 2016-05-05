package mtest

import "testing"

func TestUtil(t *testing.T){
	AssertTrue(true,"")
	AssertFalse(false,"")
	AssertEqual("a", "a", "")
	AssertEqual(1,1,"")
}


func TestTest(t *testing.T){
	TestEqual(t, "1","1")
	TestTrue(t, true)
	TestFalse(t,false)
}

