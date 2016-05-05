package mcore

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// GetType returns value type.
// support struct and pointer both.
func GetType(v interface{}) reflect.Type {
	typ := reflect.TypeOf(v)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ
}

// GetTypeName returns value type name.
func GetTypeName(v interface{}) string {
	return GetType(v).Name()
}

// GetPkgPath
func GetPkgPath(v interface{}) string {
	return GetType(v).PkgPath()
}

// GetElementType
// v should be a slicce.
func GetElementType(v interface{}) reflect.Type {
	m := GetArrayFirstElement(v)
	return GetType(m)
}

// GetElementTypeName values Type
// support struct, pointer, slice
// slice returns first element type.
func GetElementTypeName(v interface{}) string {
	m := GetArrayFirstElement(v)
	return GetTypeName(m)
}

// GetArrayFirstElement return rows first row.
// rows must be a slice.
func GetArrayFirstElement(rows interface{}) interface{} {
	v := reflect.ValueOf(rows)
	if v.Kind() == reflect.Slice {
		return v.Index(0).Interface()
	}
	// if not a slice, return input arg
	return rows
}

// GetArrayFields m must be a array type.
func GetArrayFields(m interface{}) []string {
	v := reflect.ValueOf(m)
	if v.Kind() == reflect.Slice {
		e := v.Index(0).Interface()
		return GetFields(e)
	}
	return GetFields(m)
}

// GetUsedArrayFields
func GetUsedArrayFields(m interface{}, include, exclude string) []string {
	v := reflect.ValueOf(m)
	if v.Kind() == reflect.Slice {
		e := v.Index(0).Interface()
		return GetUsedFields(e, include, exclude)
	}
	return GetUsedFields(m, include, exclude)
}

// GetFields returns Fields Name Array, m could be pointer or struct.
// Emmber Struct call in rotate
func GetFields(m interface{}) (ns []string) {
	typ := GetType(m)
	// Only structs are supported so return an empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
		return
	}
	// loop through the struct's fields
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.Anonymous {
			fkind := f.Type.Kind()
			if fkind == reflect.Struct || fkind == reflect.Ptr {
				fns := GetFields(reflect.New(f.Type).Interface())
				for _, fn := range fns {
					if String(fn).IsInArray(ns) {
						continue
					}
					ns = append(ns, fn)
				}
			}
		} else {
			if String(f.Name).IsInArray(ns) {
				continue
			}
			ns = append(ns, f.Name)
		}
	}
	return ns
}

// GetUsedFields
func GetUsedFields(v interface{}, include, exclude string) (r []string) {
	fs := GetFields(v)
	return GetFieldsUsed(fs, include, exclude)
}

// GetFieldsUsed
// Order by Include if has include
func GetFieldsUsed(fs []string, include, exclude string) (r []string) {
	includes := strings.Split(include, ",")
	excludes := strings.Split(exclude, ",")
	// has include param
	if include != "" {
		for _, in := range includes {
			inString := String(in)
			if !inString.IsInArrayIgnoreCase(fs) {
				continue
			}
			if exclude != "" && inString.IsInArrayIgnoreCase(excludes) {
				continue
			}
			if inString.IsContainIgnoreCase("id") || inString.IsContainIgnoreCase("password") {
				continue
			}
			r = append(r, in)
		}
	}
	// no include param
	if include == "" {
		for _, key := range fs {
			keyString := String(key)
			if exclude != "" && keyString.IsInArrayIgnoreCase(excludes) {
				continue
			}
			if keyString.IsContainIgnoreCase("id") || keyString.IsContainIgnoreCase("password") {
				continue
			}
			r = append(r, key)
		}
	}
	return
}

// GetFieldsUsed2
// Order by fs
func GetFieldsUsed2(fs []string, include, exclude string) (r []string) {
	includes := strings.Split(include, ",")
	excludes := strings.Split(exclude, ",")
	for _, key := range fs {
		keyString := String(key)
		// exist include and not in include, skip
		if include != "" && !keyString.IsInArrayIgnoreCase(includes) {
			continue
		}
		// in exclude, skip
		if keyString.IsInArrayIgnoreCase(excludes) {
			continue
		}
		if keyString.IsContainIgnoreCase("id") || keyString.IsContainIgnoreCase("password") {
			continue
		}
		r = append(r, key)
	}
	return
}

// GetMethods
func GetMethods(v interface{}) (r []string) {
	value := reflect.ValueOf(v)
	typ := value.Type()
	for i := 0; i < value.NumMethod(); i++ {
		r = append(r, typ.Method(i).Name)
	}
	return
}

// IsHasMethod use GetMethods to judge.
func IsHasMethod(v interface{}, methodName string) bool {
	return String(methodName).IsInArrayIgnoreCase(GetMethods(v))
}

// GetFieldType return refelect.Kind, so can switch it.
func GetFieldType(v interface{}, field string) reflect.Kind {
	var immutable reflect.Value
	immutable = GetReflectValue(v)
	val := immutable.FieldByName(field)
	return val.Kind()
}

// GetFieldValueOrigin returns interface{}
func GetFieldValueOrigin(v interface{}, field string) interface{} {
	var immutable reflect.Value
	immutable = GetReflectValue(v)
	val := immutable.FieldByName(field)
	return val.Interface()
}

// GetFieldValue v can be struct or pointer.
func GetFieldValue(v interface{}, field string) (r string) {
	var immutable reflect.Value
	immutable = GetReflectValue(v)
	val := immutable.FieldByName(field)
	switch val.Kind() {
	case reflect.Int64, reflect.Int32, reflect.Int:
		r = fmt.Sprintf("%d", val.Int())
	case reflect.Float64, reflect.Float32:
		r = fmt.Sprintf("%.2f", val.Float())
	default:
		// process time
		vi := val.Interface()
		if vc, ok := vi.(time.Time); ok {
			r = FormatTime(vc)
			break
		}
		r = fmt.Sprintf("%v", val)
	}
	return
}

// GetReflectValue, returns reflect.Value, support both struct and pointer.
// if value is a pointer, indirect into it.
func GetReflectValue(v interface{}) reflect.Value {
	var immutable reflect.Value
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		immutable = reflect.Indirect(reflect.ValueOf(v))
	} else {
		immutable = reflect.ValueOf(v)
	}
	return immutable
}

// GetReflectFieldValue returns reflect.Value.
func GetReflectFieldValue(v interface{}, fn string) reflect.Value {
	return GetReflectValue(v).FieldByName(fn)
}

// SetFieldValue change struct field value.
// should use struct pointer.
func SetFieldValue(v interface{}, field string, newValue interface{}) interface{} {
	var immutable reflect.Value
	immutable = GetReflectValue(v)
	f := immutable.FieldByName(field)
	if f.IsValid() && f.CanSet() {
		if f.Kind() == reflect.Int {
			f.SetInt(newValue.(int64))
		} else {
			f.Set(reflect.ValueOf(newValue))
		}
	}
	return v
}

// PrintType pring value type
func PrintType(v interface{}) {
	typ := reflect.TypeOf(v)
	fmt.Println(typ)
}
