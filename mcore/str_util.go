package mcore

//String Utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const (
	UP_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// IsUpLetter
func IsUpLetter(l string) bool {
	if len(l) != 1 {
		return false
	}
	return strings.Contains(UP_LETTERS, l)
}

// SubRight
func SubRight(s string, start int) string {
	return Sub(s, start, len(s)-start)
}

// SubLeft
func SubLeft(s string, end int) string {
	return Sub(s, 0, end)
}

// Sub
func Sub(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// SubByByte
func SubByByte(str string, length int) string {
	bs := []byte(str)[:length]
	bl := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] >= 0 && bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++
		case bs[i] >= 192 && bs[i] <= 253:
			cl := 0
			switch {
			case bs[i]&252 == 252:
				cl = 6
			case bs[i]&248 == 248:
				cl = 5
			case bs[i]&240 == 240:
				cl = 4
			case bs[i]&224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return ""
}

// GetFixedWidthNum used for line number
// add space left side, align right
func GetFixedWidthNum(num int, width int) string {
	s := fmt.Sprint(num)
	return GetFixedWidthString(s, width, " ", false)
}

// GetFixedWidthStringAlignLeft add blank to left
// used for pretty output.
func GetFixedWidthStringAlignLeft(v string, width int) string {
	return GetFixedWidthString(v, width, " ", true)
}

// GetFixedWidthStringAlignRight add blank right
func GetFixedWidthStringAlignRight(v string, width int) string {
	return GetFixedWidthString(v, width, " ", false)
}

// GetFixedWidthString
func GetFixedWidthString(
	v string,
	width int,
	fil string,
	alignLeft bool) string {
	n := StringWidth(v)
	// not cut long string
	if n >= width {
		return v
	}
	// less than width
	fillNums := width - n
	if fil == "" {
		fil = " "
	}
	fillStr := strings.Repeat(fil, fillNums)
	// add blank right side
	if alignLeft {
		return v + fillStr
	}
	// add blank left side
	return fillStr + v
}

// EncodeGBK
func EncodeGBK(in string) string {
	reader := transform.NewReader(bytes.NewReader([]byte(in)), simplifiedchinese.GBK.NewEncoder())
	d, _ := ioutil.ReadAll(reader)
	return string(d)
}

// UpperCaseFirst
func UpperCaseFirst(in string) string {
	in = strings.TrimSpace(in)
	begin := SubLeft(in, 1)
	begin = strings.ToUpper(begin)
	end := SubRight(in, 1)
	return begin + end
}

// StringWidth chinese char is width 2
func StringWidth(v string) int {
	r := 0
	for _, c := range []rune(v) {
		if IsChineseChar(string(c)) {
			r = r + 2
		} else {
			r = r + 1
		}
	}
	return r
}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// StringLen how many chars
func StringLen(v string) int {
	return len([]rune(v))
}

func TrimSepLast(v string, sep string) string {
	vs := strings.Split(v, sep)
	if len(vs) < 2 {
		return v
	}
	return strings.Join(vs[:len(vs)-1], sep)
}
