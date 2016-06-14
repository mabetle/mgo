package mcore

type Email string

// Validate implements Validator
func (e Email) Validate() (b bool) {
	b = String(e).IsEmail()
	return
}
