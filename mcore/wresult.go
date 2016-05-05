package mcore

import (
	"fmt"
)

// Results define results
type Results struct {
	Errors []error  //strore error
	Msgs   []string //store not error message.
}

// NewResults creates Results instanse.
func NewResults() *Results {
	return &Results{
		Errors: []error{},
		Msgs:   []string{},
	}
}

// RecordMsg records message
func (r *Results) RecordMsg(msg ...interface{}) {
	r.Msgs = append(r.Msgs, fmt.Sprint(msg...))
}

// RecordError records error
func (r *Results) RecordError(err error) {
	r.Errors = append(r.Errors, err)
}

// RecordErrorMsg record error message
func (r *Results) RecordErrorMsg(errMsg ...interface{}) {
	msg := fmt.Sprint(errMsg...)
	r.RecordError(fmt.Errorf(msg))
}

// RecordfMsg record message with format.
func (r *Results) RecordfMsg(format string, args ...interface{}) {
	r.RecordfMsg(fmt.Sprintf(format, args...))
}

// RecordfError record error with format.
func (r *Results) RecordfError(format string, args ...interface{}) {
	r.RecordErrorMsg(fmt.Sprintf(format, args...))
}

// IsHasError return has error or not
func (r *Results) IsHasError() bool {
	return len(r.Errors) > 0
}

// IsHasMsg returns has message or not.
func (r *Results) IsHasMsg() bool {
	return len(r.Msgs) > 0
}

// PrintErrors prints errors
func (r *Results) PrintErrors() {
	if !r.IsHasError() {
		fmt.Println("No errors")
		return
	}
	for i, err := range r.Errors {
		fmt.Printf("Error %d: %v\n", i+1, err)
	}
}

// PrintMsgs prints messages.
func (r *Results) PrintMsgs() {
	if !r.IsHasMsg() {
		fmt.Printf("No messages.\n")
	}
	for i, v := range r.Msgs {
		fmt.Printf("Msg %d: %v\n", i+1, v)
	}
}

// Print prints errors and messages.
func (r *Results) Print() {
	r.PrintErrors()
	r.PrintMsgs()
}
