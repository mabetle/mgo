package mcore

import (
	"fmt"
)

// Error define Error
type Error struct {
	err error
}

// NewError creates Error
func NewError() *Error {
	return &Error{err: nil}
}

// NewErrError creates Error with error arg
func NewErrError(err error) *Error {
	return &Error{err: err}
}

// PutError put error to Error
func (e *Error) PutError(err error) *Error {
	e.err = err
	return e
}

// Erorr return Error error
func (e *Error) Error() error {
	return e.err
}

// Print print Error
func (e *Error) Print() {
	fmt.Printf("%v\n", e.err)
}

// IsReturnHasError check returns
func IsReturnHasError(args ...interface{}) (r bool) {
	nargs := len(args)
	if nargs < 1 {
		return
	}
	larg := args[nargs-1]
	if IsError(larg) && larg != nil {
		r = true
		return
	}
	return
}

// IsError check error.
func IsError(v interface{}) (r bool) {
	if _, ok := v.(error); ok {
		r = true
	}
	return
}
