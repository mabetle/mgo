package msdb

import (
	"fmt"
)

func DemoSimpleTable(table SimpleTable) {
	fmt.Println("====Begin Print SimpleTable====")
	fmt.Println("Cols:", table.GetCols())
	fmt.Println("Rows:", table.GetRows())

	//print head
	fmt.Print("Row\t")
	colNames := table.GetColNames()
	//cols:=len(colNames)

	for _, colName := range colNames {
		fmt.Printf("%v\t", colName)
	}
	fmt.Println()

	// print data
	for table.Next() {
		fmt.Printf("%d\t", table.GetRowIndex())
		for _, colName := range colNames {
			fmt.Printf("%v\t", table.GetString(colName))
		}
		fmt.Println()
	}

	fmt.Println("====End.. Print SimpleTable====")
}
