package common

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
