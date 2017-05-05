package decoder

import (
	"github.com/tsg/gopacket/layers"
)

type ICMPv4 struct {
	Type     uint8  `json:"type"`
	Code     uint8  `json:code"`
	Checksum uint16 `json:"checksum"`
	Id       uint16 `json:"id"`
	Seq      uint16 `json:"seq"`
	Payload  []byte `json:"payload,omitempty"`
	ICMPContent
}

func NewICMPv4(icmp4 *layers.ICMPv4) (i *ICMPv4, pktType PktType) {
	pktType = PktTypeICMPv4
	i = &ICMPv4{}
	i.Type = uint8(icmp4.TypeCode >> 8)
	if i.Type == layers.ICMPv4TypeEchoReply || i.Type == layers.ICMPv4TypeEchoRequest {
		return nil, 0
	}
	i.Code = uint8(icmp4.TypeCode)
	i.Checksum = icmp4.Checksum
	i.Id = icmp4.Id
	i.Seq = icmp4.Seq
	i.Payload = icmp4.Payload
	//if i.Type == layers.ICMPv4TypeDestinationUnreachable && i.Code == 3 { // port unreachable
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
	//}
	return i, pktType
}
