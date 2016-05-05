package mcore

import (
	"log"
)

type StringKeyValueMap map[string]string

func NewStringKeyValueMap() StringKeyValueMap {
	return make(map[string]string)
}

// put key value
func (c StringKeyValueMap) Put(key, value string) {
	c[key] = value
}

func (c StringKeyValueMap) IsContain(key string) bool {
	_, ok := c[key]
	return ok
}

func (c StringKeyValueMap) GetString(key string) string {
	return c.GetStringWithDefault(key, "")
}

func (c StringKeyValueMap) GetStringWithDefault(key, defaultValue string) string {
	if c.IsContain(key) {
		return c[key]
	}
	log.Printf("Error: not contains key: %s", key)
	return defaultValue
}

func (c StringKeyValueMap) GetInt(key string) int {
	return String(c.GetString(key)).ToIntNoError()
}

func (c StringKeyValueMap) GetBool(key string) bool {
	return String(c.GetString(key)).ToBool()
}
