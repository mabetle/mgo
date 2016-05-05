package main

import (
	"fmt"
	"github.com/mabetle/mgo/mlog"
	. "github.com/mabetle/mgo/mdb/demo/models"
	"github.com/mabetle/mgo/mdb/sql_mysqld"
	"github.com/mabetle/mgo/mdb/xorm_mysqld"
)

var (
	sql  = sql_mysqld.NewDemoSql()
	xorm = xorm_mysqld.NewDemoXorm()
	m    = DemoXorm{}
)

func Migrate() {
	xorm.Sync(DemoXorm{})
}

func Insert() {
	d := DemoXorm{
		Name: "demo",
	}
	d2 := DemoXorm{
		Id:   "demo",
		Name: "Demo2",
	}
	//xorm.Insert(d)
	xorm.UuidSave(&d)
	xorm.SaveOrUpdate(&d2)
	xorm.Save(&d2)
	sql.PrintTable(m.TableName())
}

func Get() {
	demo := DemoXorm{}
	xorm.GetById(&demo, "demo")
	fmt.Printf("%+v\n", demo)
}

func main() {
	mlog.SetTraceLevel()
	sql.DropTable(m.TableName())
	Migrate()
	Insert()
	Get()
}
