package common

import (
	//"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSetupTime(t *testing.T) {
	Convey("开机时间格式: 2006-01-02 15:04:05", t, func() {
		println(SetupTime())
	})
}

func TestRunningTime(t *testing.T) {
	Convey("运时时间格式与最大值: 2540400:10:10", t, func() {
		println(RunningTime())
	})
}

func TestBuildTime(t *testing.T) {
	Convey("编译时间格式: 2018-12-09 15:26:26", t, func() {
		println(BuildTime())
	})
}

func TestVersion(t *testing.T) {
	Convey("版本格式: 1.0.20181209", t, func() {
		println(Version("1", "2"))
	})
}

func TestIsMachineLittleEndian(t *testing.T) {
	Convey("判断系统大小端 -- windows小端", t, func() {
		So(IsMachineLittleEndian(), ShouldBeTrue)
	})
}
