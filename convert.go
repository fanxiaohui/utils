package common

import (
	"fmt"
	"strconv"
)

type argInt []int

// 返回参数索引的值,否则返回给定默认值
func (this argInt) GetMust(i, defVal int) int {
	var rt int = defVal

	if i >= 0 && i < len(this) {
		rt = this[i]
	}

	return rt
}

// 格式基本数据类型的字符串,如果给定无给定参数,则采用默认方式格式化
// 对于正数,则为默认十进制
// 对于float32,float64,参数为精度, 精度默认为 -1, 则代表使用最少数量的、但又必需的数字来表示。
// 函数将浮点数表示为字符串(十进制)并返回。
func FormatBaseTypes(val interface{}, args ...int) (s string) {
	switch v := val.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case *bool:
		s = strconv.FormatBool(*v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).GetMust(0, -1), 32)
	case *float32:
		s = strconv.FormatFloat(float64(*v), 'f', argInt(args).GetMust(0, -1), 32)
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).GetMust(0, -1), 64)
	case *float64:
		s = strconv.FormatFloat(*v, 'f', argInt(args).GetMust(0, -1), 64)
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).GetMust(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).GetMust(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).GetMust(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).GetMust(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).GetMust(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).GetMust(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).GetMust(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).GetMust(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).GetMust(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).GetMust(0, 10))
	case *int:
		s = strconv.FormatInt(int64(*v), argInt(args).GetMust(0, 10))
	case *int8:
		s = strconv.FormatInt(int64(*v), argInt(args).GetMust(0, 10))
	case *int16:
		s = strconv.FormatInt(int64(*v), argInt(args).GetMust(0, 10))
	case *int32:
		s = strconv.FormatInt(int64(*v), argInt(args).GetMust(0, 10))
	case *int64:
		s = strconv.FormatInt(int64(*v), argInt(args).GetMust(0, 10))
	case *uint:
		s = strconv.FormatUint(uint64(*v), argInt(args).GetMust(0, 10))
	case *uint8:
		s = strconv.FormatUint(uint64(*v), argInt(args).GetMust(0, 10))
	case *uint16:
		s = strconv.FormatUint(uint64(*v), argInt(args).GetMust(0, 10))
	case *uint32:
		s = strconv.FormatUint(uint64(*v), argInt(args).GetMust(0, 10))
	case *uint64:
		s = strconv.FormatUint(uint64(*v), argInt(args).GetMust(0, 10))
	case string:
		s = v
	case *string:
		s = *v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}

	return s
}

const hexDigit = "0123456789abcdef"

// 转化为16进制字符串, 不含前缀0x
func FormatSliceHex(a []byte) string {
	if len(a) == 0 {
		return ""
	}
	buf := make([]byte, 0, len(a)*2)
	for _, b := range a {
		buf = append(buf, hexDigit[b>>4])
		buf = append(buf, hexDigit[b&0xF])
	}

	return string(buf)
}

// Convert different types to byte slice using types and functions in unsafe and reflect package.
// It has higher performance, but notice that it may be not safe when garbage collection happens.
// Use it when you need to temporary convert a long string to a byte slice and won't keep it for long time.
// func Str2ByteSliceNonCopy(val string) []byte {
//     pslc := (*reflect.SliceHeader)(unsafe.Pointer(&val))
//     pslc.Cap = pslc.Len
//     return *(*[]byte)(unsafe.Pointer(pslc))
// }

// Zero-copy convert from byte slice to a string
// func BytesSlice2StrNonCopy(val []byte) string {
//     return *(*string)(unsafe.Pointer(&val))
// }
