package appconf

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// define conf hubs
var conf = make(map[string]string)

var initLoad = false

func init() {
	LoadConf("app.conf", "confs/app.conf", "/rundata/confs/app.conf")
	initLoad = true
}

// LoadConf load app conf
func LoadConf(section string, files ...string) error {
	for _, f := range files {
		ioutil.ReadFile(f)
	}
	return nil
}

func fmtKey(key string) string {
	key = strings.TrimSpace(key)
	key = strings.ToUpper(key)
	return key
}

// GetString get config string.
func GetString(key string, dv string) string {
	v, e := conf[fmtKey(key)]
	if !e {
		return dv
	}
	return strings.TrimSpace(v)
}

// GetBool return bool
func GetBool(key string, dv bool) bool {
	v, e := conf[fmtKey(key)]
	if !e {
		return dv
	}
	v = strings.TrimSpace(v)
	v = strings.ToUpper(v)
	if v == "YES" || v == "1" || v == "T" || v == "TRUE" || v == "Y" {
		return true
	}
	return false
}

// GetInt return key int
func GetInt(key string, dv int) int {
	v, e := conf[fmtKey(key)]
	if !e {
		return dv
	}
	v = strings.TrimSpace(v)
	if iv, err := strconv.Atoi(v); err == nil {
		return iv
	}
	return 0
}
