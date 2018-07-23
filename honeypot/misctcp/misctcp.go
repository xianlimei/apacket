package misctcp

import (
	"bytes"
	"fmt"
	"github.com/Acey9/apacket/honeypot/core"
	"unicode"
)

const respCiscoJabber = "<?xml version='1.0'?><stream:stream xmlns:stream='http://etherx.jabber.org/streams' xmlns='jabber:client' " +
	"from=%s version='1.0' id='ro7kfeld3ecro6v2yngoihvwx9vuxcco0k9fyut2' xmlns:ack='http://www.xmpp.org/extensions/xep-0198.html#ns'>" +
	"<stream:features xmlns:stream='http://etherx.jabber.org/streams'><starttls xmlns='urn:ietf:params:xml:ns:xmpp-tls'>" +
	"<required/></starttls><address xmlns='http://affinix.com/jabber/address'>1.1.1.2</address>" +
	"<auth xmlns='http://jabber.org/features/iq-auth'/></stream:features>"

type Misc struct {
}

func NewMisc() *Misc {
	u := &Misc{}
	return u
}

func (s *Misc) Fingerprint(request []byte, tlsTag bool) (identify bool, ptype string, err error) {
	return
}

func (s *Misc) Parser(remoteAddr, localAddr string, request []byte, ptype string, tls bool) (response *core.Applayer) {
	return
}

func (s *Misc) ciscoJabber(request []byte) (response []byte) {
	reqLen := len(request)
	i := bytes.Index(request, []byte("to="))
	if i == -1 {
		return
	}
	if i+3 >= reqLen {
		return
	}
	request = request[i+3:]

	spaceIdx := bytes.IndexFunc(request, unicode.IsSpace)
	if spaceIdx == -1 || spaceIdx < 3 {
		return
	}
	to := request[:spaceIdx]
	res := fmt.Sprintf(respCiscoJabber, to)
	response = []byte(res)
	return
}

func (s *Misc) weblogicRCE(request []byte) (response []byte) {
	//CVE-2018-2893
	i := bytes.Index(request, []byte("t3 "))
	if i != 0 {
		return
	}

	j := bytes.Index(request, []byte("\nAS:"))
	if j == -1 {
		return
	}
	response = []byte("HELO:10.3.6.0.false\nAS:2048\nHL:19\n")
	return
}

func (s *Misc) DisguiserResponse(request []byte) (response []byte) {
	i := bytes.Index(request, []byte("http://etherx.jabber.org/streams"))
	if i != -1 {
		response = s.ciscoJabber(request)
		return
	}
	response = s.weblogicRCE(request)
	return
}
