package decoder

import (
	"fmt"
	"github.com/tsg/gopacket"
	"github.com/tsg/gopacket/layers"
)

const (
	PktTypeTCP    int8 = 1
	PktTypeUDP    int8 = 2
	PktTypeICMPv4 int8 = 3
)

type Decoder struct {
	Ip4     *IPv4   `json:"ip4,omitempty"`
	Tcp     *TCP    `json:"tcp,omitempty"`
	Udp     *UDP    `json:"udp,omitempty"`
	Icmp4   *ICMPv4 `json:"icmp4,omitempty"`
	PktType int8    `json:"ptype,omitempty"`
}

func (d *Decoder) Process(data []byte) {

	packet := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.NoCopy)
	for _, layer := range packet.Layers() {
		switch layer.LayerType() {
		case layers.LayerTypeIPv4:
			ip4l := packet.Layer(layers.LayerTypeIPv4)
			ip4, ok := ip4l.(*layers.IPv4)
			if !ok {
				fmt.Println("ip4 not ok")
				break
			}
			i := NewIP4(ip4)
			d.Ip4 = i
		case layers.LayerTypeIPv6:
			//TODO
		case layers.LayerTypeICMPv4:
			icmp4l := packet.Layer(layers.LayerTypeICMPv4)
			icmp4, ok := icmp4l.(*layers.ICMPv4)
			if !ok {
				fmt.Println("icmp4 not ok")
				break
			}
			ic4 := NewICMPv4(icmp4)
			d.Icmp4 = ic4
			d.PktType = PktTypeICMPv4
		case layers.LayerTypeICMPv6:
			//TODO
		case layers.LayerTypeUDP:
			udpl := packet.Layer(layers.LayerTypeUDP)
			udp, ok := udpl.(*layers.UDP)
			if !ok {
				fmt.Println("udp not ok")
				break
			}
			u := NewUDP(udp)
			d.Udp = u
			d.PktType = PktTypeUDP
		case layers.LayerTypeTCP:
			tcpl := packet.Layer(layers.LayerTypeTCP)
			tcp, ok := tcpl.(*layers.TCP)
			if !ok {
				fmt.Println("tcp not ok")
				break
			}
			if !tcp.SYN && !tcp.ACK {
				break
			}
			t := NewTCP(tcp)
			d.Tcp = t
			d.PktType = PktTypeTCP
		}
	}
}
