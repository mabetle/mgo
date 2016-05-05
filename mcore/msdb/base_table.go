package msdb

import (
	"fmt"
	"github.com/mabetle/mcore"
)

// BaseTable SmipleTable base
type BaseTable struct {
	*Cusor
	StringGetter
	Header []string
}

// Next for loop rows
func (t *BaseTable) Next() bool {
	return t.Cusor.Next()
}

// GetRowIndex returns row cursor
func (t *BaseTable) GetRowIndex() int {
	return t.Cusor.RowIndex
}

// GetRows returns table rows
func (t *BaseTable) GetRows() int {
	return t.Cusor.MaxIndex
}

// GetCols return table cols
func (t *BaseTable) GetCols() int {
	return len(t.Header)
}

// GetColIndex return col name index
func (t *BaseTable) GetColIndex(colName string) (colIndex int) {
	colIndex = -1 // means not exists.
	for i := 0; i < t.GetCols(); i++ {
		if colName == t.Header[i] {
			colIndex = i
		}
	}
	return
}

// GetStringByColName returns col name value.
func (t *BaseTable) GetString(colName string) (value string) {
	col := t.GetColIndex(colName)
	if col == -1 {
		return ""
	}
	return t.GetStringByIndex(col)
}

// GetInt returns column int value
func (t *BaseTable) GetInt(colName string) int {
	return mcore.NewString(t.GetString(colName)).ToIntNoError()
}

// GetFloat returns column float value
func (t *BaseTable) GetFloat(colName string) float64 {
	return mcore.NewString(t.GetString(colName)).ToFloat64NoError()
}

// GetColNames return table col names
func (t *BaseTable) GetColNames() []string {
	return t.Header
}

// IsHasColumn checks columnName exists
func (t *BaseTable) IsHasColumn(columnName string) bool {
	return mcore.String(columnName).IsInArrayIgnoreCase(t.GetColNames())
}

// Demo demos
func (t *BaseTable) Demo() {
	DemoSimpleTable(t)
}

// GetRowColString Random Access
// TODO not implement yet.
func (t *BaseTable) GetRowColString(row int, col int) (result string) {
	fmt.Printf("GetRowColString not implement\n")
	return
}
