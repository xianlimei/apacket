package decoder

import (
	"github.com/tsg/gopacket"
	"github.com/tsg/gopacket/layers"
)

type Decoder struct {
}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (d *Decoder) Process(data []byte, ci *gopacket.CaptureInfo) (*Packet, error) {

	flow := &Flow{}
	pkt := &Packet{Ts: ci.Timestamp,
		Flow: flow}

	defer func() {
		pkt.Ptype = pkt.PktType.String()
	}()

	packet := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.NoCopy)
	for _, layer := range packet.Layers() {
		switch layer.LayerType() {
		case layers.LayerTypeIPv4:
			ip4l := packet.Layer(layers.LayerTypeIPv4)
			ip4, ok := ip4l.(*layers.IPv4)
			if !ok {
				return nil, nil
			}
			pkt.Ip4 = NewIP4(ip4)

			flow.Sip = ip4.SrcIP
			flow.Dip = ip4.DstIP
			flow.Protocol = ip4.Protocol

		case layers.LayerTypeIPv6:
			ip6l := packet.Layer(layers.LayerTypeIPv6)
			ip6, ok := ip6l.(*layers.IPv6)
			if !ok {
				return nil, nil
			}
			pkt.Ip6 = NewIP6(ip6)

			flow.Sip = ip6.SrcIP
			flow.Dip = ip6.DstIP
			flow.Protocol = ip6.NextHeader

		case layers.LayerTypeICMPv4:
			icmp4l := packet.Layer(layers.LayerTypeICMPv4)
			icmp4, ok := icmp4l.(*layers.ICMPv4)
			if !ok {
				break
			}
			pkt.Icmp4, pkt.PktType = NewICMPv4(icmp4)
			if pkt.Icmp4 == nil {
				return nil, nil
			}
			return pkt, nil
		case layers.LayerTypeICMPv6:
			icmp6l := packet.Layer(layers.LayerTypeICMPv6)
			icmp6, ok := icmp6l.(*layers.ICMPv6)
			if !ok {
				break
			}
			pkt.Icmp6, pkt.PktType = NewICMPv6(icmp6)
			if pkt.Icmp6 == nil {
				return nil, nil
			}
			return pkt, nil
		case layers.LayerTypeUDP:
			udpl := packet.Layer(layers.LayerTypeUDP)
			udp, ok := udpl.(*layers.UDP)
			if !ok {
				break
			}
			pkt.Udp, pkt.PktType = NewUDP(udp)
			flow.Sport = uint16(udp.SrcPort)
			flow.Dport = uint16(udp.DstPort)
			//return pkt, nil
		case layers.LayerTypeDNS:
			dnsl := packet.Layer(layers.LayerTypeDNS)
			dns, ok := dnsl.(*layers.DNS)
			if !ok {
				break
			}
			pkt.Dns, pkt.PktType = NewDNS(dns)
			//pkt.Udp.Payload = nil
			return pkt, nil
		case layers.LayerTypeTCP:
			tcpl := packet.Layer(layers.LayerTypeTCP)
			tcp, ok := tcpl.(*layers.TCP)
			if !ok {
				break
			}
			pkt.Tcp, pkt.PktType = NewTCP(tcp)
			flow.Sport = uint16(tcp.SrcPort)
			flow.Dport = uint16(tcp.DstPort)
			return pkt, nil
		}
	}
	if pkt.PktType != 0 {
		return pkt, nil
	}
	return nil, nil
}
