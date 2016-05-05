package mtag

import (
	"github.com/mabetle/mcore"
	"strings"
)

// label tag format:
// label="zh='' en=''"

// GetLabelTag returns field "label" tag value.
func GetLabelTag(v interface{}, fieldName string) (string, bool) {
	return GetTag(v, fieldName, "label")
}

// parse string to KeyValue map.
func ParseKeyValueMap(value string) map[string]string {
	result := make(map[string]string)
	rows := strings.Split(value, " ")
	for _, row := range rows {
		// skip blank
		if strings.TrimSpace(row) == "" {
			continue
		}
		kv := strings.Split(row, "=")
		if len(kv) == 2 {
			k := strings.TrimSpace(kv[0])
			v := strings.TrimSpace(kv[1])
			v = strings.Trim(v, "'")
			v = strings.Trim(v, "\"")
			v = strings.TrimSpace(v)
			result[k] = v
		}
	}
	return result
}

// GetLocaleLabel returns field label by locale.
// locale format: en en_US  / zh zh_CN zh_HK etc.
func GetLocaleLabel(v interface{}, fieldName string, locale string) string {
	labelValue, e := GetLabelTag(v, fieldName)

	// not exist
	if !e {
		return mcore.ToLabel(fieldName)
	}
	locale = strings.Replace(locale, "-", "_", -1)

	lang := strings.Split(locale, "_")[0]

	m := ParseKeyValueMap(labelValue)

	// include lang_coutry locale
	if v, ok := m[locale]; ok {
		return v
	}

	// include lang
	if v, ok := m[lang]; ok {
		return v
	}

	// defult en
	if v, ok := m["en"]; ok {
		return v
	}

	// default return
	return mcore.ToLabel(fieldName)
}

// GetLabelZH
func GetLabelZH(v interface{}, fieldName string) string {
	return GetLocaleLabel(v, fieldName, "zh")
}

// GetLabelEN
func GetLabelEN(v interface{}, fieldName string) string {
	return GetLocaleLabel(v, fieldName, "en")
}
