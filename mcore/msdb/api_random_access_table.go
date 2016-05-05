package msdb

// RandomAccessTable
// Can GetString by Row and Col index
type RandomAccessTable interface{
	GetColNames()[]string
	GetRows() int
	GetCols() int
	GetRowColString(row int,col int)string
}

