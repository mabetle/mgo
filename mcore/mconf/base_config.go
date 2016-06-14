package mconf

import (
	"github.com/mabetle/mgo/mcore"
)

type BaseConfig struct {
	StringDefaultGetter
}

// GetInt
func (c BaseConfig) GetInt(key string) int {
	return mcore.NewString(c.GetStringWithDefault(key, "0")).ToIntNoError()
}

// GetBool
func (c BaseConfig) GetBool(key string) bool {
	return mcore.NewString(c.GetStringWithDefault(key, "F")).ToBool()
}
