package demo

import (
	"fmt"
	"github.com/mabetle/mgo/mcore/mconf"
)

// Config api Demo
func DemoConfig(mconf mconf.Config) {
	//should be false
	fmt.Println("IsContain('no-exist'):", mconf.IsContain("no-exist"))
	fmt.Println("GetString('db.spec'):", mconf.GetString("db.spec"))
	fmt.Println("GetString('db2.spec'):", mconf.GetString("db2.spec"))
}
