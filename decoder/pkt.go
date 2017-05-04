package decoder

import (
	"time"
)

const (
	PktTypeTCP            PktType = 1
	PktTypeUDP            PktType = 2
	PktTypeDNS            PktType = 3
	PktTypeICMPv4         PktType = 4
	PktTypeICMPv6         PktType = 5
	PktTypeTCPSYN         PktType = 6
	PktTypeTCPSYNACK      PktType = 7
	PktTypeICMP4DNS       PktType = 8
	PktTypeICMP4TCP       PktType = 9
	PktTypeICMP4TCPSYN    PktType = 10
	PktTypeICMP4TCPSYNACK PktType = 11
	PktTypeICMP4UDP       PktType = 12
)

type Packet struct {
	Ts      time.Time `json:"ts"`
	PktType PktType   `json:"-"`
	Ptype   string    `json:"ptype,omitempty"`
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
	case PktTypeICMP4DNS:
		typeStr = "icmp_dns"
	case PktTypeICMP4TCP:
		typeStr = "icmp_tcp"
	case PktTypeICMP4TCPSYN:
		typeStr = "icmp_syn"
	case PktTypeICMP4TCPSYNACK:
		typeStr = "icmp_synack"
	case PktTypeICMP4UDP:
		typeStr = "icmp_udp"
	}
	return typeStr
}
