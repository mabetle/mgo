package mcore

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type String string

// NewString
func NewString(v interface{}) String {
	return String(fmt.Sprintf("%v", v))
}

// String
func (s String) String() string {
	return string(s)
}

// Int trans String to int
func (s String) Int() int {
	n, err := strconv.Atoi(s.String())
	log.Printf("Error: convert %v to int error: %v ", s, err)
	return n
}

// Value equal to String()
func (s String) Value() string {
	return string(s)
}

// Len
func (s String) Len() int {
	return len(s)
}

// IsContains
func (s String) IsContains(sub string) bool {
	return strings.Contains(string(s), sub)
}

// IsContainsAny
func (s String) IsContainsAny(chars string) bool {
	return strings.ContainsAny(string(s), chars)
}

// IsContainsRune
func (s String) IsContainsRune(r rune) bool {
	return strings.ContainsRune(string(s), r)
}

// IsContainIgnoreCase
func (s String) IsContainIgnoreCase(sub string) bool {
	sl := strings.ToLower(string(s))
	subl := strings.ToLower(sub)
	return strings.Contains(sl, subl)
}

// Count
func (s String) Count(sep string) int {
	return strings.Count(string(s), sep)
}

// IsEqualFold
func (s String) IsEqualFold(t string) bool {
	return strings.EqualFold(string(s), t)
}

// Fields
func (s String) Fields() []string {
	return strings.Fields(string(s))
}

// IsHasPrefix
func (s String) IsHasPrefix(prefix string) bool {
	return strings.HasPrefix(string(s), prefix)
}

// IsStartWith equal to IsHasPrefix
func (s String) IsStartWith(start string) bool {
	return s.IsHasPrefix(start)
}

// IsStartIgnoreCase
func (s String) IsStartIgnoreCase(prefix string) bool {
	ls := strings.ToLower(string(s))
	lp := strings.ToLower(prefix)
	return strings.HasPrefix(ls, lp)
}

// IsStartInIgnoreCase
func (s String) IsStartInIgnoreCase(args []string) bool {
	for _, arg := range args {
		if String(arg).IsStartIgnoreCase(string(s)) {
			return true
		}
	}
	return false
}

// IsStartsIgnoreCase
func (s String) IsStartsIgnoreCase(prefixs ...string) bool {
	for _, prefix := range prefixs {
		if s.IsStartIgnoreCase(prefix) {
			return true
		}
	}
	return false
}

// IsStartsIgnoreCaseInArray
func (s String) IsStartsIgnoreCaseInArray(prefixs []string) bool {
	for _, prefix := range prefixs {
		if s.IsStartIgnoreCase(prefix) {
			return true
		}
	}
	return false
}

// IsEndWith equal to IsHasSuffix
func (s String) IsEndWith(end string) bool {
	return s.IsHasSuffix(end)
}

// IsHasSuffix
func (s String) IsHasSuffix(suffix string) bool {
	return strings.HasSuffix(string(s), suffix)
}

// Index
func (s String) Index(sep string) int {
	return strings.Index(string(s), sep)
}

// IndexAny
func (s String) IndexAny(chars string) int {
	return strings.IndexAny(string(s), chars)
}

// IsHasAny is string include some string.
// equal to IsContains
func (s String) IsHasAny(chars string) (r bool) {
	if s.IndexAny(chars) != -1 {
		r = true
	}
	return
}

// IndexByte
func (s String) IndexByte(c byte) int {
	return strings.IndexByte(string(s), c)
}

// IndexRune
func (s String) IndexRune(r rune) int {
	return strings.IndexRune(string(s), r)
}

// LastIndex
func (s String) LastIndex(sep string) int {
	return strings.LastIndex(string(s), sep)
}

// LastIndexAny
func (s String) LastIndexAny(chars string) int {
	return strings.LastIndexAny(string(s), chars)
}

// Repeat
func (s String) Repeat(count int) String {
	r := strings.Repeat(string(s), count)
	return String(r)
}

// Replace
func (s String) Replace(old, newStr string, n int) String {
	r := strings.Replace(string(s), old, newStr, n)
	return String(r)
}

