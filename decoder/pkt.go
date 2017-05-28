package decoder

import (
	"time"
)

const (
	PktTypeTCP           PktType = 1
	PktTypeUDP           PktType = 2
	PktTypeDNS           PktType = 3
	PktTypeICMPv4        PktType = 4
	PktTypeICMPv6        PktType = 5
	PktTypeTCPSYN        PktType = 6
	PktTypeTCPSYNACK     PktType = 7
	PktTypeICMPDNS       PktType = 8
	PktTypeICMPTCP       PktType = 9
	PktTypeICMPTCPSYN    PktType = 10
	PktTypeICMPTCPSYNACK PktType = 11
	PktTypeICMPUDP       PktType = 12
)

type Packet struct {
	Ts      time.Time `json:"ts"`
	PktType PktType   `json:"-"`
	Ptype   string    `json:"ptype,omitempty"`
	Host    string    `json:"host,omitempty"`
	Ip4     *IPv4     `json:"ip4,omitempty"`
	Ip6     *IPv6     `json:"ip6,omitempty"`
	Tcp     *TCP      `json:"tcp,omitempty"`
	Udp     *UDP      `json:"udp,omitempty"`
	Dns     *DNS      `json:"dns,omitempty"`
	Icmp4   *ICMPv4   `json:"icmp4,omitempty"`
	Icmp6   *ICMPv6   `json:"icmp6,omitempty"`
	Flow    *Flow     `json:"-"`
}

type PktType uint8

func (pt PktType) String() string {
	var typeStr string
	switch pt {
	case PktTypeTCP:
		typeStr = "tcp"
	case PktTypeUDP:
		typeStr = "udp"
	case PktTypeDNS:
		typeStr = "dns"
	case PktTypeICMPv4:
		typeStr = "icmp4"
	case PktTypeICMPv6:
		typeStr = "icmp6"
	case PktTypeTCPSYN:
		typeStr = "syn"
	case PktTypeTCPSYNACK:
		typeStr = "synack"
	case PktTypeICMPDNS:
		typeStr = "icmpdns"
	case PktTypeICMPTCP:
		typeStr = "icmptcp"
	case PktTypeICMPTCPSYN:
		typeStr = "icmpsyn"
	case PktTypeICMPTCPSYNACK:
		typeStr = "icmpsynack"
	case PktTypeICMPUDP:
		typeStr = "icmp4_udp"
	}
	return typeStr
}
