package common

import (
    "fmt"
    "strconv"
)

// Convert Basic Types to string(十进制), return empty string if the convertion fails (bool,float32,float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64,*bool,*float32,*float64, *int, *uint, *int8, *uint8, *int16, *uint16, *int32, *uint32, *int64, *uint64)
// 对于float32,float64 精度默认为 -1
func FormatBaseTypes(val interface{}) string {
    switch v := val.(type) {
    case bool:
        return strconv.FormatBool(v)
    case *bool:
        return strconv.FormatBool(*v)
    case float32, *float32, float64, *float64:
        return FormatFloat(v, -1)
    case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
        return fmt.Sprintf("%d", v)
    case *int:
        return strconv.FormatInt(int64(*v), 10)
    case *uint:
        return strconv.FormatUint(uint64(*v), 10)
    case *int8:
        return strconv.FormatInt(int64(*v), 10)
    case *uint8:
        return strconv.FormatUint(uint64(*v), 10)
    case *int16:
        return strconv.FormatInt(int64(*v), 10)
    case *uint16:
        return strconv.FormatUint(uint64(*v), 10)
    case *int32:
        return strconv.FormatInt(int64(*v), 10)
    case *uint32:
        return strconv.FormatUint(uint64(*v), 10)
    case *int64:
        return strconv.FormatInt(int64(*v), 10)
    case *uint64:
        return strconv.FormatUint(uint64(*v), 10)
    default:
        return ""
    }
}

// 函数将浮点数表示为字符串(十进制)并返回。
//prec控制精度（排除指数部分）：对十进制，它表示小数点后的数字个数,如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
func FormatFloat(val interface{}, prec int) string {
    switch v := val.(type) {
    case float32:
        return strconv.FormatFloat(float64(v), 'f', prec, 32)
    case *float32:
        return strconv.FormatFloat(float64(*v), 'f', prec, 32)
    case float64:
        return strconv.FormatFloat(v, 'f', prec, 64)
    case *float64:
        return strconv.FormatFloat(*v, 'f', prec, 64)
    default:
        return ""
    }
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
