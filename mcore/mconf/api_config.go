package mconf

// Combine them
// Any Config must impl this interface
type Config interface {
	ContainChecker
	KeyValuePutter
	StringDefaultGetter
	StringGetter
	BoolGetter
	IntGetter
}

type ContainChecker interface {
	IsContain(key string) bool
}

type KeyValuePutter interface {
	Put(key, value string)
}

type StringDefaultGetter interface {
	GetStringWithDefault(key string, dv string) string
}

type StringGetter interface {
	GetString(key string) string
}

type BoolGetter interface {
	GetBool(key string) bool
}

type IntGetter interface {
	GetInt(key string) int
}
