package decoder

import (
	"github.com/tsg/gopacket/layers"
)

type ICMPv6 struct {
	TypeCode  layers.ICMPv6TypeCode `json:"typecode"`
	Checksum  uint16                `json:"checksum"`
	TypeBytes []byte                `json:"type,omitempty"`
	Payload   []byte                `json:"payload,omitempty"`
}

func NewICMPv6(icmp6 *layers.ICMPv6) (i *ICMPv6, pktType int8) {
	pktType = PktTypeICMPv6
	i = &ICMPv6{}
	i.TypeCode = icmp6.TypeCode
	i.Checksum = icmp6.Checksum
	i.TypeBytes = icmp6.TypeBytes
	i.Payload = icmp6.Payload
	return i, pktType
}
