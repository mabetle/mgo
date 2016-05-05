package mconf

// Combine them
// Any Config must impl this interface
type Config interface {
	IsContain(key string) bool
	GetStringWithDefault(key string, defaultValue string) string

	GetString(key string) string
	// extends GetString()
	GetInt(key string) int
	GetBool(key string) bool
	Put(key, value string)
}
