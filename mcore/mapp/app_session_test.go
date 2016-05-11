package mapp_test

import (
	"fmt"
	"testing"

	"github.com/mabetle/mgo/mcore/mapp"
)

func init() {
	InitLoad()
}

func InitLoad() {
	file := "testdata/app.conf"
	if err := mapp.LoadAppSessionFile(file); err != nil {
		return
	}
}

func TestAppSession(t *testing.T) {
	if mapp.GetAppSession("appname") != "demo" {
		t.Error("error get app name")
	}
}

func ExampleAppSession() {
	fmt.Printf("%s\n", mapp.GetAppSession("appname"))
	// Output:
	// demo
}
