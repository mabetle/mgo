package csvsdb

import (
	"fmt"
	"strconv"

	"github.com/mabetle/mgo/mcore/msdb"
	"github.com/mabetle/mgo/mcsv"
	"github.com/mabetle/mgo/mlog"
)

var logger = mlog.GetLogger("github.com/mabetle/mgo/mcsv/csvsdb")

type CsvTable struct {
	*msdb.BaseTable
	body [][]string
}

func NewSimpleTable(file string) msdb.SimpleTable {
	return NewCsvTable(file)
}

func NewCsvTable(file string) *CsvTable {
	table := new(CsvTable)
	bt := new(msdb.BaseTable)
	table.BaseTable = bt

	csv := mcsv.NewCSV(file)
	table.BaseTable.Header = csv.GetHeaderRow()
	table.body = csv.GetData()
	table.BaseTable.Cusor = new(msdb.Cusor)
	table.BaseTable.Cusor.MaxIndex = len(table.body)

	table.StringGetter = table

	return table
}

// overide BaseTable
func (t *CsvTable) GetString(col string) (value string) {
	//TODO out of range?
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	iCol, err := strconv.Atoi(col)
	logger.CheckError(err)
	value = t.body[t.Cusor.RowIndex-1][iCol]
	return
}

// Random Access
func (t *CsvTable) GetRowColString(row, col int) (value string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	value = t.body[row][col]
	return
}
