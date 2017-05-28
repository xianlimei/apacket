package decoder

import (
	"github.com/tsg/gopacket"
	"github.com/tsg/gopacket/layers"
)

const (
	ipv4 uint8 = 4
	ipv6 uint8 = 6
)

type ICMPContent struct {
	Ip4 *IPv4 `json:"ip4,omitempty"`
	Ip6 *IPv6 `json:"ip6,omitempty"`
	Tcp *TCP  `json:"tcp,omitempty"`
	Udp *UDP  `json:"udp,omitempty"`
	Dns *DNS  `json:"dns,omitempty"`
}

func DecoderICMP(data []byte, ipVersion uint8) (*ICMPContent, PktType) {
	var pktType PktType

	var packet gopacket.Packet
	if ipVersion == ipv4 {
		packet = gopacket.NewPacket(data, layers.LayerTypeIPv4, gopacket.NoCopy)
	} else if ipVersion == ipv6 {
		packet = gopacket.NewPacket(data, layers.LayerTypeIPv6, gopacket.NoCopy)
	} else {
		return nil, pktType
	}

	icmpc := &ICMPContent{}
	for _, layer := range packet.Layers() {
		switch layer.LayerType() {
		case layers.LayerTypeIPv4:
			ip4l := packet.Layer(layers.LayerTypeIPv4)
			ip4, ok := ip4l.(*layers.IPv4)
			if !ok {
				return nil, 0
			}
			icmpc.Ip4 = NewIP4(ip4)
		case layers.LayerTypeIPv6:
			ip6l := packet.Layer(layers.LayerTypeIPv6)
			ip6, ok := ip6l.(*layers.IPv6)
			if !ok {
				return nil, 0
			}
			icmpc.Ip6 = NewIP6(ip6)
		case layers.LayerTypeDNS:
			dnsl := packet.Layer(layers.LayerTypeDNS)
			dns, ok := dnsl.(*layers.DNS)
			if !ok {
				break
			}
			icmpc.Dns, _ = NewDNS(dns)
			pktType = PktTypeICMPDNS
			icmpc.Udp.Payload = nil
			return icmpc, pktType
		case layers.LayerTypeTCP:
			tcpl := packet.Layer(layers.LayerTypeTCP)
			tcp, ok := tcpl.(*layers.TCP)
			if !ok {
				break
			}

			icmpc.Tcp, pktType = NewTCP(tcp)
			if pktType == PktTypeTCPSYN {
				pktType = PktTypeICMPTCPSYN
			} else if pktType == PktTypeTCPSYNACK {
				pktType = PktTypeICMPTCPSYNACK
			} else {
				pktType = PktTypeICMPTCP
			}
			return icmpc, pktType
		case layers.LayerTypeUDP:
			udpl := packet.Layer(layers.LayerTypeUDP)
			udp, ok := udpl.(*layers.UDP)
			if !ok {
				break
			}
			icmpc.Udp, _ = NewUDP(udp)
			pktType = PktTypeICMPUDP
		}
	}
	if pktType != 0 {
		return icmpc, pktType
	}
	return nil, 0
}
