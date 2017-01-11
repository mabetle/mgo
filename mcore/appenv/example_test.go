package appenv_test

import (
	"fmt"

	"github.com/mabetle/mgo/mcore/appenv"
)

func ExampleGet() {
	appenv.LoadConf(confFile)
	// show be true
	v := appenv.GetString("SHOW_SQL", "False")
	fmt.Printf("%s\n", v)
	// Output: true
}
