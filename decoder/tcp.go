package decoder

import (
	"github.com/tsg/gopacket/layers"
)

type TCP struct {
	SrcPort    layers.TCPPort     `json:"sport"`
	DstPort    layers.TCPPort     `json:"dport"`
	Seq        uint32             `json:"seq"`
	Ack        uint32             `json:"acknum"`
	DataOffset uint8              `json:"offset"`
	FIN        bool               `json:"fin"`
	SYN        bool               `json:"syn"`
	RST        bool               `json:"rst"`
	PSH        bool               `json:"psh"`
	ACK        bool               `json:"ack"`
	URG        bool               `json:"urg"`
	ECE        bool               `json:"ece"`
	CWR        bool               `json:"cwr"`
	NS         bool               `json:"ns"`
	Window     uint16             `json:"window"`
	Checksum   uint16             `json:"checksum"`
	Urgent     uint16             `json:"urgent"`
	Options    []layers.TCPOption `json:"-"`
	Padding    []byte             `json:"padding,omitempty"`
	Payload    []byte             `json:"payload,omitempty"`
}

func NewTCP(tcp *layers.TCP) *TCP {
	t := &TCP{}
	t.SrcPort = tcp.SrcPort
	t.DstPort = tcp.DstPort
	t.Seq = tcp.Seq
	t.Ack = tcp.Ack
	t.DataOffset = tcp.DataOffset
	t.FIN = tcp.FIN
	t.SYN = tcp.SYN
	t.RST = tcp.RST
	t.PSH = tcp.PSH
	t.ACK = tcp.ACK
	t.URG = tcp.URG
	t.ECE = tcp.ECE
	t.CWR = tcp.CWR
	t.NS = tcp.NS
	t.Window = tcp.Window
	t.Checksum = tcp.Checksum
	t.Urgent = tcp.Urgent
	t.Options = tcp.Options
	t.Padding = tcp.Padding
	t.Payload = tcp.Payload
	return t
}
