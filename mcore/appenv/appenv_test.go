package appenv_test

import (
	"fmt"
	"testing"

	"github.com/mabetle/mgo/mcore/appenv"
)

var confFile = "testdata/demo.conf"

func Init() {
	appenv.LoadConf(confFile)
}

func TestGet(t *testing.T) {
	Init()
	v := appenv.GetString("SHOW_SQL", "False")
	fmt.Printf("SHOW_SQL:%s\n", v)
	if v != "true" {
		t.Errorf("ShowSql should be true")
	}
}
