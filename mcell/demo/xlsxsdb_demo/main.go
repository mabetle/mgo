package main

import (
	"github.com/mabetle/mgo/mcell/xlsxsdb"
	"github.com/mabetle/mgo/mcore/msdb"
	//"github.com/mabetle/mgo/mcore"
	"github.com/mabetle/mgo/mcell/wxlsx"
	"github.com/mabetle/mgo/mlog"
)

var(
	logger = mlog.GetLogger("main")
	file="../data/demo.xlsx"
	sheetIndex = 0
)

func PrintSheet(){
	wxlsx.PrintSheetByIndex(file,sheetIndex)
}

func DemoRead(){
	table, err:=xlsxsdb.NewSimpleTable(file, sheetIndex)
	logger.CheckError(err)
	msdb.DemoSimpleTable(table)
}

func main() {
	//mlog.SetDebugLevel()
	mlog.SetInfoLevel()
	PrintSheet()
	DemoRead()
}


