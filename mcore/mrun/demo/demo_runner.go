package demo

import (
	"fmt"
)

type RunA struct{}
type RunB struct{}


func init(){
	AddRunner(RunA{}, "runa")
	AddRunner(RunB{}, "runb")
}

func (r RunA)Run(){
	fmt.Printf("%s\n", "run a")
}

func (r RunB)Run(){
	fmt.Printf("%s\n", "Run B")
}

