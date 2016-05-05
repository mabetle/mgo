package mcore

import (
	"github.com/mabetle/mcore/mtest"
	"testing"
)

func TestStringValidator(t *testing.T) {
	mtest.RegTest(t)

	//test email
	mtest.Equal(String("demo@demo.com").IsEmail(), true, "demo@demo.com is a email")
	mtest.Equal(String("demo.com").IsEmail(), false, "demo.com is not a email")

	//test english
	mtest.Equal(String("你好").IsChinese(), true, "你好 is no Engllish")
	mtest.Equal(String("Hello").IsEnglish(), true, "Hello is English")

	//test chinese
	mtest.Equal(String("你好").IsChinese(), true, "你好 is Chinese")
	mtest.Equal(String("Hello").IsEnglish(), true, "Hello is not Chinese")

	//test Idcard
	mtest.Equal(String("110110198801018721").IsIdCardNo(), true, "Is a IdCard number")
	mtest.Equal(String("hello").IsIdCardNo(), false, "Not a IdCard number")

	//test Phone number
	mtest.Equal(String("13900910001").IsPhoneNumber(), true, "Is a phone number")
	mtest.Equal(String("139sss00910001").IsPhoneNumber(), false, "Not a phone number")

	mtest.Equal(String("asdf").IsNumber(), false, "not a number")
	mtest.Equal(String("1012,312.12").IsNumber(), true, "not a number")

}
