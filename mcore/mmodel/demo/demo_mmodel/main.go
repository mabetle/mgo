package main

import (
	"fmt"
	"github.com/mabetle/mgo/mcore/mmodel"
)

type Model struct {
	Id   string
	Name string
}

func (m Model) TableName() string {
	return "demo_model"
}

func Demo(m1 interface{}) {
	fmt.Println(mmodel.GetModelId(m1))
	m2 := mmodel.AddModelUuid(m1)
	fmt.Printf("%+v\n", m1)
	fmt.Printf("%+v\n", m2)
}

func DemoA() {
	m1 := &Model{Id: "1", Name: "Demo1"}
	Demo(m1)
	m2 := Model{Id: "1", Name: "Demo1"}
	Demo(m2)
}

func DemoTableName() {
	m := Model{}
	fmt.Printf("%s\n", mmodel.GetModelTableName(m))
	m2 := &Model{}
	fmt.Printf("%s\n", mmodel.GetModelTableName(m2))
}

func main() {
	DemoTableName()
}
