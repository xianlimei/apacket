package dns

import (
	"github.com/Acey9/apacket/honeypot/core"
	"github.com/Acey9/apacket/logp"
	"github.com/tsg/gopacket"
	"github.com/tsg/gopacket/layers"
)

const (
	PtypeDNS = "dns"
)

type DNS struct {
}

func NewDNS() *DNS {
	return &DNS{}
}

func (dns *DNS) Fingerprint(request []byte, tlsTag bool) (identify bool, ptype string, err error) {
	ptype = PtypeDNS
	packet := gopacket.NewPacket(request, layers.LayerTypeDNS, gopacket.NoCopy)
	if dnsLayer := packet.Layer(layers.LayerTypeDNS); dnsLayer != nil {
		dnsPkt, ok := dnsLayer.(*layers.DNS)
		if !ok {
			return
		}
		if !dnsPkt.QR {
			identify = true
			logp.Debug("dns", "dns.QR:%v", dnsPkt.QR)
			return
		}
	}
	return
}

func (dns *DNS) Parser(remoteAddr, localAddr string, request []byte, ptype string, tls bool) (response *core.Applayer) {
	return
}

func (dns *DNS) DisguiserResponse(request []byte) (response []byte) {
	//TODO
	//response = []byte("\x00")
	return
}
