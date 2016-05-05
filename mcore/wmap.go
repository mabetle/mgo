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

func GetMapKeys(m map[string]interface{}) (r []string) {
	for k, _ := range m {
		r = append(r, k)
	}
	return
}

func IsMapHasKey(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}
