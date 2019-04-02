package utils

import "strings"

// AppendStr appends string to slice with no duplicates.
// 追加字符串到无重复项的切片
func AppendStr(strs []string, str string) []string {
	for _, s := range strs {
		if s == str {
			return strs
		}
	}
	return append(strs, str)
}

// AppendUint appends uint to slice with no duplicates.
// 追加uint到无重复项的切片
func AppendUint(s []uint, e uint) []uint {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// AppendUint16 appends uint16 to slice with no duplicates.
// 追加uint16到无重复项的切片
func AppendUint16(s []uint16, e uint16) []uint16 {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// 删除string切片中的 第一个出现的指定元素
func DeleteFromSliceStr(s []string, e string) []string {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}

	return s
}

// 删除string切片中的 所有出现的指示元素
func DeleteFromSliceStrAll(s []string, e string) []string {
	var tmpS []string

	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}

	return tmpS
}

// 删除uint16切片中的 第一个出现的指定元素
func DeleteFromSliceUint16(s []uint16, e uint16) []uint16 {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}

	return s
}

// 删除uint16切片中的 所有出现的指示元素
func DeleteFromSliceUint16All(s []uint16, e uint16) []uint16 {
	var tmpS []uint16

	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}

	return tmpS
}

// 删除uint切片中的 第一个出现的指定元素
func DeleteFromSliceUint(s []uint, e uint) []uint {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}

	return s
}

// 删除uint切片中的 所有出现的指示元素
func DeleteFromSliceUintAll(s []uint, e uint) []uint {
	var tmpS []uint

	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}

	return tmpS
}

// CompareSliceStr compares two 'string' type slices.
// It returns true if elements and order are both the same.
// 比较两个字符串切片,要求元素和顺序都一致才返回true
func CompareSliceStr(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// CompareSliceStr compares two 'string' type slices.
// It returns true if elements are the same, and ignores the order.
// 比较两个字符串切片,要求元素一致,但忽略顺序才返回true
func CompareSliceStrU(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		for j := len(s2) - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				s2 = append(s2[:j], s2[j+1:]...)
				break
			}
		}
	}
	if len(s2) > 0 {
		return false
	}
	return true
}

// IsSliceContainsStr returns true if the string exists in given slice
// 字符串切片是否含有指定的元素,大小写敏感
func IsSliceContainsStr(sl []string, str string) bool {
	for _, s := range sl {
		if s == str {
			return true
		}
	}
	return false
}

// IsSliceContainsStrNocase returns true if the string exists in given slice, ignore case.
// 字符串切片是否含有指定的元素,忽略大小写
func IsSliceContainsStrNocase(sl []string, str string) bool {
	str = strings.ToLower(str)
	for _, s := range sl {
		if strings.ToLower(s) == str {
			return true
		}
	}
	return false
}

// IsSliceContainsInt64 returns true if the int64 exists in given slice.
// int64切片是否含有指定的元素
func IsSliceContainsInt64(sl []int64, i int64) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}

// IsSliceContainsUint returns true if the uint exists in given slice.
// uint切片是否含有指定的元素
func IsSliceContainsUint(sl []uint, i uint) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}

// IsSliceContainsUint16 returns true if the uint16 exists in given slice.
// uint16切片是否含有指定的元素
func IsSliceContainsUint16(sl []uint16, i uint16) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}

// 反转[]byte
func ReverseSliceBytes(b []byte) []byte {
	for from, to := 0, len(b)-1; from < to; from, to = from+1, to-1 {
		b[from], b[to] = b[to], b[from]
	}

	return b
}

// 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}
