package mconf

import (
	"github.com/mabetle/mcore"
)

type StringGetter interface {
	GetString(key string, defaultValue string) string
}

type BaseConfig struct {
	StringGetter
}

// GetInt
func (c BaseConfig) GetInt(key string) int {
	return mcore.NewString(c.GetString(key, "0")).ToIntNoError()
}

// GetBool
func (c BaseConfig) GetBool(key string) bool {
	return mcore.NewString(c.GetString(key, "F")).ToBool()
}
