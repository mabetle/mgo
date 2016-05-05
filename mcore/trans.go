package mcore

import (
	"fmt"
	"strconv"
)

type Any struct {
	Value interface{}
}

func NewAny(value interface{}) Any {
	return Any{Value: value}
}

func (v Any) String() string {
	return fmt.Sprint("%v", v.Value)
}

func (v Any) Int() (i int64, err error) {
	s := v.String()
	return strconv.ParseInt(s, 10, 64)
}

func (v Any) Float() (float64, error) {
	s := v.String()
	return strconv.ParseFloat(s, 64)
}
