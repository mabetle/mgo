package mcore

import (
	"fmt"
	"strings"
)

// Join joins args with sep
func Join(sep string, args ...interface{}) string {
	return JoinArray(sep, args)
}

// JoinArray joins array args with sep
func JoinArray(sep string, args []interface{}) string {
	n := len(args)
	if n < 1 {
		return ""
	}
	vs := make([]string, n)
	for i, arg := range args {
		vs[i] = fmt.Sprint(arg)
	}
	return strings.Join(vs, sep)
}

// GetMappedName return mapped name with mapRules
func GetMappedName(name string, mapRules string) string {
	if mapRules == "" {
		return name
	}
	kvLines := strings.Split(mapRules, ",")
	for _, kvLine := range kvLines {
		kvLine = strings.TrimSpace(kvLine)
		if strings.HasPrefix(kvLine, "#") {
			//skip comment line
			continue
		}
		kvs := strings.Split(kvLine, "=")

		if len(kvs) != 2 {
			// not kv line
			continue
		}
		k := kvs[0]
		v := kvs[1]
		// found
		if k == name {
			return v
		}
	}
	// not found, means not define
	return name
}
