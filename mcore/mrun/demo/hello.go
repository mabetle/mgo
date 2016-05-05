package demo

import (
	"fmt"
)

func init(){
	AddFunc(Hello, "hello")
}

func Hello(){
	fmt.Printf("%s\n", "Hello")
}

