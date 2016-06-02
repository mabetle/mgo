package demo

import (
	"fmt"

	"github.com/mabetle/mgo/mcore/msdb"
)

// DemoSimpleTable
// demo SimpleTable api
func DemoSimpleTable(table msdb.SimpleTable) {

	var rows int = table.GetRows()
	var cols int = table.GetCols()

	fmt.Println("Rows:", rows)
	fmt.Println("Cols:", cols)

	//print head

	//print data
	for table.Next() {
		for i := 0; i < cols; i++ {
			fmt.Printf("%s\t", table.GetString(string(i)))
		}
		fmt.Println()
	}
}
