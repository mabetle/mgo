package mcore

import (
	"testing"
)

func TestV(t *testing.T) {
	renderArgs := make(map[string]interface{})
	renderArgs["hi"] = "hi"
	wa := NewWrapArgs(renderArgs, "demo=demo")
	if wa.GetArgString("demo", "") != "demo" {
		t.Error("demo not equal")
	}
	if wa.GetArgString("hi", "") != "hi" {
		t.Error("hi not equal")
	}
	if wa.GetArgString("none", "none") != "none" {
		t.Error("none wrong")
	}
	if wa.GetArgBool("hi", false) != false {
		t.Error("bool false")
	}

}
