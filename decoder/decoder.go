package decoder

import (
	"github.com/tsg/gopacket"
	"github.com/tsg/gopacket/layers"
	"time"
)

const (
	PktTypeTCP    int8 = 1
	PktTypeUDP    int8 = 2
	PktTypeDNS    int8 = 3
	PktTypeICMPv4 int8 = 4
	PktTypeICMPv6 int8 = 5
)

type Packet struct {
	Ts      time.Time `json:"ts"`
	Ip4     *IPv4     `json:"ip4,omitempty"`
	Ip6     *IPv6     `json:"ip6,omitempty"`
	Tcp     *TCP      `json:"tcp,omitempty"`
	Udp     *UDP      `json:"udp,omitempty"`
	Dns     *DNS      `json:"dns,omitempty"`
	Icmp4   *ICMPv4   `json:"icmp4,omitempty"`
	Icmp6   *ICMPv6   `json:"icmp6,omitempty"`
	PktType int8      `json:"ptype,omitempty"`
}

type Decoder struct {
}

func (d *Decoder) Process(data []byte, ci *gopacket.CaptureInfo) (*Packet, error) {

	pkt := &Packet{Ts: ci.Timestamp}
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
		case layers.LayerTypeIPv6:
			ip6l := packet.Layer(layers.LayerTypeIPv6)
			ip6, ok := ip6l.(*layers.IPv6)
			if !ok {
				return nil, nil
			}
			pkt.Ip6 = NewIP6(ip6)
		case layers.LayerTypeICMPv4:
			icmp4l := packet.Layer(layers.LayerTypeICMPv4)
			icmp4, ok := icmp4l.(*layers.ICMPv4)
			if !ok {
				break
			}
			pkt.Icmp4 = NewICMPv4(icmp4)
			pkt.PktType = PktTypeICMPv4
			return pkt, nil
		case layers.LayerTypeICMPv6:
			icmp6l := packet.Layer(layers.LayerTypeICMPv6)
			icmp6, ok := icmp6l.(*layers.ICMPv6)
			if !ok {
				break
			}
			pkt.Icmp6 = NewICMPv6(icmp6)
			pkt.PktType = PktTypeICMPv6
			return pkt, nil
		case layers.LayerTypeUDP:
			udpl := packet.Layer(layers.LayerTypeUDP)
			udp, ok := udpl.(*layers.UDP)
			if !ok {
				break
			}
			pkt.Udp = NewUDP(udp)
			pkt.PktType = PktTypeUDP
			//return pkt, nil
		case layers.LayerTypeDNS:
			dnsl := packet.Layer(layers.LayerTypeDNS)
			dns, ok := dnsl.(*layers.DNS)
			if !ok {
				break
			}
			pkt.Dns = NewDNS(dns)
			pkt.PktType = PktTypeDNS
			return pkt, nil
		case layers.LayerTypeTCP:
			tcpl := packet.Layer(layers.LayerTypeTCP)
			tcp, ok := tcpl.(*layers.TCP)
			if !ok {
				break
			}
			//if !tcp.SYN && !(tcp.SYN && tcp.ACK) {
			//	break
			//}
			pkt.Tcp = NewTCP(tcp)
			pkt.PktType = PktTypeTCP
			return pkt, nil
		}
	}
	if pkt.PktType != 0 {
		return pkt, nil
	}
	return nil, nil
}
