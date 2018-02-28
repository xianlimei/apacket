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
	IPv         uint8     `json:"ipv"`
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
	Plen        uint      `json:"plen,omitempty"`
}

func (pkt *Packet) Compress(source []byte) bytes.Buffer {
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	w.Write(source)
	w.Close()
	return buf
}

func (pkt *Packet) Sha1HexDigest(str string) string {
	h := sha1.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

func (pkt *Packet) CompressPayload() (psha1 string, plen uint) {
	switch pkt.PktType {
	case PktTypeTCPACK:
		plen = uint(len(pkt.Tcp.Payload))
		if plen != 0 {
			psha1 = pkt.Sha1HexDigest(string(pkt.Tcp.Payload))
			cPayload := pkt.Compress(pkt.Tcp.Payload)
			pkt.Tcp.Payload = cPayload.Bytes()
		}
	//case PktTypeTCPSYNACK:
	//	pl = pkt.Tcp.Payload
	case PktTypeTCP:
		plen = uint(len(pkt.Tcp.Payload))
		if plen != 0 {
			psha1 = pkt.Sha1HexDigest(string(pkt.Tcp.Payload))
			cPayload := pkt.Compress(pkt.Tcp.Payload)
			pkt.Tcp.Payload = cPayload.Bytes()
		}
	case PktTypeUDP:
		plen = uint(len(pkt.Udp.Payload))
		if plen != 0 {
			psha1 = pkt.Sha1HexDigest(string(pkt.Udp.Payload))
			cPayload := pkt.Compress(pkt.Udp.Payload)
			pkt.Udp.Payload = cPayload.Bytes()
		}
	}
	return
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
