package demo

import (
	"fmt"
	"github.com/mabetle/mgo/mcore"
	"runtime"
)

func init() {
	AddFunc(ShowOS, "showos")
	AddFunc(JudgeOS, "judgeos")
}

// ShowOS show os
func ShowOS() {
	fmt.Printf("%v\n", runtime.GOOS)
}

// JudgeOS judge os
func JudgeOS() {
	w := mcore.IsWindows()
	fmt.Printf("Is Win:%v\n", w)
}
