package main

import (
	"mabetle/libs/hubs"

	"github.com/mabetle/mgo/mdb/demo"
)

var (
	sql = hubs.GetDemoSql()
)

// SqlDemo
func Demo() {
	demo.InitDB(sql)
}

func main() {
	Demo()
}
