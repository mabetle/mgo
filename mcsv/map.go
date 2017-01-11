package mcsv

import (
	"io/ioutil"
	"strings"
)

func ReadMap(location string, keyCol, valueCol int) map[string]string {
	vs := make(map[string]string)
	bs, err := ioutil.ReadFile(location)
	if err != nil {
		return vs
	}
	lines := strings.Split(string(bs), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		items := strings.Split(line, ",")
		n := len(items)
		key := ""
		value := ""
		// keyCol exceed range
		if keyCol > n {
			continue
		}
		key = strings.TrimSpace(items[keyCol])
		if key == "" {
			continue
		}
		if valueCol < n {
			value = strings.TrimSpace(items[valueCol])
		}
		// put value
		vs[key] = value
	}
	return vs
}
