package mcrypt

// return as input.
type Plain struct {
}

func NewPlain() *Plain {
	return &Plain{}
}

// implements Encoder interface.
func (s Plain) Encode(old string) string {
	return old
}
