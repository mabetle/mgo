package mcore

import (
	"strings"
)

// ToTableName
// example: DemoTable to demo_table
func ToTableName(name string) string {
	result := ""

	lastIndexDot := strings.LastIndex(name, ".")
	if lastIndexDot != -1 {
		name = SubRight(name, lastIndexDot+1)
	}

	b := []byte(name)

	for i := 0; i < len(b); i++ {
		l := string(b[i])

		//first letter
		if i == 0 {
			result = result + strings.ToLower(l)
			continue
		}

		//Up letter
		if IsUpLetter(l) {
			result = result + "_" + strings.ToLower(l)
		} else {
			result = result + l
		}
	}

	return result
}

// ToCamel
// demo_demo  > DemoDemo
func ToCamel(name string) string {
	result := ""
	b := []byte(name)
	size := len(b)

	//upcase first letter
	result = strings.ToUpper(string(b[0]))

	for i := 1; i < size; i++ {

		l := string(b[i])
		bl := string(b[i-1])

		if "_" == l {
			continue
		}

		//front letter is _, then upcase add append
		if bl == "_" {
			result = result + strings.ToUpper(l)
			continue
		}

		//front letter is not _ append it
		result = result + l

	}
	return result

}

// dir such as abc/def or abc\\def for windows, return def
func GetPackageNameFromDir(dir string) string {
	return String(dir).ReplaceAll("\\", "/").TrimSuffix("/").SepEnd("/").String()
}

// ToLabel returns space spilt string.
func ToLabel(name string) string {
	name = strings.Replace(name, "-", " ", -1)
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Replace(name, ".", " ", -1)
	sb := NewStringBuffer()
	for i := 0; i < len(name); i++ {
		c := string(name[i])
		cb := ""
		// before letter
		if i > 0 {
			cb = string(name[i-1])
		}

		if IsUpLetter(c) {
			if IsUpLetter(cb) {
				// cb is upcase, and c is upcase
				sb.Append(c)
			} else {
				// cb not upcase
				sb.Append(" ", c)
			}
		} else {
			// other as original
			sb.Append(c)
		}
	}
	as := strings.Split(sb.String(), " ")
	sb2 := NewStringBuffer()
	for i, a := range as {
		// skip space.
		if strings.TrimSpace(a) == "" {
			continue
		}
		// upcase first letter
		a = UpperCaseFirst(a)

		if i == len(as) {
			sb2.Append(a)
		} else {
			sb2.Append(a, " ")
		}
	}
	return sb2.String()
}
