package demo

import (
	"fmt"
	"github.com/mabetle/mcore"
)

func init() {
	AddFunc(DemoUpper, "upper")
}

// DemoUpper demo
func DemoUpper() {

	fmt.Println(mcore.ToUpper("Hello"))
}