// ReplaceAll
func (s String) ReplaceAll(old, newStr string) String {
	r := s.Replace(old, newStr, -1)
	return String(r)
}

// ReplaceAllNumber Replace all numbers
func (s String) ReplaceAllNumber(newStr string) String {
	s = s.ReplaceAll("0", newStr)
	s = s.ReplaceAll("1", newStr)
	s = s.ReplaceAll("2", newStr)
	s = s.ReplaceAll("3", newStr)
	s = s.ReplaceAll("4", newStr)
	s = s.ReplaceAll("5", newStr)
	s = s.ReplaceAll("6", newStr)
	s = s.ReplaceAll("7", newStr)
	s = s.ReplaceAll("8", newStr)
	s = s.ReplaceAll("9", newStr)
	return s
}

// Split
func (s String) Split(sep string) []string {
	return strings.Split(string(s), sep)
}

// Title
func (s String) Title() String {
	return String(strings.Title(string(s)))
}

// ToLower
func (s String) ToLower() String {
	return String(strings.ToLower(string(s)))
}

// ToUpper
func (s String) ToUpper() String {
	return String(strings.ToUpper(string(s)))
}

// Trim
func (s String) Trim(cutset string) String {
	return String(strings.Trim(string(s), cutset))
}

// TrimLeft
func (s String) TrimLeft(cutset string) String {
	return String(strings.TrimLeft(string(s), cutset))
}

// TrimStart equal to TrimLeft
func (s String) TrimStart(start string) String {
	return s.TrimLeft(start)
}

// TrimStarts
func (s String) TrimStarts(starts ...string) String {
	for _, start := range starts {
		s = s.TrimStart(start)
	}
	return s
}

// TrimEnds
func (s String) TrimEnds(ends ...string) String {
	for _, end := range ends {
		s = s.TrimEnd(end)
	}
	return s
}

// TrimRight
func (s String) TrimRight(cutset string) String {
	return String(strings.TrimRight(string(s), cutset))
}

// TrimEnd equal to TrimRight
func (s String) TrimEnd(end string) String {
	return s.TrimRight(end)
}

// TrimSuffix
func (s String) TrimSuffix(suffix string) String {
	return String(strings.TrimSuffix(string(s), suffix))
}

// TrimPrefix
func (s String) TrimPrefix(prefix string) String {
	return String(strings.TrimPrefix(string(s), prefix))
}

// TrimSpace
func (s String) TrimSpace() String {
	return String(strings.TrimSpace(string(s)))
}

// IsEqualIgnoreCase
func (s String) IsEqualIgnoreCase(v string) bool {
	if strings.ToLower(string(s)) == strings.ToLower(v) {
		return true
	}
	return false
}

// IsIn is in array
func (s String) IsIn(args ...string) bool {
	for _, v := range args {
		if v == string(s) {
			return true
		}
	}
	return false
}

// IsNotIn reverse for IsIn
func (s String) IsNotIn(args ...string) bool {
	return !s.IsIn(args...)
}

// IsInSepString
func (s String) IsInSepString(arg string, sep string) bool {
	args := String(arg).Split(sep)
	return s.IsInArray(args)
}

// IsNotInSepString
func (s String) IsNotInSepString(arg string, sep string) bool {
	return !s.IsInSepString(arg, sep)
}

// IsInSepStringIgnoreCase
func (s String) IsInSepStringIgnoreCase(arg string, sep string) bool {
	args := String(arg).Split(sep)
	return s.IsInArrayIgnoreCase(args)
}

// IsNotInSepStringIgnoreCase
func (s String) IsNotInSepStringIgnoreCase(arg, sep string) bool {
	return !s.IsInSepStringIgnoreCase(arg, sep)
}

// IsInIgnoreCase is in array ignore case.
func (s String) IsInIgnoreCase(args ...string) bool {
	for _, v := range args {
		if s.IsEqualIgnoreCase(v) {
			return true
		}
	}
	return false
}

// IsNotInIgnoreCase
func (s String) IsNotInIgnoreCase(args ...string) bool {
	return !s.IsInIgnoreCase(args...)
}

