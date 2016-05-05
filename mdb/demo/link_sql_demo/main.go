package main


import (
	"github.com/mabetle/mgo/mdb"
	"github.com/mabetle/mgo/mdb/demo/ioc"
)

var(
	db  =  ioc.GetSql()
)

func main() {
	mdb.NewLinkSql(db).SetTable("demo_table").Print()
}

