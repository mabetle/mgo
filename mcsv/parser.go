package mcsv

import(
	"encoding/csv"
	"strings"
	"io/ioutil"
	"fmt"
)

var(
	row []string
	c *CSV
	rowCussor int
)


func Init(){
	rowCussor = 0
}

// return  *CSV
func NewCSV(file string) *CSV{
	c = &CSV{}

	fileContent, err := ioutil.ReadFile(file)

	if err != nil{
		panic(err)
	}

	csvReader := csv.NewReader(strings.NewReader(string(fileContent)))

	headerRow, errRead := csvReader.Read()

	if errRead != nil{
		panic(errRead)
	}

	c.headerRow=headerRow


	data, errReadAll :=csvReader.ReadAll()

	if errReadAll != nil {
		panic(errReadAll)
	}

	c.data=data

	return c
}

// return CSV HeaderRow
func (c *CSV) GetHeaderRow()[]string{
	return c.headerRow
}

// return CSV Data
func (c *CSV) GetData()[][]string{
	return c.data
}

// return CSV Rows
func (c *CSV) GetRows() int {
	return len(c.data)
}

// return CSV Columns
func (c *CSV) GetColumns() int {
	return len(c.headerRow)
}

// parameter row
// return CSV Row
func (c *CSV) GetRow(row int)[]string{
	if row>c.GetRows() {
		return nil
	}
	return c.GetData()[row]
}


// return CSV row and column value
func (c *CSV)GetString(row int, column int)string{

	if row > c.GetRows(){
		return ""
	}

	if column > c.GetColumns(){
		return ""
	}

	return c.GetRow(row)[column]
}

// show CSV Content
func (c *CSV)ShowContent(){
	fmt.Println(c.GetHeaderRow())

	rows := c.GetRows()

	for row := 0 ;row < rows; row++{
		fmt.Println(c.GetRow(row))
	}
}

func (c *CSV)Next() bool{
	rowCussor++
	if rowCussor>c.GetRows()-1{
		return false
	}
	return true
}

func (c *CSV)GetStringByName(columnName string) string {
	return c.GetStringByColumnIndex(c.GetColumnNameIndex(columnName))
}

func (c *CSV)GetStringByColumnIndex(column int) string{
	if rowCussor==0 {
		return ""
	}
	if column==-1{
		return ""
	}
	return c.GetString(rowCussor,column)
}

func (c *CSV)GetColumnNameIndex(columnName string) int{
	h:=c.GetHeaderRow()
	for i:=0;i<len(h);i++{
		if h[i]==columnName {
			return i
		}
	}
	return -1
}

