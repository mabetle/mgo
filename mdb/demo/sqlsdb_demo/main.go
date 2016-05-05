package main

import (
	"github.com/mabetle/mgo/mdb/demo/ioc"
	"github.com/mabetle/mgo/mdb/sqlsdb"
)

var(
	sql = ioc.GetSql()
)

func Demo(){
	t:="demo"
	//db.PrintTable(t)

	sql.PrintQueryData("select * from demo")

	table:=sqlsdb.NewSqlTable(sql, t)
	table.Demo()
}

func main() {
	Demo()
}


