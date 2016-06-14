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
)

var (
	VersionMain  = "1"
	VersionMinus = "0"

	VendorName = VendorMabetle
	VendorURL  = VendorMabetleUrl

	LicenseTo = "demo"

	AccountSchema = "web_common"
)

var (
	KeyLicenseTo  = "license.to"
	KeyVendorName = "vendor.name"
)
