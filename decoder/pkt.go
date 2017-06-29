package decoder

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"io"
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
	PktTypeTCPACK        PktType = 13
)

type Packet struct {
	Ts          time.Time `json:"ts"`
	PktType     PktType   `json:"-"`
	Ptype       string    `json:"ptype,omitempty"`
	Host        string    `json:"host,omitempty"`
	Ip4         *IPv4     `json:"ip4,omitempty"`
	Ip6         *IPv6     `json:"ip6,omitempty"`
	Tcp         *TCP      `json:"tcp,omitempty"`
	Udp         *UDP      `json:"udp,omitempty"`
	Dns         *DNS      `json:"dns,omitempty"`
	Icmp4       *ICMPv4   `json:"icmp4,omitempty"`
	Icmp6       *ICMPv6   `json:"icmp6,omitempty"`
	Flow        *Flow     `json:"-"`
	PayloadSha1 string    `json:"psha1,omitempty"`
}

func (pkt *Packet) CompressPayload(payload []byte) bytes.Buffer {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(payload)
	w.Close()
	return b
}

func (pkt *Packet) PayloadSha1Hex(payload string) string {
	h := sha1.New()
	io.WriteString(h, payload)
	return hex.EncodeToString(h.Sum(nil))
}

func (pkt *Packet) CalPayloadSha1() string {
	var pl []byte

	switch pkt.PktType {
	case PktTypeTCPACK:
		pl = pkt.Tcp.Payload
	//case PktTypeTCPSYNACK:
	//	pl = pkt.Tcp.Payload
	case PktTypeTCP:
		pl = pkt.Tcp.Payload
	case PktTypeUDP:
		pl = pkt.Udp.Payload
	}
	if len(pl) != 0 {
		cPayload := pkt.CompressPayload(pl)
		return pkt.PayloadSha1Hex(cPayload.String())
	}
	return ""
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
	case PktTypeTCPACK:
		typeStr = "ack"
	case PktTypeICMPDNS:
		typeStr = "icmpdns"
	case PktTypeICMPTCP:
		typeStr = "icmptcp"
	case PktTypeICMPTCPSYN:
		typeStr = "icmpsyn"
	case PktTypeICMPTCPSYNACK:
		typeStr = "icmpsynack"
	case PktTypeICMPUDP:
		typeStr = "icmpudp"
	}
	return typeStr
}
