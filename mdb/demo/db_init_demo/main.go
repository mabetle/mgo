package main

import (
	"github.com/mabetle/mgo/mdb/demo/ioc"
	"github.com/mabetle/mgo/mdb/demo"
)

var (
	sql = ioc.GetSql()
)

func InitDB(){
	demo.InitDB(sql)
}

func main() {
	InitDB()
}

