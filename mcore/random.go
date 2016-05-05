package mcore

import (
	"fmt"
	"math/rand"
	"time"
)

// NewRandom generate new random integer number.
func NewRandom() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()
}

//NewRangeRandom generate new random integer number range in 0~n
func NewRangeRandom(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}

// NewRandomPassword
// generate new number password.
func NewRandomPassword() string {
	s := fmt.Sprintf("%v", NewRangeRandom(10000000))
	return s
}
