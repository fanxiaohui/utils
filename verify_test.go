package common

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCrcChecksum(t *testing.T) {
	var tab = []byte{0x04, 0x00, 0x00, 0x00, 0x03, 0xb0, 0x0b}

	// Only pass t into top-level Convey calls
	Convey("crc校验", t, func() {
		So(0xc79a, ShouldEqual, CrcChecksum(tab))
	})
}

func TestXorCheckSum(t *testing.T) {
	var tab = []byte{0x04, 0x00, 0x00, 0x00, 0x03, 0xb0, 0x0b}

	// Only pass t into top-level Convey calls
	Convey("crc校验", t, func() {
		So(0xbc, ShouldEqual, XorCheckSum(tab))
	})
}
