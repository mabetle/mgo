package mcore

import (
	"fmt"
)

// ErrorType define
type ErrorType int

const (
	// TYPE_EXCEPTION define
	TYPE_EXCEPTION ErrorType = iota
	// TYPE_RUNTIME_EXCEPTION define
	TYPE_RUNTIME_EXCEPTION
)

// Exception app can not go on. app should check it.
type Exception struct {
	ErrorType
	err error
}

// NewExceptionf create
func NewExceptionf(format string, args ...interface{}) Exception {
	return NewException(fmt.Errorf(format, args...))
}

// NewException create
func NewException(err error) Exception {
	return Exception{ErrorType: TYPE_EXCEPTION, err: err}
}

func (e Exception) Error() string {
	return e.err.Error()
}

// ErrorString error string
func (e Exception) ErrorString() string {
	return e.Error()
}

// RuntimeException depends on apps, not check.
type RuntimeException struct {
	Exception
}

// NewRuntimeExceptionf create
func NewRuntimeExceptionf(format string, args ...interface{}) RuntimeException {
	return NewRuntimeException(fmt.Errorf(format, args...))
}

// NewRuntimeException create
func NewRuntimeException(err error) RuntimeException {
	e := Exception{ErrorType: TYPE_RUNTIME_EXCEPTION, err: err}
	return RuntimeException{Exception: e}
}

// FileNotFoundError define
type FileNotFoundError struct {
	RuntimeException
}

// NewFileNotFoundError create
func NewFileNotFoundError(name string) FileNotFoundError {
	e := NewRuntimeExceptionf("File not found. File: %s ", name)
	return FileNotFoundError{RuntimeException: e}
}

// IsRuntimeException is runtime Exception
func IsRuntimeException(e Exception) bool {
	return e.ErrorType == TYPE_RUNTIME_EXCEPTION
}

// IsException is Exception
func IsException(e Exception) bool {
	return e.ErrorType == TYPE_EXCEPTION
}

// IsRuntimeExceptionType is runtime Exception type
func IsRuntimeExceptionType(err error) bool {
	if _, ok := err.(RuntimeException); ok {
		return true
	}
	return false
}

func IsExceptionType(err error) bool {
	return !IsRuntimeExceptionType(err)
}
