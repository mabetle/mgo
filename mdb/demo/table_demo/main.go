package main

import (
	"fmt"
	"mabetle/libs/hubs"
)

var (
	db    = hubs.GetDemoSql()
	table = db.NewTable("demo_table")
)

func Demo() {
	fmt.Println(table.CountRows())
	fmt.Println(table.CountColumns())
	fmt.Println(table.GetColumns())
	table.Print()
}

func main() {
	Demo()
}