// IsInArray Is In array
func (s String) IsInArray(args []string) bool {
	for _, v := range args {
		if v == string(s) {
			return true
		}
	}
	return false
}

// IsNotInArray
func (s String) IsNotInArray(args []string) bool {
	return !s.IsInArray(args)
}

// IsInArrayIgnoreCase
func (s String) IsInArrayIgnoreCase(args []string) bool {
	for _, v := range args {
		if strings.ToLower(v) == strings.ToLower(string(s)) {
			return true
		}
	}
	return false
}

// IsContainsInSepString
func (s String) IsContainsInSepString(arg string, sep string) bool {
	return s.IsContainInArray(strings.Split(arg, sep))
}

// IsContainsInSepStringIgnoreCase
func (s String) IsContainsInSepStringIgnoreCase(arg string, sep string) bool {
	return s.IsContainInArrayIgnoreCase(strings.Split(arg, sep))
}

// IsContainInArrayIgnoreCase
func (s String) IsContainInArrayIgnoreCase(args []string) bool {
	for _, v := range args {
		v = strings.ToLower(v)
		if s.ToLower().IsContains(v) {
			return true
		}
	}
	return false
}

// IsContainInArray
func (s String) IsContainInArray(args []string) bool {
	for _, v := range args {
		if s.IsContains(v) {
			return true
		}
	}
	return false
}

// IsNotInArrayIgnoreCase
func (s String) IsNotInArrayIgnoreCase(args []string) bool {
	return !s.IsInArrayIgnoreCase(args)
}

// IsBlank if string is blank
func (s String) IsBlank() (r bool) {
	if string(s.TrimSpace()) == "" {
		r = true
	}
	return
}

// IsNotBlank
func (s String) IsNotBlank() bool {
	return !s.IsBlank()
}

// Sub string
func (s String) Sub(start, length int) String {
	r := Sub(string(s), start, length)
	return String(r)
}

// SubLeftSep
func (s String) SubLeftSep(sep string) string {
	if !s.IsContains(sep) {
		return s.String()
	}
	as := s.Split(sep)
	return strings.Join(as[:1], "")
}

// SubRightSep
func (s String) SubRightSep(sep string) string {
	if !s.IsContains(sep) {
		return s.String()
	}
	as := s.Split(sep)
	return strings.Join(as[1:], "")
}

// SubLeft
func (s String) SubLeft(length int) String {
	r := SubLeft(string(s), length)
	return String(r)
}

// SubRight
func (s String) SubRight(length int) String {
	r := SubRight(string(s), length)
	return String(r)
}

// AddStart
func (s String) AddStart(start string) String {
	return String(start + string(s))
}

// AddEnd
func (s String) AddEnd(end string) String {
	return String(string(s) + end)
}

// AddPrefix
func (s String) AddPrefix(pre string) String {
	if s.IsHasPrefix(pre) {
		return s
	}
	return s.AddStart(pre)
}

// AddSuffix
func (s String) AddSuffix(suf string) String {
	if s.IsHasSuffix(suf) {
		return s
	}
	return s.AddEnd(suf)
}

// SepEnd "abc.go" SepEnd(".") return go.
func (s String) SepEnd(sep string) (r String) {
	if !s.IsHasAny(sep) {
		return
	}
	a := s.Split(sep)
	item := a[len(a)-1]
	return String(item)
}

// SepStart
// abc.go return abc
func (s String) SepStart(sep string) (r String) {
	if !s.IsHasAny(sep) {
		return
	}
	return String(s.Split(sep)[0])
}

// TrimQuotes "abc" to abc
func (s String) TrimQuotes() String {
	if s == "" {
		return s
	}
	if s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

// TrimEndIndex eg: "a/b/c" TrimEndIndex("/") returns a/b
func (s String) TrimEndIndex(sep string) String {
	if !s.IsContains(sep) {
		println("not contains ")
		return s
	}
	n := s.LastIndex(sep)
	println(n)
	return s.SubLeft(n)
}

// TrimBeginIndex
func (s String) TrimBeginIndex(sep string) String {
	if !s.IsContains(sep) {
		return s
	}
	n := s.Index(sep) + len(sep)
	return s.Sub(n, len(s))
}
