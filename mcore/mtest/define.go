package mtest

import "fmt"

type InputExpect struct {
	Input  interface{}
	Expect interface{}
	Msg    string
}

type InputExpects struct {
	rows []InputExpect
}

func NewInputExpects() *InputExpects {
	rows := []InputExpect{}
	return &InputExpects{rows: rows}
}

func (s *InputExpects) Put(input, expect interface{}, msgs ...interface{}) *InputExpects {
	msg := fmt.Sprint(msgs...)
	row := InputExpect{Input: input, Expect: expect, Msg: msg}
	s.rows = append(s.rows, row)
	return s
}

func (s *InputExpects) Rows() []InputExpect {
	return s.rows
}
