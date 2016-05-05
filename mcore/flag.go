package mcore

import (
	"flag"
)

// FIXME not work properly.
func GetFlagString(name string)string{
	if !flag.Parsed(){
		flag.Parse()
	}
	if v:=flag.Lookup(name); v== nil{
		return ""
	}else{
		return v.Value.String()
	}
}

