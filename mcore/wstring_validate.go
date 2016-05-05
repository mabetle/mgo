package mcore

import (
	"regexp"
)

func (in String) IsMatchString(exp string) bool {
	m, _ := regexp.MatchString(exp, in.String())
	return m
}

func (in String) IsMustCompileMatch(pattern string) bool {
	v := regexp.MustCompile(pattern)
	return v.MatchString(in.String())
}

func (in String) IsEnglish() bool {
	return in.IsMatchString("^[a-zA-Z]+$")
}

func (in String) IsChinese() bool {
	return in.IsMatchString("^[\\x{4e00}-\\x{9fa5}]+$")
}

func (in String) IsEmail() bool {
	var emailPattern = "^[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?$"
	//emailPattern = `^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`
	return in.IsMatchString(emailPattern)
}

func (in String) IsPhoneNumber() bool {
	return in.IsMatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`)
}

func (in String) IsNumber() bool {
	return in.IsMatchString(`^[0-9\,\.]+$`)
}

func (in String) IsIdCardNo() bool {
	return in.IsMatchString(`^(\d{15})$`) || in.IsMatchString(`^(\d{17})([0-9]|X)$`)
}
