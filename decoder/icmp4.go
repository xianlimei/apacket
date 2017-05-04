package decoder

import (
	"github.com/tsg/gopacket"
	"github.com/tsg/gopacket/layers"
)

type ICMPContent struct {
	Tcp *TCP `json:"tcp,omitempty"`
	Udp *UDP `json:"udp,omitempty"`
	Dns *DNS `json:"dns,omitempty"`
}

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
	if i.Type == layers.ICMPv4TypeDestinationUnreachable && i.Code == 3 { // port unreachable
		pkt, pt := DecoderDestinationUnreachable(i.Payload, i)
		if pkt != nil {
			pkt.Payload = nil
			return pkt, pt
		}
	}
	return i, pktType
}

func DecoderDestinationUnreachable(data []byte, icmp4 *ICMPv4) (*ICMPv4, PktType) {
	var pktType PktType
	pktType = 0
	packet := gopacket.NewPacket(data, layers.LayerTypeIPv4, gopacket.NoCopy)
	for _, layer := range packet.Layers() {
		switch layer.LayerType() {
		case layers.LayerTypeDNS:
			dnsl := packet.Layer(layers.LayerTypeDNS)
			dns, ok := dnsl.(*layers.DNS)
			if !ok {
				break
			}
			icmp4.Dns, _ = NewDNS(dns)
			pktType = PktTypeICMP4DNS
			icmp4.Udp.Payload = nil
			return icmp4, pktType
		case layers.LayerTypeTCP:
			tcpl := packet.Layer(layers.LayerTypeTCP)
			tcp, ok := tcpl.(*layers.TCP)
			if !ok {
				break
			}

			icmp4.Tcp, pktType = NewTCP(tcp)
			if pktType == PktTypeTCPSYN {
				pktType = PktTypeICMP4TCPSYN
			} else if pktType == PktTypeTCPSYNACK {
				pktType = PktTypeICMP4TCPSYNACK
			} else {
				pktType = PktTypeICMP4TCP
			}
			return icmp4, pktType
		case layers.LayerTypeUDP:
			udpl := packet.Layer(layers.LayerTypeUDP)
			udp, ok := udpl.(*layers.UDP)
			if !ok {
				break
			}
			icmp4.Udp, pktType = NewUDP(udp)
		}
	}
	if pktType != 0 {
		return icmp4, pktType
	}
	return nil, 0
}
