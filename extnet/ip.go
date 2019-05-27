// 扩展net包, 提供ip到数值的转换

package extnet

import (
	"encoding/binary"
	"errors"
	"net"
)

// IP2Numer net.IP转换为数值
func IP2Numer(p net.IP) (uint32, error) {
	ip := p.To4()
	if ip == nil || len(ip) < 4 {
		return 0, errors.New("invalid ipv4 address")
	}

	return binary.BigEndian.Uint32(ip), nil
}

// Numer2IP 数值转换为net.IP
func Numer2IP(l uint32) net.IP {
	return net.IPv4(byte(l>>24), byte(l>>16), byte(l>>8), byte(l))
}

// Dot2Numer 点分十进制字符串转换数值
func Dot2Numer(s string) (uint32, error) {
	return Net2Numer(net.ParseIP(s))
}

// Number2Dot 数值转换为点分十进制字符串
func Number2Dot(l uint32) string {
	return Numer2Net(l).String()
}

// Mask2Dot mask 转为点分十进制作字符串
func Mask2Dot(mask net.IPMask) string {
	return net.IP(mask).String()
}

// Dot2Mask parses s as an IP address, returning the result.
// The string s can be in dotted decimal ("192.0.2.1")
// or IPv6 ("2001:db8::68") form.
// If s is not a valid textual representation of an IP address,
// ParseIP returns nil.
func Dot2Mask(s string) net.IPMask {
	return net.IPMask(net.ParseIP(s))
}
