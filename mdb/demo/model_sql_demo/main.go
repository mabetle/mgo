package main

import (
	"github.com/mabetle/mgo/mdb/demo/ioc"
	"fmt"
	. "github.com/mabetle/mgo/mdb/demo/models"
)


var (
	sql = ioc.GetSql()
	m=new(DemoTable)
)

func Demo(){
	sql.PrintModelTable(m)
	//var	rows []DemoTable
	//sql.ModelQuery(&rows, "")
	//fmt.Println(rows)
}

func main() {
	Demo()
	fmt.Println()
}

