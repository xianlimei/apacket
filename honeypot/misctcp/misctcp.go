package misctcp

import (
	"bytes"
	"encoding/binary"
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
	i := bytes.Index(request, []byte("http://etherx.jabber.org/streams"))
	if i != -1 {
		return
	}
	reqLen := len(request)
	i = bytes.Index(request, []byte("to="))
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

func (s *Misc) weblogicCVE20182893(request []byte) (response []byte) {
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

func (s *Misc) linksysBackdoor(request []byte) (response []byte) {
	//https://thehackernews.com/2014/01/hacking-wireless-dsl-routers-via.html
	i := bytes.Index(request, []byte("\x4D\x4D\x63\x53")) //MMcS
	if i != 0 {
		return
	}
	if len(request) < 12 {
		return
	}
	cmd := binary.LittleEndian.Uint32(request[4:8])
	if cmd < 1 || cmd > 13 {
		return
	}

	buf := bytes.Buffer{}
	buf.WriteString("\x4D\x4D\x63\x53\x00\x00\x00\x00")

	var msg string
	msg = "ok"

	if cmd == 1 {
		msg = "user=root\x00|\x01password=123456"
	}

	len_buf := new(bytes.Buffer)
	binary.Write(len_buf, binary.LittleEndian, uint32(len(msg)))
	buf.Write(len_buf.Bytes())
	buf.WriteString(msg)
	response = buf.Bytes()

	return
}

func (s *Misc) DisguiserResponse(request []byte, remoteAddr string) (response []byte) {
	response = s.weblogicCVE20182893(request)
	if len(response) != 0 {
		return
	}
	response = s.ciscoJabber(request)
	if len(response) != 0 {
		return
	}
	response = s.linksysBackdoor(request)
	return
}
