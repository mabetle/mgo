package msdb

// SimpleTable defines SimpleTable API
type SimpleTable interface {
	GetColNames() []string
	GetRows() int
	GetRowIndex() int
	GetCols() int
	Next() bool
	StringGetter
	GetString(colName string) string
	GetInt(colName string) int
	GetFloat(colName string) float64
	IsHasColumn(columName string) bool
}

// StringGetter define GetString API
type StringGetter interface {
	GetStringByIndex(col int) string //default by col index
}
