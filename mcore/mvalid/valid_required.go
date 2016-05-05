package mvalid

import (
	"fmt"
)

type RequiredValidator struct {
	Value interface{}
	Field string
	Validator
}

func NewRequiredValidator(v interface{}, field string) *RequiredValidator {
	return &RequiredValidator{Value: v, Field: field}
}

// Validate implement Validator.
func (vd *RequiredValidator) Validate() (bool, string) {
	// valid a value.
	if vd.Field == "" {
		v := fmt.Sprintf("%v", vd.Value)
		if v == "" {
			err := ""
			return false, err
		}
	}
	// field not null.

	return true, ""
}
