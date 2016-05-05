package mvalid

import (
	"fmt"
	"github.com/mabetle/mgo/mcore"
	"strings"
)

const (
	TAG_VALIDATOR = "validator"
	V_REQUIRED    = "required"
	V_MIN         = "min"
	V_MAX         = "max"
	V_MIN_LENGTH  = "minlength"
	V_MAX_LENGTH  = "maxlength"
	V_EMAIL       = "email"
	V_IDCARD      = "idcard"
	V_REGEX       = "regex"
	V_NUMBER      = "number"
	V_PHONE       = "phone"
	V_LENGTH      = "length"
)

// Validate
func Validate(v interface{}) map[string]string {
	m := make(map[string]string)
	typ := mcore.GetType(v)
	for i := 0; i < typ.NumField(); i++ {
		fn := typ.Field(i).Name
		logger.Debugf("Validate field:%s", fn)
		msg := ValidateField(v, fn)
		if msg == "" {
			// no error, skip
			continue
		}
		m[fn] = msg
	}
	return m
}

func GetValidatorTag(v interface{}, fn string) string {
	typ := mcore.GetType(v)
	f, e := typ.FieldByName(fn)
	if !e {
		// not found field
		logger.Warnf("not found field:%s", fn)
		return ""
	}
	return f.Tag.Get(TAG_VALIDATOR)
}

// ValidateField
func ValidateField(v interface{}, fn string) (err string) {
	tv := GetValidatorTag(v, fn)
	logger.Debugf("Validate Tag: %s", tv)
	if tv == "" {
		// field no tag, skip
		return ""
	}
	vds := strings.Split(tv, ",")
	fv := mcore.GetFieldValue(v, fn)
	logger.Debugf("Field  Value: %s", fv)

	for _, item := range vds {
		kv := strings.Split(item, "=")
		v_key := strings.TrimSpace(kv[0])
		v_key = strings.ToLower(v_key)
		v_value := ""
		if len(kv) > 1 {
			v_value = kv[1]
			v_value = strings.TrimSpace(v_value)
			v_value = strings.TrimLeft(v_value, "'")
			v_value = strings.TrimRight(v_value, "'")
		}
		switch v_key {
		case V_REQUIRED:
			if fv == "" {
				err = err + "null or blank."
			}
		case V_MIN:
			i_v_value, err1 := mcore.StrToInt(v_value)
			if err1 != nil {
				// parse int error, skip
				continue
			}
			i_f_value, err2 := mcore.StrToInt(fv)
			if err2 != nil {
				// field value parse int error, skip
				continue
			}
			if i_f_value < i_v_value {
				err = err + fmt.Sprintf("%s less than %s.", fv, v_value)
			}
		case V_MAX:
			i_v_value, err1 := mcore.StrToInt(v_value)
			if err1 != nil {
				// parse int error, skip
				continue
			}
			i_f_value, err2 := mcore.StrToInt(fv)
			if err2 != nil {
				// field value parse int error, skip
				continue
			}
			if i_f_value > i_v_value {
				err = err + fmt.Sprintf("%s great than %s.", fv, v_value)
			}
		case V_MAX_LENGTH:
			length, err1 := mcore.StrToInt(v_value)
			if err1 != nil {
				// parse int error, skip
				continue
			}
			if len(fv) > length {
				err = err + fmt.Sprintf("%s length great than %s.", fv, v_value)
			}
		case V_MIN_LENGTH:
			length, err1 := mcore.StrToInt(v_value)
			if err1 != nil {
				// parse int error, skip
				continue
			}
			if len(fv) < length {
				err = err + fmt.Sprintf("%s length less than %s.", fv, v_value)
			}
		case V_LENGTH:
			length, err1 := mcore.StrToInt(v_value)
			if err1 != nil {
				// parse int error, skip
				continue
			}
			if len(fv) != length {
				err = err + fmt.Sprintf("%s length not equals %s.", fv, v_value)
			}
		case V_EMAIL:
			if !mcore.NewString(fv).IsEmail() {
				err = err + fmt.Sprintf("%s not a Email.", fv)
			}
		case V_IDCARD:
			if !mcore.NewString(fv).IsIdCardNo() {
				err = err + fmt.Sprintf("%s not a ID Card number.", fv)
			}
		case V_NUMBER:
			if !mcore.NewString(fv).IsNumber() {
				err = err + fmt.Sprintf("%s not a number.", fv)
			}
		case V_PHONE:
			if !mcore.NewString(fv).IsPhoneNumber() {
				err = err + fmt.Sprintf("%s not a phone number.", fv)
			}
		case V_REGEX:
			if !mcore.NewString(fv).IsMatchString(v_value) {
				err = err + fmt.Sprintf("%s not match %s.", fv, v_value)
			}
		default:
			// unknow validator tag.
			logger.Warnf("Unknown validate tag: %s ", item)
		}
	}
	logger.Debugf("Validate Result: %s", err)
	return
}

func PrintValidate(v interface{}) {
	m := Validate(v)
	fmt.Printf("Vildate Result:\n")
	for k, v := range m {
		fmt.Printf("Filed:%s Result: %s \n", k, v)
	}
}
