package margs

import (
	"fmt"
	"strconv"
	"strings"
)

type KeyValue struct {
	key   string
	value interface{}
}

type Args struct {
	kvs []KeyValue
}

// NewArgs
func NewArgs(renderArgs map[string]interface{}, args ...string) *Args {
	kvs := []KeyValue{}
	for k, v := range renderArgs {
		kvs = append(kvs, KeyValue{key: k, value: v})
	}

	for _, arg := range args {
		kv := strings.Split(arg, "=")
		k := kv[0]
		v := ""
		if len(kv) > 1 {
			v = strings.Join(kv[1:], "=")
		}
		kvs = append(kvs, KeyValue{key: k, value: v})
	}
	return &Args{kvs: kvs}
}

func (s *Args) IsHasArg(name string) bool {
	for _, kv := range s.kvs {
		if kv.key == name {
			return true
		}
	}
	return false
}

func (s *Args) GetArray(name string) []interface{} {
	rows := []interface{}{}
	for _, kv := range s.kvs {
		if kv.key == name {
			rows = append(rows, kv.value)
		}
	}
	return rows
}

func (s *Args) GetString(name string, dv string) string {
	for _, kv := range s.kvs {
		if kv.key == name {
			return fmt.Sprint(kv.value)
		}
	}
	return dv
}

func (s *Args) GetInt(name string, dv int) int {
	sv := s.GetString(name, "")
	if sv == "" {
		return dv
	}
	iv, err := strconv.Atoi(sv)
	if err != nil {
		return dv
	}
	return iv
}

func (s *Args) GetInt64(name string, dv int64) int64 {
	sv := s.GetString(name, "")
	if sv == "" {
		return dv
	}
	iv, err := strconv.Atoi(sv)
	if err != nil {
		return dv
	}
	return int64(iv)
}

func (s *Args) GetFloat64(name string, dv float64) float64 {
	sv := s.GetString(name, "")
	if sv == "" {
		return dv
	}
	fv, err := strconv.ParseFloat(sv, 64)
	if err != nil {
		return fv
	}
	return dv
}

func (s *Args) GetBool(name string, dv bool) bool {
	sv := s.GetString(name, "")
	if sv == "" {
		return dv
	}
	sv = strings.ToLower(sv)
	if sv == "true" || sv == "t" || sv == "yes" || sv == "y" || sv == "1" {
		return true
	}
	return false
}

func (s *Args) GetLocale() string {
	return s.GetString("currentLocale", "en_US")
}

func (s *Args) GetInclude() string {
	return s.GetString("include", "")
}

func (s *Args) GetExclude() string {
	return s.GetString("exclude", "")
}

func (s *Args) String() string {
	rows := []string{}
	for _, kv := range s.kvs {
		rows = append(rows, fmt.Sprintf("%s=%v", kv.key, kv.value))
	}
	return strings.Join(rows, "\n")
}
