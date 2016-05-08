package mstat

// Sorter .
type Sorter interface {
	Sort(values []interface{}) []interface{}
}

// DescSorter .
type DescSorter interface {
	DescSort(values []interface{}) []interface{}
}
