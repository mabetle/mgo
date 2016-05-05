package mcsv

import (
	"github.com/mabetle/mgo/mcore"
	"github.com/mabetle/mgo/mlog"
)

var (
	logger = mlog.GetLogger("github.com/mabetle/mgo/mcsv")
)

func (csv CSV) Write(file string) {

}

// WriteData
func WriteData(location string, data [][]string) error {
	if location == "" {
		location = "/rundata/export.csv"
	}
	content := BuildCsvContent(data)
	_, err := mcore.WriteFile(location, content)
	logger.Debug("Write CSV data to :", location, ", Result: ", err)
	return err
}

func BuildCsvContent(data [][]string) (r string) {
	rows := len(data)
	cols := len(data[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if col == cols-1 {
				r = r + data[row][col]
			} else {
				r = r + data[row][col] + ","
			}
		}
		r = r + "\n"
	}
	return
}
