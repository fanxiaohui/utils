package common

import "testing"

func TestCrc(t *testing.T) {
	var tab = []byte{0x04, 0x00, 0x00, 0x00, 0x03, 0xb0, 0x0b}
	if CrcChecksum(tab) != 0xc79a {
		t.Log("failed")
	} else {
		t.Log("success")
	}
}
