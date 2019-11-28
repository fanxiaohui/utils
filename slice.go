package utils

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

// AppendInt64 appends int64 to slice with no duplicates.
// 追加int16到无重复项的切片
func AppendInt64(s []int64, e int64) []int64 {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}

// DeleteFromSliceUint16 删除uint16切片中的 第一个出现的指定元素
func DeleteFromSliceUint16(s []uint16, e uint16) []uint16 {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// DeleteFromSliceUint16All 删除uint16切片中的 所有出现的指示元素
func DeleteFromSliceUint16All(s []uint16, e uint16) []uint16 {
	if s == nil {
		return s
	}

	tmpS := make([]uint16, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}
	return tmpS
}

// DeleteFromSliceUint 删除uint切片中的 第一个出现的指定元素
func DeleteFromSliceUint(s []uint, e uint) []uint {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// DeleteFromSliceUintAll 删除uint切片中的 所有出现的指示元素
func DeleteFromSliceUintAll(s []uint, e uint) []uint {
	if s == nil {
		return s
	}

	tmpS := make([]uint, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}

	return tmpS
}

// DeleteFromSliceInt64 删除int64切片中的 第一个出现的指定元素
func DeleteFromSliceInt64(s []int64, e int64) []int64 {
	for i, v := range s {
		if v == e {
			return append(s[:i], s[i+1:]...)
		}
	}

	return s
}

// DeleteFromSliceInt64All 删除int64切片中的 所有出现的指示元素
func DeleteFromSliceInt64All(s []int64, e int64) []int64 {
	if s == nil {
		return s
	}
	tmpS := make([]int64, 0, len(s))
	for _, v := range s {
		if v != e {
			tmpS = append(tmpS, v)
		}
	}

	return tmpS
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

// ReverseSliceBytes 反转[]byte
func ReverseSliceBytes(b []byte) []byte {
	for from, to := 0, len(b)-1; from < to; from, to = from+1, to-1 {
		b[from], b[to] = b[to], b[from]
	}

	return b
}
