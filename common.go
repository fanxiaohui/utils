package utils

import (
	"unsafe"
)

//const char *Buildtime(void);
//const char *VersionDate(void);
import "C"

// BuildDateTime 编译时间 format : 2018-12-09 15:26:26
func BuildDateTime() string {
	return C.GoString(C.Buildtime())
}

// IsMachineLittleEndian 判断系统大小端
func IsMachineLittleEndian() bool {
	var i int16 = 0x0001

	u := unsafe.Pointer(&i)
	b := *((*byte)(u))

	return b == 0x01
}
