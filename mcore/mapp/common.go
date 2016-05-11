package mapp

import (
	"github.com/mabetle/mgo/mlog"
)

var (
	logger = mlog.GetLogger("github.com/mabetle/mgo/mcore/mapp")
)

const (
	VendorMabetle    = "Mabetle"
	VendorMabetleUrl = "http://www.mabetle.com"
	ModeDev          = "dev"
)

var (
	VersionMain  = "1"
	VersionMinus = "0"

	VendorName = VendorMabetle
	VendorURL  = VendorMabetleUrl

	LicenseTo = "demo"
	RunMode   = "dev"

	AccountSchema = "web_common"
)

var (
	KeyLicenseTo  = "license.to"
	KeyVendorName = "vendor.name"
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
