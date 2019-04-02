package utils

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFormatBaseTypes(t *testing.T) {
	ts := []struct {
		got  interface{}
		want string
	}{
		{true, "true"},
		{float32(100.2), "100.2"},
		{float64(100.2), "100.2"},
		{int(-8), "-8"},
		{int8(8), "8"},
		{int16(-3000), "-3000"},
		{int32(50000), "50000"},
		{int64(4324), "4324"},
		{uint(190), "190"},
		{uint8(8), "8"},
		{uint16(3000), "3000"},
		{uint32(5000000), "5000000"},
		{uint64(4324), "4324"},
		{"hello", "hello"},
		{[]byte{'1', '2', '3'}, "123"},
	}

	Convey("基本数据类型 转 字符串", t, func() {
		for _, v := range ts {
			So(strings.EqualFold(FormatBaseTypes(v.got), v.want), ShouldBeTrue)
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
