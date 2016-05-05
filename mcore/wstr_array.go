package mcore

import (
	"fmt"
	"strings"
)

type StringArray []string

func NewStringArray(items ...string) StringArray {
	return StringArray(items)
}

// new line one item
func (s StringArray) Merge() string {
	sb := NewStringBuffer()
	for _, v := range s {
		sb.Appendf("%s\n", v)
	}
	return sb.String()
}

// MergeWithOsNewLine
func (s StringArray) MergeWithOsNewLine() string {
	sb := NewStringBuffer()
	for _, v := range s {
		sb.Appendf("%s%s", v, O_NL)
	}
	return sb.String()
}

func (s StringArray) AppendStringArray(a StringArray) StringArray {
	for _, v := range a {
		s = append(s, v)
	}
	return s
}

func (s StringArray) Join(sep string) string {
	return s.RoundJoin("", sep)
}

// example: {"a", "b" , "c"}  to "a,b,c"
func (s StringArray) RoundJoin(round, sep string) string {
	n := len(s) - 1
	sb := NewStringBuffer()
	for i := 0; i <= n; i++ {
		item := fmt.Sprintf("%s%s%s", round, s[i], round)
		if i == n {
			sb.Append(item)
		} else {
			sb.Appendf("%s%s", item, sep)
		}
	}
	return sb.String()
}

func (s StringArray) RemoveBlank() (r StringArray) {
	for _, line := range s {
		line = strings.TrimSpace(line)
		if line != "" {
			r = append(r, line)
		}
	}
	return
}

func (s StringArray) RemoveComment(start string) (r StringArray) {
	for _, line := range s {
		if !strings.HasPrefix(line, start) {
			r = append(r, line)
		}
	}
	return
}

func (s StringArray) RemoveSharpComment() StringArray {
	return s.RemoveComment("#")
}

func (s StringArray) RemoveSlashComment() StringArray {
	return s.RemoveComment("//")
}

func (s StringArray) RemoveBlankAndComment(start string) StringArray {
	return s.RemoveBlank().RemoveComment(start)
}

func (s StringArray) RemoveBlankAndSharpComment() StringArray {
	return s.RemoveBlank().RemoveComment("#")
}

func (s StringArray) RemoveBlankAndSlashComment() StringArray {
	return s.RemoveBlank().RemoveComment("//")
}
