package mcsv


type RandomAccessTable struct{

}

// GetRandomAccessTable
func NewRandomAccessTable(file string) *RandomAccessTable{
	return &RandomAccessTable{}
}

func (table RandomAccessTable) GetRows()int{
	return 5
}

func (table RandomAccessTable) GetCols()int{
	return 5
}


// has out of index error
func (table RandomAccessTable) GetString(row int,col int)string{
	return "demo"
}

