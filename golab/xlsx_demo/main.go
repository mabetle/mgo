package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {

	fileName:="/rundata/demo.xlsx"

	var row int=3
	var col int=2
	sheetIndex:=0

	xlFile, _ := xlsx.OpenFile(fileName)
	sheet := xlFile.Sheets[sheetIndex]

	cell:=sheet.Cell(row,col)
	fmt.Printf("%v\n",cell.String())

	fmt.Println("End")
}


