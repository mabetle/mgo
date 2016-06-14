package mcore

type KeyValueMap map[string]interface{}

func NewKeyValueMap() KeyValueMap {
	return make(map[string]interface{})
}

func (c KeyValueMap) GetKeys() []string {
	return GetMapKeys(c)
}

func (c KeyValueMap) IsHasKey(key string) bool {
	return IsMapHasKey(c, key)
}
