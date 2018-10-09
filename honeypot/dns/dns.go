package dns

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/Acey9/apacket/honeypot/core"
	"github.com/Acey9/apacket/logp"
	"github.com/tsg/gopacket"
	"github.com/tsg/gopacket/layers"
	"math/rand"
	"net"
	"strings"
	"time"
)

const (
	PtypeDNS = "dns"
)

var SLD = []string{
	"www.Itiscertain",
	"ns.Itisdecidedlyso",
	"cname.Withoutadoubt",
	"Yesdefinitely",
	"Youmayrelyonit",
	"AsIseeityes",
	"Mostlikely",
	"www.Outlookgood",
	"ns.Yes",
	"mx.Signspointtoyes",
	"mx.Replyhazytryagain",
	"Askagainlater",
	"Betternottellyounow",
	"Cannotpredictnow",
	"Concentrateandaskagain",
	"Dontcountonit",
	"Myreplyisno",
	"Mysourcessayno",
	"Outlooknotsogood",
	"Verydoubtful",
}

var SUFFIX = []string{
	"com",
	"me",
	"org",
	"net",
	"tk",
	"win",
}

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

func (dns *DNS) DisguiserResponse(request []byte, remoteAddr string) (response []byte) {
	packet := gopacket.NewPacket(request, layers.LayerTypeDNS, gopacket.NoCopy)
	dnsLayer := packet.Layer(layers.LayerTypeDNS)
	if dnsLayer == nil {
		return
	}
	dnsPkt, ok := dnsLayer.(*layers.DNS)
	if !ok {
		return
	}
	if dnsPkt.QR { //dns response
		return
	}
	msg := NewMsg(dnsPkt)
	response = msg.Bytes()
	return
}

type Msg struct {
	query *layers.DNS
	buf   bytes.Buffer
}

func (m *Msg) uint16Field(n uint16) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, n)
	m.buf.Write(b)
}

func (m *Msg) uint32Field(n uint32) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, n)
	m.buf.Write(b)
}

func (m *Msg) stringField(s string) {
	m.buf.WriteString(s)
}

func (m *Msg) encodeName(name string) []byte {
	buf := bytes.Buffer{}
	names := strings.Split(name, ".")
	for _, rn := range names {
		b := []byte{uint8(len(rn))}
		buf.Write(b)
		buf.WriteString(rn)
	}
	buf.WriteByte(0x00)
	return buf.Bytes()
}

func (m *Msg) queryRecordField(questions []layers.DNSQuestion) {
	for _, rr := range questions {
		name := m.encodeName(string(rr.Name))
		m.buf.Write(name)
		m.uint16Field(uint16(rr.Type))
		m.uint16Field(uint16(rr.Class))
	}
}

func (m *Msg) answerRecordField(answers []layers.DNSResourceRecord) {
	for _, rr := range answers {
		m.buf.Write(m.encodeName(string(rr.Name)))
		m.uint16Field(uint16(rr.Type))
		m.uint16Field(uint16(rr.Class))
		m.uint32Field(uint32(rr.TTL))
		if rr.Type == layers.DNSTypeMX {
			m.uint16Field(rr.DataLength + 2)
			m.uint16Field(0x32)
		} else {
			m.uint16Field(rr.DataLength)
		}
		m.buf.Write(rr.Data)
	}
}

func (m *Msg) appendRR(rname []byte, rdata string, rtype layers.DNSType, class layers.DNSClass) {
	rRecord := layers.DNSResourceRecord{
		Name:  rname,
		Type:  rtype,
		Class: class,
		TTL:   900,
	}
	switch rtype {
	case layers.DNSTypeA:
		ip := net.ParseIP(rdata)
		rRecord.Data = ip[12:16]
	case layers.DNSTypeAAAA:
		rRecord.Data = net.ParseIP(rdata)
	case layers.DNSTypeCNAME:
		rRecord.Data = m.encodeName(rdata)
		rRecord.DataLength = uint16(len(rRecord.Data))
		m.query.Answers = append(m.query.Answers, rRecord)

		cnameIP := m.randomIPv4()
		m.appendRR([]byte(rdata), cnameIP, layers.DNSTypeA, layers.DNSClassIN)
		return
	default:
		rRecord.Data = m.encodeName(rdata)
	}

	rRecord.DataLength = uint16(len(rRecord.Data))
	m.query.Answers = append(m.query.Answers, rRecord)
}

func (m *Msg) Bytes() (payload []byte) {
	//DNS Header
	m.uint16Field(m.query.ID)

	//Flag
	flag := uint16(0x8180)
	m.uint16Field(flag)

	//QDCOUNT
	m.uint16Field(uint16(len(m.query.Questions)))

	//ANCOUNT
	m.uint16Field(uint16(len(m.query.Answers)))

	//NSCOUNT
	m.uint16Field(uint16(len(m.query.Authorities)))

	//ARCOUNT
	m.uint16Field(uint16(len(m.query.Additionals)))

	//Questions
	m.queryRecordField(m.query.Questions)

	//Answer
	m.answerRecordField(m.query.Answers)

	payload = m.buf.Bytes()
	return
}

func (m *Msg) randomDomain() (fqdn string) {
	rand.Seed(time.Now().UnixNano())
	sld := SLD[rand.Intn(len(SLD))]
	suffix := SUFFIX[rand.Intn(len(SUFFIX))]
	fqdn = fmt.Sprintf("%s.%s", sld, suffix)
	return
}

func (m *Msg) randomIPv4() (ip string) {
	rand.Seed(time.Now().UnixNano())
	ip = fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return
}

func (m *Msg) randomRdata(rtype layers.DNSType, name []byte) (rdata string) {
	switch rtype {
	case layers.DNSTypeA:
		rdata = m.randomIPv4()
	case layers.DNSTypeAAAA:
		rdata = "::1"
	case layers.DNSTypeCNAME, layers.DNSTypeNS, layers.DNSTypeMX:
		rdata = m.randomDomain()
	case layers.DNSTypeTXT:
		rdata = "v=spf1 mx ~all"
	default:
		rdata = "com"
	}
	return
}

func NewMsg(query *layers.DNS) (msg *Msg) {
	msg = &Msg{
		query: query,
		buf:   bytes.Buffer{},
	}

	for _, rr := range msg.query.Questions {
		if rr.Type == 255 {
			//add A
			rdata := msg.randomRdata(layers.DNSTypeA, rr.Name)
			msg.appendRR(rr.Name, rdata, layers.DNSTypeA, rr.Class)
			//add CNAME
			rdata = msg.randomRdata(layers.DNSTypeCNAME, rr.Name)
			msg.appendRR(rr.Name, rdata, layers.DNSTypeCNAME, rr.Class)
			//add TXT
			rdata = "v=spf1 mx ~all"
			msg.appendRR(rr.Name, rdata, layers.DNSTypeTXT, rr.Class)
			//add MX
			rdata = msg.randomRdata(layers.DNSTypeMX, rr.Name)
			msg.appendRR(rr.Name, rdata, layers.DNSTypeMX, rr.Class)
		} else {
			rdata := msg.randomRdata(rr.Type, rr.Name)
			msg.appendRR(rr.Name, rdata, rr.Type, rr.Class)
		}
	}
	return
}
