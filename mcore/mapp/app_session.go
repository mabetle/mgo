package mapp

import (
	"io/ioutil"
	"strings"
)

type KeyValue interface {
	Key() string
	Value() interface{}
}

var appSession = make(map[string]interface{})

func PutAppSession(key string, value interface{}) {
	appSession[key] = value
}

func IsContainAppSession(key string) bool {
	_, ok := appSession[key]
	return ok
}

func GetAppSession(key string) interface{} {
	return appSession[key]
}

// LoadAppSessionFile read configs
func LoadAppSessionFile(file string) error {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	text := string(bs)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		kv := strings.Split(line, "=")
		if len(kv) == 1 {
			continue
		}
		key := strings.TrimSpace(kv[0])
		value := strings.Join(kv[1:], "=")
		value = strings.TrimSpace(value)
		PutAppSession(key, value)
	}
	return nil
}

func GetAppSessions() []string {
	as := []string{}
	for k, _ := range appSession {
		as = append(as, k)
	}
	return as
}

func PutAppSessionKeyValues(kvs ...KeyValue) {
	for _, kv := range kvs {
		PutAppSession(kv.Key(), kv.Value())
	}
}
