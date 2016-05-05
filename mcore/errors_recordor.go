package mcore

import (
	"fmt"
	"strings"
)

// Errors record errors
type Errors struct {
	errs []string
}

// NewErrors creates new Errors
func NewErrors() *Errors {
	return &Errors{errs: []string{}}
}

// Record records error
func (r *Errors) Record(err error) {
	if err != nil {
		r.errs = append(r.errs, fmt.Sprint(err))
	}
}

// HasError check has errors
func (r *Errors) HasError() bool {
	return len(r.errs) > 0
}

// Error return errors merged error message.
func (r *Errors) Error() error {
	if r.HasError() {
		return fmt.Errorf("%s", strings.Join(r.errs, ","))
	}
	return nil
}

// Len return errors nums
func (r Errors) Len() int {
	return len(r.errs)
}

// Print prints errors
func (r *Errors) Print() {
	if !r.HasError() {
		fmt.Printf("No error\n")
	}
	for i, err := range r.errs {
		fmt.Printf("Error %d: %s\n", i+1, err)
	}
}
