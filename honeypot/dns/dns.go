package dns

import (
	"github.com/Acey9/apacket/honeypot/core"
	"github.com/Acey9/apacket/logp"
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
	if len(request) < 31 {
		return
	}
	logp.Debug("dns.Fingerprint", "dns...")
	tid := request[0:2]
	flag := request[2:4]
	questions := request[4:6]
	anserRRS := request[6:8]
	authorityRRS := request[8:10]
	additionalRRS := request[10:12]
	logp.Debug("dns.Fingerprint", "tid:% 2x, flag:% 2x, questions:% 2x, anserRRS:% 2x, authorityRRS:% 2x, additionalRRS:% 2x", tid, flag, questions, anserRRS, authorityRRS, additionalRRS)
	queries := request[12:]
	logp.Debug("dns.Fingerprint", "queries:% 2x", queries)
	//TODO
	return
}

func (dns *DNS) Parser(remoteAddr, localAddr string, request []byte, ptype string, tls bool) (response *core.Applayer) {
	return
}

func (dns *DNS) DisguiserResponse(request []byte) (response []byte) {
	return
}
