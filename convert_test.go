package common

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFormatBaseTypes(t *testing.T) {
	ts := map[interface{}]string{
		true:            "true",
		false:           "false",
		int(8):          "8",
		int(-8):         "-8",
		float32(100.2):  "100.2",
		float32(-100.2): "-100.2",
		float64(100.2):  "100.2",
		float64(-100.2): "-100.2",
	}

	Convey("基本数据类型 转 字符串", t, func() {
		for k, v := range ts {
			So(strings.EqualFold(FormatBaseTypes(k), v), ShouldBeTrue)
		}
	})

	Convey("基本数据类型(指针) 转 字符串", t, func() {
		var tb bool = true
		var ptb *bool = &tb
		var ti16 int16 = 103
		var pti16 *int16 = &ti16
		var tu32 uint32 = 104
		var ptu32 *uint32 = &tu32
		var tf32 float32 = 108.1
		var ptf32 *float32 = &tf32
		var tf64 float64 = 109.2
		var ptf64 *float64 = &tf64

		So(strings.EqualFold(FormatBaseTypes(ptb), "true"), ShouldBeTrue)
		So(strings.EqualFold(FormatBaseTypes(pti16), "103"), ShouldBeTrue)
		So(strings.EqualFold(FormatBaseTypes(ptu32), "104"), ShouldBeTrue)
		So(strings.EqualFold(FormatBaseTypes(ptf32, -1), "108.1"), ShouldBeTrue)
		So(strings.EqualFold(FormatBaseTypes(ptf64), "109.2"), ShouldBeTrue)
	})
}

func TestFormatSliceHex(t *testing.T) {
	Convey("转化为16进制字符串,不含前缀0x, 切片有值", t, func() {
		hex := []byte{0x12, 0x34, 0x56, 0x78}
		So(FormatSliceHex(hex), ShouldEqual, "12345678")
	})
	Convey("转化为16进制字符串,不含前缀0x, 切片无值", t, func() {
		hex := []byte{}
		So(FormatSliceHex(hex), ShouldEqual, "")
	})
}
