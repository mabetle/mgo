package mcore

import (
	"strconv"
	"errors"
)
var ErrInvalidNumber = errors.New("Invalid number format")

// error
type NumberFormatError struct{

}

// if cannot change, return 0
func StrToInt(v string)(r int, err error){
	r, err=strconv.Atoi(v)
	if nil!= err{
		err = ErrInvalidNumber
		return
	}
	return
}

// StrToFloat64
func StrToFloat64(v string)(r float64, err error){
	r, err = strconv.ParseFloat(v, 64)
	if nil != err{
		err = ErrInvalidNumber
		return
	}
	return
}

func StrToInt64(v string)(r int64, err error){
	r, err = strconv.ParseInt(v, 10, 0)
	if nil != err{
		err = ErrInvalidNumber
		return
	}
	return
}


