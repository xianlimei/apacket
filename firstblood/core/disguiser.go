package core

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"time"
)

type IP4 struct {
	Sip      string `json:"sip"`
	Dip      string `json:"dip"`
	Protocol uint16 `json:"proto"`
}

type IP6 struct {
	Sip      string `json:"sip"`
	Dip      string `json:"dip"`
	Protocol uint16 `json:"proto"`
}

type TCP struct {
	Sport   uint16 `json:"sport"`
	Dport   uint16 `json:"dport"`
	Payload []byte `json:"payload,omitempty"`
}

type UDP struct {
	Sport   uint16 `json:"sport"`
	Dport   uint16 `json:"dport"`
	Payload []byte `json:"payload,omitempty"`
}

type Applayer struct {
	Ts    time.Time `json:"ts"`
	Ptype string    `json:"ptype"`
	Psha1 string    `json:"psha1,omitempty"`
	Plen  uint      `json:"plen,omitempty"`
	TLS   bool      `json:"tls,omitempty"`
	IPv   uint8     `json:"ipv"`
	IP4   *IP4      `json:"ip4,omitempty"`
	IP6   *IP6      `json:"ip6,omitempty"`
	TCP   *TCP      `json:"tcp,omitempty"`
	UDP   *UDP      `json:"udp,omitempty"`
	//Http  *HTTPMsg  `json:"http,omitempty"`
	Appl interface{} `json:"appl,omitempty"`
}

func (app *Applayer) Compress(source []byte) bytes.Buffer {
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	w.Write(source)
	w.Close()
	return buf
}

func (app *Applayer) Sha1HexDigest(str string) string {
	h := sha1.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

type Disguiser interface {
	Fingerprint(request []byte, tlsTag bool) (identify bool, ptype string, err error)
	DisguiserResponse(request []byte) (response []byte)
	Parser(remoteAddr, localAddr string, request []byte, ptype string, tls bool) (response *Applayer)
}

func NewApplayer(remoteAddr, localAddr, ptype string, proto uint16, payload []byte, tls bool) (applayer *Applayer, err error) {

	ip4 := &IP4{}
	ip6 := &IP6{}
	tcp := &TCP{}
	udp := &UDP{}

	applayer = &Applayer{
		Ts:    time.Now(),
		Ptype: ptype,
		TLS:   tls,
	}

	sip, sport, ipv, err := GetIPPort(remoteAddr)
	if err != nil {
		return
	}

	dip, dport, ipv, err := GetIPPort(localAddr)
	if err != nil {
		return
	}

	applayer.IPv = uint8(ipv)

	if ipv == IPv4 {
		ip4.Sip = sip
		ip4.Dip = dip
		ip4.Protocol = proto
		applayer.IP4 = ip4
	} else {
		ip6.Sip = sip
		ip6.Dip = dip
		ip6.Protocol = proto
		applayer.IP6 = ip6
	}

	var cPayload []byte
	if len(payload) != 0 {
		cP := applayer.Compress(payload)
		cPayload = cP.Bytes()
		applayer.Psha1 = applayer.Sha1HexDigest(string(payload))
		applayer.Plen = uint(len(payload))

	}

	if proto == TransportTCP {
		tcp.Sport = sport
		tcp.Dport = dport
		if len(cPayload) != 0 {
			tcp.Payload = cPayload
		}
		applayer.TCP = tcp
	} else if proto == TransportUDP {
		udp.Sport = sport
		udp.Dport = dport
		if len(cPayload) != 0 {
			udp.Payload = cPayload
		}
		applayer.UDP = udp
	}

	return
}
