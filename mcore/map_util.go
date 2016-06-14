package mcore

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
