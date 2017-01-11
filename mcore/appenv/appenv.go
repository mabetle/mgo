package appenv

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// define conf hubs
var conf = make(map[string]string)

var initLoad = false

func init() {
	LoadConf("app.conf", "confs/app.conf", "/rundata/confs/app.conf")
	LoadEnv()
	initLoad = true
}

func LoadEnv() {
	// env format key=value
	for _, line := range os.Environ() {
		PutEnvLine(line)
	}
}

func PutEnv(key, value string) {
	conf[fmtKey(key)] = value
}

func PutEnvLine(line string) {
	// not a key=value line
	if !strings.Contains(line, "=") {
		return
	}
	kvs := strings.Split(line, "=")
	// make sure kvs array valid
	if len(kvs) != 2 {
		return
	}
	key := strings.TrimSpace(kvs[0])
	value := strings.TrimSpace(kvs[1])
	PutEnv(key, value)
}

func LoadSectionConf(section string, files ...string) {

}

// LoadConf load app conf
func LoadConf(files ...string) {
	for _, f := range files {
		bs, err := ioutil.ReadFile(f)
		//file not found
		if err != nil {
			continue
		}
		log.Printf("load conf file: %s", f)
		// parse text
		PutEnvText(string(bs))
	}
}

func PutEnvText(text string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		PutEnvLine(line)
	}
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

func PrintConfs() {
	fmt.Printf("===Confs===\n")
	for k, v := range conf {
		fmt.Printf("%s:%s\n", k, v)
	}
}
