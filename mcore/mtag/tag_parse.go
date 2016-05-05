package mtag

import (
	"reflect"
)

// GetTag returns tag string
// v should be struct, pointer or array(slice)
// if field not exists, returns "" and false.
func GetTag(v interface{}, fieldName string, tag string) (string, bool) {
	typ := reflect.TypeOf(v)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Slice {
		typ = typ.Elem()
	}

	f, b := typ.FieldByName(fieldName)
	// not found field.
	if !b {
		return "", false
	}
	r := f.Tag.Get(tag)
	return string(r), true
}
