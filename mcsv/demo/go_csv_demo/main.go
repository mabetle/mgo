package main

import (
	"github.com/mabetle/mgo/mcsv/csvsdb"
	"github.com/mabetle/mgo/mcore/msdb"
	"github.com/mabetle/mgo/mcsv"
	"github.com/mabetle/mgo/mcore"
)

var csvFile = "../data/demo.csv"

func PrintFile(){
	//mcore.PrintFile(csvFile)
	mcore.PrintFile2(csvFile)
}

func DemoRead(){
	csv:=mcsv.NewCSV(csvFile)
	csv.ShowContent()
}

func DemoSimpleTableRead(){
	//table:=csvsdb.NewCsvTable(csvFile)
	table:=csvsdb.NewSimpleTable(csvFile)
	msdb.DemoSimpleTable(table)
	//table.Demo()
}

func main() {
	//DemoRead()
	PrintFile()
	DemoSimpleTableRead()
}

