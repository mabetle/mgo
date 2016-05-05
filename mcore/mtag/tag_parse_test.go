package mtag_test

import (
	"github.com/mabetle/mgo/mcore/mtag"
	"testing"
)

type Model struct {
	Name string `tag:"hello"`
}

func TestParse(t *testing.T) {
	v, _ := mtag.GetTag(Model{}, "Name", "tag")
	if v != "hello" {
		t.Error("tag value not equals as expect.")
	}
}
