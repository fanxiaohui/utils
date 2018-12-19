package common

import "encoding/binary"

type little_Endian struct{}
type big_Endian struct{}

// LittleEndian is the little-endian implementation of ByteOrder.
var Little_Endian little_Endian

// BigEndian is the big-endian implementation of ByteOrder.
var Big_Endian big_Endian

func BreakUint16(v uint16) (LoByte, HiByte byte) {
	return byte(v), byte(v >> 8)
}

func BreakUint32(v uint32) (Byte0, Byte1, Byte2, Byte3 byte) {
	return byte(v), byte(v >> 8), byte(v >> 16), byte(v >> 24)
}

func BuildUint16(loByte, Hibyte byte) uint16 {
	return uint16(loByte) | uint16(Hibyte)<<8
}

func BuildUint32(Byte0, Byte1, Byte2, Byte3 byte) uint32 {
	return uint32(Byte0) | (uint32(Byte1) << 8) | (uint32(Byte2) << 16) | (uint32(Byte3) << 24)
}

func (little_Endian) Putuint16(v uint16) []byte {
	s := make([]byte, 2)

	binary.LittleEndian.PutUint16(s, v)

	return s
}

func (little_Endian) Putuint32(v uint32) []byte {
	s := make([]byte, 4)

	binary.LittleEndian.PutUint32(s, v)

	return s
}

func (little_Endian) Putuint64(v uint64) []byte {
	s := make([]byte, 8)

	binary.LittleEndian.PutUint64(s, v)

	return s
}

func (big_Endian) Putuint16(v uint16) []byte {
	s := make([]byte, 2)

	binary.BigEndian.PutUint16(s, v)

	return s
}

func (big_Endian) Putuint32(v uint32) []byte {
	s := make([]byte, 4)

	binary.BigEndian.PutUint32(s, v)

	return s
}

func (big_Endian) Putuint64(v uint64) []byte {
	s := make([]byte, 8)

	binary.BigEndian.PutUint64(s, v)

	return s
}