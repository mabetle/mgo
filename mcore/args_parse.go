package mcore

import (
	"fmt"
	"strings"
)

// Args wraps args
type Args []string

// ParseStringToArgs line to args
// FIXME not work as expect
func ParseStringToArgs(s string) []string {
	r := []string{}
	s = strings.TrimSpace(s)
	if s == "" {
		return r
	}
	args := strings.Split(s, " ")
	for _, arg := range args {
		arg = strings.TrimSpace(arg)
		arg = strings.Trim(arg, `"`) // delete wrap "
		if arg != "" {
			r = append(r, arg)
		}
	}
	return r
}

// NewArgsFromString create Args
func NewArgsFromString(s string) Args {
	return Args(ParseStringToArgs(s))
}

// NewArgs creates Args
func NewArgs(args []string) Args {
	return Args(args)
}

// IsHasFlag returns is has flag
func (a Args) IsHasFlag(flag string) bool {
	return String(flag).IsInArray(a)
}

// ParseString parse string
func (a Args) ParseString(flag string) (r string) {

	return
}

// ParseInt parse int
func (a Args) ParseInt(flag string) (r int) {

	return
}

// NArgs number of args
func (a Args) NArgs() int {
	return len(a)
}

// NFlags number of flags
func (a Args) NFlags() (r int) {

	return
}

// VArgs value of args indexed
func (a Args) VArgs(index int) string {
	return a[index]
}

func GetParseArgs(args ...string) []string {
	parseArgs := []string{}
	for _, arg := range args {
		// muliti lines args
		if strings.Contains(arg, "\n") {
			nas := strings.Split(arg, "\n")
			parseArgs = append(parseArgs, GetParseArgs(nas...)...)
		}
		// seperate by blank
		if strings.Contains(arg, " ") {
			nas := strings.Split(arg, " ")
			parseArgs = append(parseArgs, GetParseArgs(nas...)...)
		}
		parseArgs = append(parseArgs, arg)
	}
	return parseArgs
}

// GetArgString arg format: a=b
func GetArgString(name string, defaultValue string, args ...string) string {
	// process Args
	for _, a := range GetParseArgs(args...) {
		arg := GetString(a)
		//arg format: a=b
		arg = strings.TrimSpace(arg)
		kv := strings.Split(arg, "=")
		value := ""
		if len(kv) == 2 {
			value = strings.TrimSpace(kv[1])
		}
		key := strings.TrimSpace(kv[0])
		if NewString(key).IsEqualIgnoreCase(name) {
			return value
		}
	}
	return defaultValue
}

// GetArgBool return bool value
func GetArgBool(name string, defaultValue bool, args ...string) bool {
	dVS := "0"
	if defaultValue {
		dVS = "1"
	}
	bStr := GetArgString(name, dVS, args...)
	return NewString(bStr).ToBool()
}

// GetArgInt returns int value
func GetArgInt(name string, defaultValue int, args ...string) int {
	iStr := GetArgString(name, "", args...)
	if iStr == "" {
		return defaultValue
	}
	n, err := NewString(iStr).ToInt()
	if err != nil {
		return defaultValue
	}
	return n
}

// IsArgExists is exists arg
func IsArgExists(name string, args ...string) bool {
	for _, a := range GetParseArgs(args...) {
		arg := GetString(a)
		arg = strings.TrimSpace(arg)
		//arg format: a=b
		key := strings.Split(arg, "=")[0]
		key = strings.TrimSpace(key)
		if NewString(key).IsEqualIgnoreCase(name) {
			return true
		}
	}
	return false
}

// JoinArgs join args
func JoinArgs(renderArgs map[string]interface{}, args ...string) []string {
	r := []string{}
	for k, v := range renderArgs {
		r = append(r, fmt.Sprintf("%s=%v", k, v))
	}
	r = append(r, GetParseArgs(args...)...)
	return r
}
