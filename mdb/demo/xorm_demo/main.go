package main

import (
	"fmt"
	"mabetle/libs/hubs"

	. "github.com/mabetle/mgo/mdb/demo/models"
	"github.com/mabetle/mgo/mlog"
)

var (
	sql  = hubs.GetDemoSql()
	xorm = hubs.GetDemoXorm()
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
