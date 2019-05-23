package utils

import (
	"fmt"
	"time"
	"unsafe"
)

//const char *Buildtime(void);
//const char *VersionDate(void);
import "C"

var markStartTime = time.Now().Local()

const layout = "2006-01-02 15:04:05"

// SetupTime 开机时间
func SetupTime() string {
	return markStartTime.Format(layout)
}

// RunningTimeSeconds 运行时间 秒数
func RunningTimeSeconds() uint64 {
	return uint64(time.Since(markStartTime) / time.Second)
}

// RunningTime 运行时间 Largest time is 2540400:10:10
func RunningTime() string {
	// get second
	u := uint64(time.Since(markStartTime) / time.Second)

	second := u % 60
	u /= 60
	minutes := u % 60
	u /= 60

	return fmt.Sprintf("%02d:%02d:%02d", u, minutes, second)
}

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
