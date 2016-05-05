package mcore

// Validator API

// Validator
type Validetor interface {
	Validate(v interface{}) bool
}
