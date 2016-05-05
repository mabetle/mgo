package demo

import (
	"fmt"
	"github.com/mabetle/mcore/msdb"
)

func DemoRandomAccessTable(table msdb.RandomAccessTable){
	var rows int = table.GetRows()
	var cols int = table.GetCols()

	fmt.Println("Rows:",rows)
	fmt.Println("Cols:",cols)

	//print head

	//print data
	for row := 0; row < rows; row++ {
		for col	:= 0; col < cols; col++ {
			fmt.Printf("%s\t",table.GetString(row,col))
		}
		fmt.Println()
	}

}
