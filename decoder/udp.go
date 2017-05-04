package decoder

import (
	"github.com/tsg/gopacket/layers"
)

type UDP struct {
	SrcPort  layers.UDPPort `json:"sport"`
	DstPort  layers.UDPPort `json:"dport"`
	Length   uint16         `json:"len"`
	Checksum uint16         `json:"checksum"`
	Payload  []byte         `json:"payload,omitempty"`
}

func NewUDP(udp *layers.UDP) (u *UDP, pktType PktType) {
	pktType = PktTypeUDP
	u = &UDP{}
	u.SrcPort = udp.SrcPort
	u.DstPort = udp.DstPort
	u.Checksum = udp.Checksum
	u.Payload = udp.Payload
	return u, pktType
}
