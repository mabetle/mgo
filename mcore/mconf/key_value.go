package mconf

import (
	"github.com/mabetle/mcore"
)

// KeyValueConfig implements Config interface.
type KeyValueConfig struct {
	mcore.StringKeyValueMap // store key and value.
	KeyValueLoader       //load keys and values.
}

// KeyValueLoader
type KeyValueLoader interface {
	LoadKeyValue() mcore.StringKeyValueMap
}

// NewConfig
func NewConfig(loader KeyValueLoader) Config {
	c := &KeyValueConfig{}
	c.KeyValueLoader = loader
	c.StringKeyValueMap = c.LoadKeyValue()
	return c
}
