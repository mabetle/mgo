package mcsv

//implements for sdb.SimpleTable
// a valid table csv file must have head in first line
type SimpleTable struct{
	file string
	head []string
	body [][]string
	rowIndex int
}

func GetSimpleTable(file string)*SimpleTable{
	return &SimpleTable{}
}

//implements for sdb.SimpleTable
func (t *SimpleTable) GetRows() int{
	return 5
}

func (t *SimpleTable) GetCols() int{
	return 5
}

func (t *SimpleTable) Next() bool{
	t.rowIndex++

	if t.rowIndex<t.GetRows(){
		return true
	}
	return false
}

func (t *SimpleTable)GetString(col int)string{
	return "demo"
}

func (t *SimpleTable)GetStringByColName(colName string)string{
	return "demo"
}

