package mcmd

import (
	"os"

	"github.com/mabetle/mgo/mcore"
)

// return which app is run.
func AppName() string {
	return mcore.String(os.Args[0]).ReplaceAll("\\", "/").SepEnd("/").String()
}

// promote user input args to run application.
func ReadNotBlankLineWithMsg(msg string) string {
	return mcore.ReadNotBlankLineWithMsg(msg)
}
