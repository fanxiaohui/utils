package utils

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
	Convey("运时时间格式和最大值: 2540400:10:10", t, func() {
		println(RunningTime())
	})
	Convey("运时j时秒数", t, func() {
		println(RunningTimeSeconds())
	})

}

func TestBuildTime(t *testing.T) {
	Convey("编译时间格式: 2018-12-09 15:26:26", t, func() {
		println(BuildDateTime())
	})
}

func TestIsMachineLittleEndian(t *testing.T) {
	Convey("判断系统大小端 -- windows小端", t, func() {
		So(IsMachineLittleEndian(), ShouldBeTrue)
	})
}
