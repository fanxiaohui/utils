package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

//const char *Buildtime(void);
//const char *VersionDate(void);
import "C"

var markStartTime time.Time = time.Now().Local()

const layout = "2006-01-02 15:04:05"

// 开机时间
func SetupTime() string {
	return markStartTime.Format(layout)
}

// 运行时间 秒数
func RunningTimeSeconds() uint64 {
	return uint64(time.Since(markStartTime) / time.Second)
}

// 运行时间 Largest time is 2540400:10:10
func RunningTime() string {
	// get second
	u := uint64(time.Since(markStartTime) / time.Second)

	second := u % 60
	u /= 60
	minutes := u % 60
	u /= 60

	return fmt.Sprintf("%02d:%02d:%02d", u, minutes, second)
}

/*编译时间 format : 2018-12-09 15:26:26*/
func BuildDateTime() string {
	return C.GoString(C.Buildtime())
}

/*编译版本 format : major.minor.fixed - 1.0.1 Beta*/
func Version(major, minor, fixed int, isBeta bool) string {
	version := []string{
		strconv.FormatInt(int64(major), 10),
		strconv.FormatInt(int64(minor), 10),
		strconv.FormatInt(int64(fixed), 10),
	}

	if isBeta {
		return "V" + strings.Join(version, ".") + " Beta"
	}

	return "V" + strings.Join(version, ".")
}

// 判断系统大小端
func IsMachineLittleEndian() bool {
	var i int16 = 0x0001

	u := unsafe.Pointer(&i)
	b := *((*byte)(u))

	return b == 0x01
}
