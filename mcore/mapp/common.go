package mapp

import (
	"github.com/mabetle/mlog"
)

var (
	logger = mlog.GetLogger("github.com/mabetle/mcore/mapp")
)

const (
	V_VENDOR_MABETLE     = "Mabetle"
	V_VENDOR_MABETLE_URL = "http://www.mabetle.com"
	MODE_DEV             = "dev"
)

var (
	VersionMain  = "1"
	VersionMinus = "0"

	VendorName = V_VENDOR_MABETLE
	VendorURL  = V_VENDOR_MABETLE_URL

	LicenseTo = "DEMO"
	RunMode   = "dev"

	AccountSchema = "web_common"
)

var (
	KEY_LICENSE_TO  = "license.to"
	KEY_VENDOR_NAME = "vendor.name"
)

type License struct {
	To string
}

type Vendor struct {
	Name string
	URL  string
}

type Version struct {
	Main  string
	Minus string
}
