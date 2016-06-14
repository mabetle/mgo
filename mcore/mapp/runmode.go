package mapp

import "github.com/mabetle/mgo/mcore"

func GetRunMode() string {
	if mcore.IsHasDevArg() {
		return "dev"
	} else if mcore.IsHasTestArg() {
		return "test"
	}
	return "prod"
}

func IsDevMode() bool {
	return GetRunMode() == "dev"
}
