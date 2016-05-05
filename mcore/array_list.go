package mcore

import (
	"fmt"
	"reflect"
)

// List interface.
type List interface {
	Next() bool
	Len() int
	Put(values ...interface{})
	Value() interface{}
}

// ArrayList implements List interface.
type ArrayList struct {
	cursor int // cursor -1 means not move. cursor is values index.
	values []interface{}
	typ    reflect.Type
}

// NewArrayList create a ArrayList instance.
// model
func NewArrayList(model interface{}) *ArrayList {
	l := &ArrayList{}
	v := []interface{}{}
	l.cursor = -1
	l.values = v
	l.typ = reflect.TypeOf(model)
	return l
}

// Len return length of value
func (s *ArrayList) Len() int {
	return len(s.values)
}

// Next for loop
func (s *ArrayList) Next() bool {
	if s.Len() == 0 {
		return false
	}
	s.cursor++
	if s.cursor > s.Len()-1 {
		s.cursor = s.Len() - 1
		return false
	}
	return true
}

// Put put value
func (s *ArrayList) Put(values ...interface{}) {
	for _, v := range values {
		s.values = append(s.values, v)
	}
}

// Value returns cursor value
func (s *ArrayList) Value() interface{} {
	return s.values[s.cursor]
}

// Type return reflect type
func (s *ArrayList) Type() reflect.Type {
	return s.typ
}

// Remove remove
func (s *ArrayList) Remove(value interface{}) {
	tmp := []interface{}{}
	for _, e := range s.values {
		if e == value {
			continue
		}
		tmp = append(tmp, e)
	}
	s.values = tmp
}

// IsContain is contain value
func (s *ArrayList) IsContain(value interface{}) bool {
	for _, e := range s.values {
		if e == value {
			return true
		}
	}
	return false
}

// Print out put
func (s *ArrayList) Print() {
	fmt.Println("Data in ArrayList:")
	for s.Next() {
		fmt.Printf("%d:%v\n", s.cursor, s.Value())
	}
}
