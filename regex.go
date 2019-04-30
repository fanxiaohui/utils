package utils

import (
	"regexp"
)

const (
	regexChineseParttern    = `^[\x{4e00}-\x{9fa5}]+$`
	regexEnglishParttern    = `^[a-zA-Z]+$`
	regexPhoneParttern      = `^(1[3|4|5|8][0-9]\d{4,8})$`
	regexIDcard15Parttern   = `^(\d{15})$`
	regexIDcard18Parttern   = `^(\d{17})([0-9]|X)$`
	regexDigitParttern      = `^[0-9]+$`
	regexEmailPattern       = `(?i)[A-Z0-9._%+-]+@(?:[A-Z0-9-]+\.)+[A-Z]{2,6}`
	regexStrictEmailPattern = `(?i)[A-Z0-9!#$%&'*+/=?^_{|}~-]+` +
		`(?:\.[A-Z0-9!#$%&'*+/=?^_{|}~-]+)*` +
		`@(?:[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?\.)+` +
		`[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?`
	regexURLPattern = `(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?`
)

// IsChinese 是否是中文
func IsChinese(s string) bool {
	return regexp.MustCompile(regexChineseParttern).MatchString(s)
}

// IsEnglish 是否是英文
func IsEnglish(s string) bool {
	return regexp.MustCompile(regexEnglishParttern).MatchString(s)
}

// IsPhone 是否是电话号码
func IsPhone(s string) bool {
	return regexp.MustCompile(regexPhoneParttern).MatchString(s)
}

// IsIDCard 是否是身份证号
func IsIDCard(s string) bool {
	if len(s) == 15 {
		return regexp.MustCompile(regexIDcard15Parttern).MatchString(s)
	} else if len(s) == 18 {
		return regexp.MustCompile(regexIDcard18Parttern).MatchString(s)
	}

	return false
}

// IsDigit 是否是数字
func IsDigit(s string) bool {
	return regexp.MustCompile(regexDigitParttern).MatchString(s)
}

// IsEmail 是否是email地址
func IsEmail(email string) bool {
	return regexp.MustCompile(regexEmailPattern).MatchString(email)
}

// IsEmailRFC 是否是email地址
// this validation omits RFC 2822
func IsEmailRFC(email string) bool {
	return regexp.MustCompile(regexStrictEmailPattern).MatchString(email)
}

// IsURL 是否是url地址
func IsURL(url string) bool {
	return regexp.MustCompile(regexURLPattern).MatchString(url)
}
