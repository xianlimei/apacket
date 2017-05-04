package decoder

import (
	"github.com/tsg/gopacket/layers"
)

type ICMPv4 struct {
	TypeCode layers.ICMPv4TypeCode `json:"typecode"`
	Checksum uint16                `json:"checksum"`
	Id       uint16                `json:"id"`
	Seq      uint16                `json:"seq"`
	Payload  []byte                `json:"payload,omitempty"`
}

func NewICMPv4(icmp4 *layers.ICMPv4) (i *ICMPv4, pktType int8) {
	pktType = PktTypeICMPv4
	i = &ICMPv4{}
	i.TypeCode = icmp4.TypeCode
	i.Checksum = icmp4.Checksum
	i.Id = icmp4.Id
	i.Seq = icmp4.Seq
	i.Payload = icmp4.Payload
	return i, pktType
}
