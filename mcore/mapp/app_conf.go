package mapp

import (
	"log"

	"github.com/mabetle/mgo/mcore/mconf"
	"github.com/mabetle/mgo/mcore/mconf/ini"
)

// App Config
// config can hold in files and db

var appConfLocations = []string{}

// hold configs
var appConf mconf.Config

// you can change this
var AppName string

var appConfLoaded = false

// SetAppConfLocations load config from locations
func SetAppConfLocations(locations ...string) {
	appConfLocations = locations
}

// LoadAppConf
func LoadAppConf() {
	if len(appConfLocations) == 0 {
		panic("no app conf files")
	}
	log.Printf("Load app conf: %v", appConfLocations)
	appConf = ini.NewConfig(appConfLocations...)
	appConfLoaded = true
}

var scanDirs = []string{
	".",
	"conf",
	"/rundata/conf",
}

// ScanConfs, lazy to set config locations
func ScanConfs() {

}

func GetAppName() string {
	if AppName == "" {
		AppName = GetAppConfStringWithDefault("app.name", "dbc")
	}
	return AppName
}

// GetAppConfStringWithDefault is main config entry point.
func GetAppConfStringWithDefault(key string, dv string) string {
	if !appConfLoaded {
		LoadAppConf()
	}
	return appConf.GetStringWithDefault(key, dv)
}

func GetAppConfString(key string) string {
	return GetAppConfStringWithDefault(key, "")
}

// VendorName
func GetVendorName() string {
	return GetAppConfString(KeyVendorName)
}
