package mcrypt

type Encoder interface {
	Encode(old string) string
}
