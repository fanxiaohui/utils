package utils

import (
	//"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
