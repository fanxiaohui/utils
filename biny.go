package utils

// BreakUint16 拆分uint16 为高低字节
func BreakUint16(v uint16) (LoByte, HiByte byte) {
	return byte(v), byte(v >> 8)
}

// BreakUint32 拆分uint32 为4个字节
func BreakUint32(v uint32) (Byte0, Byte1, Byte2, Byte3 byte) {
	return byte(v), byte(v >> 8), byte(v >> 16), byte(v >> 24)
}

// BuildUint16 将高低字节组uint16
func BuildUint16(loByte, Hibyte byte) uint16 {
	return uint16(loByte) | uint16(Hibyte)<<8
}

// BuildUint32 将4个字节组成uint32
func BuildUint32(Byte0, Byte1, Byte2, Byte3 byte) uint32 {
	return uint32(Byte0) | (uint32(Byte1) << 8) | (uint32(Byte2) << 16) | (uint32(Byte3) << 24)
}
