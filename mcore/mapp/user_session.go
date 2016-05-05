package mapp

import (
	"fmt"
	"strings"
)

var userSession = make(map[string]interface{})

func userKey(userName, key string) string {
	return fmt.Sprintf("%s-%s", userName, key)
}

func PutUserSession(userName string, key string, value interface{}) {
	userSession[userKey(userName, key)] = value
}

func IsContainUserSession(userName, key string) bool {
	_, ok := userSession[userKey(userName, key)]
	return ok
}

func GetUserSession(userName, key string) interface{} {
	return userSession[userKey(userName, key)]
}

func GetUserSessions(userName string) []string {
	as := []string{}
	prefix := userName + "-"
	for k, _ := range userSession {
		if strings.HasPrefix(k, prefix) {
			as = append(as, strings.TrimPrefix(k, prefix))
		}
	}
	return as
}

func PutUserSessionKeyValues(kvs ...KeyValue) {
	for _, kv := range kvs {
		userSession[kv.Key()] = kv.Value()
	}
}
