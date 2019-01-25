package common

import (
    "regexp"
)

const (
    regex_chinese_parttern     = `^[\x{4e00}-\x{9fa5}]+$`
    regex_english_parttern     = `^[a-zA-Z]+$`
    regex_phone_parttern       = `^(1[3|4|5|8][0-9]\d{4,8})$`
    regex_IDcard15_parttern    = `^(\d{15})$`
    regex_IDcard18_parttern    = `^(\d{17})([0-9]|X)$`
    regex_digit_parttern       = `^[0-9]+$`
    regex_email_pattern        = `(?i)[A-Z0-9._%+-]+@(?:[A-Z0-9-]+\.)+[A-Z]{2,6}`
    regex_strict_email_pattern = `(?i)[A-Z0-9!#$%&'*+/=?^_{|}~-]+` +
        `(?:\.[A-Z0-9!#$%&'*+/=?^_{|}~-]+)*` +
        `@(?:[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?\.)+` +
        `[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?`
    regex_url_pattern = `(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?`
)

// 是否是中文
func IsChinese(s string) bool {
    return regexp.MustCompile(regex_chinese_parttern).MatchString(s)
}

// 是否是英文
func IsEnglish(s string) bool {
    return regexp.MustCompile(regex_english_parttern).MatchString(s)
}

// 是否是电话号码
func IsPhone(s string) bool {
    return regexp.MustCompile(regex_phone_parttern).MatchString(s)
}

// 是否是身份证号
func IsIDCard(s string) bool {
    if len(s) == 15 {
        return regexp.MustCompile(regex_IDcard15_parttern).MatchString(s)
    } else if len(s) == 18 {
        return regexp.MustCompile(regex_IDcard18_parttern).MatchString(s)
    }

    return false
}

func IsDigit(s string) bool {
    return regexp.MustCompile(regex_digit_parttern).MatchString(s)
}

// 是否是email地址
func IsEmail(email string) bool {
    return regexp.MustCompile(regex_email_pattern).MatchString(email)
}

// 是否是email地址
// this validation omits RFC 2822
func IsEmailRFC(email string) bool {
    return regexp.MustCompile(regex_strict_email_pattern).MatchString(email)
}

// 是否是url地址
func IsUrl(url string) bool {
    return regexp.MustCompile(regex_url_pattern).MatchString(url)
}
