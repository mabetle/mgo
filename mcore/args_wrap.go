package mcore

import (
	"fmt"
	"strings"
)

// WrapArgs wrap renderArgs and args together
type WrapArgs struct {
	renderArgs map[string]interface{}
	args       []string
	wrapArgs   map[string]string
}

// NewWrapArgs creates instance
func NewWrapArgs(renderArgs map[string]interface{}, args ...string) *WrapArgs {
	ins := &WrapArgs{renderArgs: renderArgs, args: args}
	ins.wrapArgs = make(map[string]string)
	for k, v := range renderArgs {
		ins.wrapArgs[k] = fmt.Sprint(v)
	}
	for _, v := range args {
		kv := strings.Split(v, "=")
		if len(kv) != 2 {
			continue
		}
		key := kv[0]
		value := kv[1]
		ins.wrapArgs[key] = value
	}
	return ins
}

// GetArgString returns name string value
func (a *WrapArgs) GetArgString(name string, dv string) string {
	v, ok := a.wrapArgs[name]
	if !ok {
		return dv
	}
	return v
}

// IsArgExists returns name exists or not
func (a *WrapArgs) IsArgExists(name string) bool {
	_, ok := a.wrapArgs[name]
	return ok
}

// GetArgBool
func (a *WrapArgs) GetArgBool(name string, dv bool) bool {
	v, ok := a.wrapArgs[name]
	if !ok {
		return dv
	}
	return NewString(v).ToBool()
}

// GetArgInt returns int
func (a *WrapArgs) GetArgInt(name string, dv int) int {
	v, ok := a.wrapArgs[name]
	if !ok {
		return dv
	}
	iv, err := NewString(v).ToInt()
	if err != nil {
		return dv
	}
	return iv
}

// GetArgFloat64 returns Float64
func (a *WrapArgs) GetArgFloat64(name string, dv float64) float64 {
	v, ok := a.wrapArgs[name]
	if !ok {
		return dv
	}
	fv, err := NewString(v).ToFloat64()
	if err != nil {
		return dv
	}
	return fv
}
