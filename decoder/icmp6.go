package decoder

import (
	"github.com/tsg/gopacket/layers"
)

type ICMPv6 struct {
	Type      uint8  `json:"type"`
	Code      uint8  `json:code"`
	Checksum  uint16 `json:"checksum"`
	TypeBytes []byte `json:"type,omitempty"`
	Payload   []byte `json:"payload,omitempty"`
	ICMPContent
}

func NewICMPv6(icmp6 *layers.ICMPv6) (i *ICMPv6, pktType PktType) {
	pktType = PktTypeICMPv6
	i = &ICMPv6{}

	i.Type = uint8(icmp6.TypeCode >> 8)
	if i.Type == layers.ICMPv6TypeEchoRequest || i.Type == layers.ICMPv6TypeEchoReply {
		return nil, 0
	}
	i.Code = uint8(icmp6.TypeCode)

	i.Checksum = icmp6.Checksum
	i.TypeBytes = icmp6.TypeBytes
	i.Payload = icmp6.Payload
	icmpContent, pt := DecoderICMP(i.Payload, ipv4)
	if icmpContent != nil {
		pktType = pt
		//i.Payload = nil
		i.Ip4 = icmpContent.Ip4
		i.Ip6 = icmpContent.Ip6
		i.Tcp = icmpContent.Tcp
		i.Udp = icmpContent.Udp
		i.Dns = icmpContent.Dns
	}
	return i, pktType
}
