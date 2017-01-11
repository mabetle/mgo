package main

import (
	"fmt"

	"github.com/mabetle/mgo/mcore/appenv"
)

func init() {
	appenv.LoadConf("../../testdata/demo.conf")
}

func DemoGet() {
	// show be true
	v := appenv.GetString("SHOW_SQL", "False")
	fmt.Printf("%s\n", v)
}

func main() {
	//appenv.PrintConfs()
	DemoGet()
}
