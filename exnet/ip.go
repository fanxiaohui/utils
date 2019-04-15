// 扩展net包, 提供ip到数值的转换

package exnet

import (
	"errors"
	"math"
	"net"
)

// net.IP转换为数值
func ParseIP2L(p net.IP) (uint, error) {
	ip := p.To4()
	if ip == nil || len(ip) < 4 {
		return 0, errors.New("invalid ipv4 address")
	}

	return (uint(ip[3]) | uint(ip[2])<<8 | uint(ip[1])<<16 | uint(ip[0])<<24), nil
}

// 数值转换为net.IP
func ParseL2IP(l uint) (net.IP, error) {
	if l > math.MaxUint32 {
		return nil, errors.New("out of range")
	}

	return net.IPv4(byte(l>>24), byte(l>>16), byte(l>>8), byte(l)), nil
}

// 点分十进制字符串转换数值
func ParseString2L(s string) (uint, error) {
	return ParseIP2L(net.ParseIP(s))
}

// 数值转换为点分十进制字符串
func ParseL2String(l uint) (string, error) {
	ip, err := ParseL2IP(l)
	if err != nil {
		return "", err
	}

	return ip.String(), nil
}
