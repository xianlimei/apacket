package utils

import (
	"net"
)

func InterfaceAddrsByName(ifaceName string) ([]string, error) {

	var buf []string

	i, err := net.InterfaceByName(ifaceName)
	if err != nil {
		return nil, err
	}
	addrs, err := i.Addrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
			buf = append(buf, ip.String())
		case *net.IPAddr:
			ip = v.IP
			buf = append(buf, ip.String())
		}
	}
	return buf, nil
}
