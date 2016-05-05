package msdb

type Cusor struct{
	MaxIndex	int
	RowIndex	int
}

// ideas from jdbc
func (c *Cusor)Next()bool{
	c.RowIndex++
	if c.RowIndex<c.MaxIndex+1{
		return true
	}
	return false
}


func (c *Cusor)First(){
	c.RowIndex = 0
}

func (c *Cusor)Last(){
	c.RowIndex = c.MaxIndex
}

func (c *Cusor)Move(rowIndex int){
	if rowIndex > c.MaxIndex{
		rowIndex = c.MaxIndex
	}
	c.RowIndex = rowIndex
}


