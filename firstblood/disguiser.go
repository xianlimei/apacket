package firstblood

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"strconv"
	"strings"
	"time"
)

const (
	TransportTCP = 6
	TransportUDP = 17
	IPv4         = 4
	IPv6         = 6
)

const (
	PtypeHTTP  = "http"
	PtypeOther = "other"
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
	IPv   uint8     `json:"ipv"`
	IP4   *IP4      `json:"ip4,omitempty"`
	IP6   *IP6      `json:"ip6,omitempty"`
	TCP   *TCP      `json:"tcp,omitempty"`
	UDP   *UDP      `json:"udp,omitempty"`
	Http  *HTTPMsg  `json:"http,omitempty"`
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
	Fingerprint(request []byte) (identify bool, err error)
	DisguiserResponse(request []byte) (response []byte)
	Parser(remoteAddr, localAddr string, request []byte) (response *Applayer)
}

var DisguiserMap []Disguiser

func init() {
	http := NewHTTP()
	DisguiserMap = append(DisguiserMap, http)
}

func getIPPort(addr string) (ip string, port uint16, ipv int, err error) {
	var iPort int
	ipPort := strings.Split(addr, ":")
	_len := len(ipPort)
	if _len > 2 {
		ip = strings.Join(ipPort[0:_len-1], ":")
		iPort, err = strconv.Atoi(ipPort[_len-1])
		ipv = IPv6
	} else {
		ip = ipPort[0]
		iPort, err = strconv.Atoi(ipPort[_len-1])
		ipv = IPv4
	}
	if err == nil {
		port = uint16(iPort)
	}
	return
}

func NewApplayer(remoteAddr, localAddr, ptype string, proto uint16, payload []byte) (applayer *Applayer, err error) {

	ip4 := &IP4{}
	ip6 := &IP6{}
	tcp := &TCP{}
	udp := &UDP{}

	applayer = &Applayer{
		Ts:    time.Now(),
		Ptype: ptype,
	}

	sip, sport, ipv, err := getIPPort(remoteAddr)
	if err != nil {
		return
	}

	dip, dport, ipv, err := getIPPort(localAddr)
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

func parseInt(line []byte) (int, error) {
	i, err := strconv.Atoi(string(line))
	return int(i), err
	// TODO: is it an error if 'buf.Len() != 0 {}' ?
}

func trim(buf []byte) []byte {
	return trimLeft(trimRight(buf))
}

func trimLeft(buf []byte) []byte {
	for i, b := range buf {
		if b != ' ' && b != '\t' {
			return buf[i:]
		}
	}
	return nil
}

func trimRight(buf []byte) []byte {
	for i := len(buf) - 1; i > 0; i-- {
		b := buf[i]
		if b != ' ' && b != '\t' {
			return buf[:i+1]
		}
	}
	return nil
}

func toLower(buf, in []byte) []byte {
	if len(in) > len(buf) {
		goto unbufferedToLower
	}

	for i, b := range in {
		if b > 127 {
			goto unbufferedToLower
		}

		if 'A' <= b && b <= 'Z' {
			b = b - 'A' + 'a'
		}
		buf[i] = b
	}
	return buf[:len(in)]

unbufferedToLower:
	return bytes.ToLower(in)
}
