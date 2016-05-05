package mrun


import (
	"testing"
)

func Hello(){
	//"hello"
}

func TestRunFuc(t *testing.T){
	AddFunc(Hello, "hello")
	RunAllFuncs()
}


