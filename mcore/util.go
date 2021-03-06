package mcore

import (
	"fmt"
	"strings"
)

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

func CheckError(err error, okMsgs ... interface{}) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if len(okMsgs) != 0 {
		fmt.Printf("Success:%s\n", fmt.Sprint(okMsgs ...))
	}
}
