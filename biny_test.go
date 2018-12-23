package common

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUint16(t *testing.T) {
	var val16 uint16 = 0x1234
	var lo8 byte = 0x34
	var hi8 byte = 0x12

	Convey("uint16分成高低字节", t, func() {
		tmpl, tmph := BreakUint16(val16)
		So(tmpl, ShouldEqual, lo8)
		So(tmph, ShouldEqual, hi8)
	})
	Convey("高低字节组成uint16", t, func() {
		So(BuildUint16(lo8, hi8), ShouldEqual, val16)
	})
}

func TestUint32(t *testing.T) {
	var val32 uint32 = 0x12345678
	var byte0 byte = 0x78
	var byte1 byte = 0x56
	var byte2 byte = 0x34
	var byte3 byte = 0x12
	Convey("uint32分成4字节", t, func() {
		bt0, bt1, bt2, bt3 := BreakUint32(val32)
		So(bt0, ShouldEqual, byte0)
		So(bt1, ShouldEqual, byte1)
		So(bt2, ShouldEqual, byte2)
		So(bt3, ShouldEqual, byte3)
	})
	Convey("4字节组成uint32", t, func() {
		So(BuildUint32(byte0, byte1, byte2, byte3), ShouldEqual, val32)
	})
}

func TestPutUint(t *testing.T) {
	s2 := []byte{1, 2}
	s4 := []byte{1, 2, 3, 4}
	s8 := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	Convey("uint16转换成切片<大小端>", t, func() {
		So(bytes.Equal(Little_Endian.Putuint16(0x0201), s2), ShouldBeTrue)

		So(bytes.Equal(Big_Endian.Putuint16(0x0102), s2), ShouldBeTrue)
	})

	Convey("uint32转换成切片<大小端>", t, func() {
		So(bytes.Equal(Little_Endian.Putuint32(0x04030201), s4), ShouldBeTrue)

		So(bytes.Equal(Big_Endian.Putuint32(0x01020304), s4), ShouldBeTrue)
	})

	Convey("uint64转换成切片<大小端>", t, func() {
		So(bytes.Equal(Little_Endian.Putuint64(0x0807060504030201), s8), ShouldBeTrue)

		So(bytes.Equal(Big_Endian.Putuint64(0x0102030405060708), s8), ShouldBeTrue)
	})

}
