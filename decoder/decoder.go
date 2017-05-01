package decoder

import (
	"fmt"
	"github.com/tsg/gopacket"
	"github.com/tsg/gopacket/layers"
	"time"
)

const (
	PktTypeTCP    int8 = 1
	PktTypeUDP    int8 = 2
	PktTypeICMPv4 int8 = 3
)

type Packet struct {
	Ts      time.Time `json:"ts"`
	Ip4     *IPv4     `json:"ip4,omitempty"`
	Tcp     *TCP      `json:"tcp,omitempty"`
	Udp     *UDP      `json:"udp,omitempty"`
	Icmp4   *ICMPv4   `json:"icmp4,omitempty"`
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
				fmt.Println("ip4 not ok")
				return nil, nil
			}
			i := NewIP4(ip4)
			pkt.Ip4 = i
		case layers.LayerTypeIPv6:
			//TODO
			return nil, nil
		case layers.LayerTypeICMPv4:
			icmp4l := packet.Layer(layers.LayerTypeICMPv4)
			icmp4, ok := icmp4l.(*layers.ICMPv4)
			if !ok {
				fmt.Println("icmp4 not ok")
				break
			}
			ic4 := NewICMPv4(icmp4)
			pkt.Icmp4 = ic4
			pkt.PktType = PktTypeICMPv4
			return pkt, nil
		case layers.LayerTypeICMPv6:
			//TODO
			return nil, nil
		case layers.LayerTypeUDP:
			udpl := packet.Layer(layers.LayerTypeUDP)
			udp, ok := udpl.(*layers.UDP)
			if !ok {
				fmt.Println("udp not ok")
				break
			}
			u := NewUDP(udp)
			pkt.Udp = u
			pkt.PktType = PktTypeUDP
			return pkt, nil
		case layers.LayerTypeTCP:
			tcpl := packet.Layer(layers.LayerTypeTCP)
			tcp, ok := tcpl.(*layers.TCP)
			if !ok {
				fmt.Println("tcp not ok")
				break
			}
			if !tcp.SYN && !(tcp.SYN && tcp.ACK) {
				break
			}
			t := NewTCP(tcp)
			pkt.Tcp = t
			pkt.PktType = PktTypeTCP
			return pkt, nil
		}
	}
	return nil, nil
}
