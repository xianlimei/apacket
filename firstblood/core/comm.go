package core

import (
	"bytes"
	"encoding/binary"
	"net"
	"strconv"
	"strings"
)

func IntranetIP(ipStr string) bool {
	var ipInt uint32
	ip, _, err := net.ParseCIDR(ipStr + "/32")
	if err != nil {
		return false
	}

	if ip.IsLoopback() {
		return true
	}

	if len(ip) == 16 {
		ipInt = binary.BigEndian.Uint32(ip[12:16])
	} else {
		ipInt = binary.BigEndian.Uint32(ip)
	}

	//192.168.0.0/16
	if ipInt >= 3232235520 && ipInt <= 3232301055 {
		return true
	}

	//10.0.0.0/8
	if ipInt >= 167772160 && ipInt <= 184549375 {
		return true
	}

	//172.16.0.0/12
	if ipInt >= 2886729728 && ipInt <= 2887778303 {
		return true
	}
	return false
}

func GetIPPort(addr string) (ip string, port uint16, ipv int, err error) {
	var iPort int
	ipPort := strings.Split(addr, ":")
	_len := len(ipPort)
	if _len > 2 {
		ip = strings.Join(ipPort[0:_len-1], ":")
		iPort, err = strconv.Atoi(ipPort[_len-1])
		ipv = IPv6
	} else {
		ip = ipPort[0]
		iPort, err = strconv.Atoi(ipPort[_len-1])
		ipv = IPv4
	}
	if err == nil {
		port = uint16(iPort)
	}
	return
}

func ParseInt(line []byte) (int, error) {
	i, err := strconv.Atoi(string(line))
	return int(i), err
	// TODO: is it an error if 'buf.Len() != 0 {}' ?
}

func Trim(buf []byte) []byte {
	return TrimLeft(TrimRight(buf))
}

func TrimLeft(buf []byte) []byte {
	for i, b := range buf {
		if b != ' ' && b != '\t' {
			return buf[i:]
		}
	}
	return nil
}

func TrimRight(buf []byte) []byte {
	for i := len(buf) - 1; i > 0; i-- {
		b := buf[i]
		if b != ' ' && b != '\t' {
			return buf[:i+1]
		}
	}
	return nil
}

func ToLower(buf, in []byte) []byte {
	if len(in) > len(buf) {
		goto unbufferedToLower
	}

	for i, b := range in {
		if b > 127 {
			goto unbufferedToLower
		}

		if 'A' <= b && b <= 'Z' {
			b = b - 'A' + 'a'
		}
		buf[i] = b
	}
	return buf[:len(in)]

unbufferedToLower:
	return bytes.ToLower(in)
}
