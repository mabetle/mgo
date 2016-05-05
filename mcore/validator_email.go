package mcore

type Email string

// implement Validate
func (e Email) Validate() (b bool) {
	b = String(e).IsEmail()
	return
}
