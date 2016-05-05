package msdb

import (
	"fmt"
)

func DemoRandomAccessTable(table RandomAccessTable){
	fmt.Println("====Begin Print RandomAccessTable====")

	rows:=table.GetRows()

	if rows<1{
		fmt.Println("Table is empty.")
		return
	}
	
	colNames:=table.GetColNames()
	cols:=len(colNames)
	
	for col := 0; col < cols; col++ {
		fmt.Println("%v\t", colNames[col])
	}
	fmt.Println()
	
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			fmt.Printf("%v\t", table.GetRowColString(row,col))
		}
		fmt.Println()
	}
	
	fmt.Println("====End.. Print RandomAccessTable====")
}

