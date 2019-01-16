package common

import "strings"

// AppendStr appends string to slice with no duplicates.
func AppendStr(strs []string, str string) []string {
	for _, s := range strs {
		if s == str {
			return strs
		}
	}
	return append(strs, str)
}

// AppendUint appends uint to slice with no duplicates.
func AppendUint(s []uint, e uint) []uint {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// AppendUint16 appends uint16 to slice with no duplicates.
func AppendUint16(s []uint16, e uint16) []uint16 {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// 删除uint16切片中的第一个指定元素
func DeleteFromSliceUint16(s []uint16, e uint16) []uint16 {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}

	return s
}

// 删除uint16切片中的所有指示元素
func DeleteFromSliceUint16All(s []uint16, e uint16) []uint16 {
	var tmpS []uint16

	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}

	return tmpS
}

// 删除uint切片中的第一个指定元素
func DeleteFromSliceUint(s []uint, e uint) []uint {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}

	return s
}

// 删除uint切片中的所有指示元素
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

// IsSliceContainsStr returns true if the string exists in given slice, ignore case.
func IsSliceContainsStr(sl []string, str string) bool {
	str = strings.ToLower(str)
	for _, s := range sl {
		if strings.ToLower(s) == str {
			return true
		}
	}
	return false
}

// IsSliceContainsInt64 returns true if the int64 exists in given slice.
func IsSliceContainsInt64(sl []int64, i int64) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}

// IsSliceContainsUint returns true if the uint exists in given slice.
func IsSliceContainsUint(sl []uint, i uint) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}

// IsSliceContainsUint16 returns true if the uint16 exists in given slice.
func IsSliceContainsUint16(sl []uint16, i uint16) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}
