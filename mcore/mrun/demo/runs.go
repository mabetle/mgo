package demo

import (
	"fmt"
)

func init() {
	AddStruct(&Runs{})
}

func OldAdd() {
	runs := Runs{}
	AddFunc(runs.RunsA, "runsA")
	AddFunc(runs.RunsB, "runsB")
	AddFunc(runs.RunsC, "runsC")
}

// Runs collects
type Runs struct {
}

func (r Runs) RunsA() {
	fmt.Println("Runs A")
}

func (r Runs) RunsB() {
	fmt.Println("Runs B")
}

func (r Runs) RunsC() {
	fmt.Println("Runs C")
}
