package mcore

import (
	"fmt"
)

type KeyValue struct {
	Key   string
	Value interface{}
}

// String
func (kv KeyValue) String() string {
	if v, b := kv.Value.(string); b {
		return v
	}
	return fmt.Sprint(kv.Value)
}

func (kv KeyValue) Int(dv int) int {
	if v, b := kv.Value.(int); b {
		return v
	}
	n, err := String(kv.String()).ToInt()
	if err != nil {
		return dv
	}
	return n
}

func (kv KeyValue) Bool() bool {
	if v, b := kv.Value.(bool); b {
		return v
	}
	s := kv.String()
	return String(s).ToBool()
}

func (kv KeyValue) Float64(dv float64) float64 {
	if v, b := kv.Value.(float64); b {
		return v
	}
	fv, err := String(kv.String()).ToFloat64()
	if err != nil {
		return dv
	}
	return fv
}
