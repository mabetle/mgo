package mvalid

// Validator
type Validator interface {
	// run validate
	// Validator provide Validate method.
	Validate() (bool, string)
}
