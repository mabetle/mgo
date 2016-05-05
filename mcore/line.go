package mcore

import (
	"fmt"
)

// GetFileConfigValue returns config value from config files
func GetFileConfigValue(location string, key string, sep string) (string, error) {
	lines, err := ReadFileLines(location)
	if err != nil {
		return "", err
	}
	return GetConfigValue(lines, key, sep)
}

// GetConfigValue returns config value from lines.
func GetConfigValue(lines []string, key string, sep string) (string, error) {
	for _, kv := range GetKeyValueArray(lines, sep) {
		if NewString(kv.Key).IsEqualIgnoreCase(key) {
			return kv.String(), nil
		}
	}
	// not found
	return "", fmt.Errorf("not found value for key: %s", key)
}

// GetKeyValueArrayFromFile
func GetKeyValueArrayFromFile(location string, sep string) ([]KeyValue, error) {
	kvs := []KeyValue{}
	lines, err := ReadFileLines(location)
	if err != nil {
		return kvs, err
	}
	return GetKeyValueArray(lines, sep), nil
}

// GetKeyValueArray
func GetKeyValueArray(lines []string, sep string) []KeyValue {
	kvs := []KeyValue{}
	for _, line := range lines {
		lineS := NewString(line).TrimSpace()
		if lineS.IsStartsIgnoreCase("#", "//") || lineS.IsBlank() {
			continue
		}
		var kvp []string
		kvp = lineS.Split(sep)
		// bad key value
		if len(kvp) < 2 {
			continue
		}
		k := kvp[0]
		v := kvp[1]
		kv := KeyValue{Key: k, Value: v}
		kvs = append(kvs, kv)
	}
	return kvs
}
