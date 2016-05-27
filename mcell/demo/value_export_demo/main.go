package main

import (
	"mabetle/libs/dbc"
	"mabetle/libs/wagerecord"

	"github.com/mabetle/mgo/mcell/wxlsx"
	"github.com/mabetle/mgo/mlog"
)

var (
	logger = mlog.GetLogger("main")
	xorm   = dbc.GetXorm()
)

func main() {
	m := &wagerecord.WageRecord{}
	m.UserName = "zsc"
	m.TheYear = 2014
	var rows []wagerecord.WageRecord
	xorm.Find(&rows, m)
	f, err := wxlsx.ValueToExcel("", rows, "", "Id")
	if logger.CheckError(err) {
		return
	}
	f.Save("/rundata/output/demo.xlsx")
}
