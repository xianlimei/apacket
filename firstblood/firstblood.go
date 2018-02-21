package firstblood

import (
	//"encoding/base64"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/Acey9/apacket/config"
	"github.com/Acey9/apacket/outputs"
	"net"
	"time"
)

const PAYLOAD_MAX_LEN = 524288 //512KB
const (
	TypeHandshake   uint8 = 0x16
	TypeClientHello uint8 = 0x01
	VersionSSLH     uint8 = 0x03
)

type FirstBlood struct {
	ListenAddr    string
	TLSListenAddr string
	outputer      outputs.Outputer
	//sha1Filter *outputs.ShaOneFilter
}

func NewFirstBlood() *FirstBlood {

	var o outputs.Outputer
	var err error

	if config.Cfg.LogServer != "" {
		o, err = outputs.NewSapacketOutputer(config.Cfg.LogServer, config.Cfg.Token)
	} else if len(config.Cfg.NsqdTCPAddress) != 0 {
		o, err = outputs.NewNSQOutputer(config.Cfg.NsqdTCPAddress, config.Cfg.NsqdTopic)
	} else {
		o, err = outputs.NewFileOutputer()
	}

	//o, err = outputs.NewFileOutputer() //TODO DELETE

	if err != nil {
		panic(err)
	}

	//shaone := outputs.NewShaOneFilter()

	fb := &FirstBlood{
		ListenAddr:    config.Cfg.ListenAddr,
		TLSListenAddr: config.Cfg.TLSListenAddr,
		outputer:      o,
		//sha1Filter: shaone,
	}
	return fb
}

func (fb *FirstBlood) Start() {
	if fb.TLSListenAddr != "" {
		go fb.TLSListen("tcp", fb.TLSListenAddr)
	}
	fb.Listen("tcp", fb.ListenAddr)
}

func (fb *FirstBlood) Listen(network, address string) error {
	srv, err := net.Listen(network, address)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for {
		conn, err := srv.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		go fb.initHandler(conn, false)
	}
	return nil
}

func (fb *FirstBlood) TLSListen(network, address string) error {
	cer, err := tls.LoadX509KeyPair(config.Cfg.ServerCrt, config.Cfg.ServerKey)
	if err != nil {
		fmt.Println(err)
		return err
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	srv, err := tls.Listen(network, address, config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for {
		conn, err := srv.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		go fb.initHandler(conn, true)
	}
	return nil
}

func (fb *FirstBlood) tlsRedirect(payload []byte, conn net.Conn) (response []byte) {
	l, err := conn.Write(payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := make([]byte, PAYLOAD_MAX_LEN)
	l, err = conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	response = buf[:l]
	return
}

func (fb *FirstBlood) getTLSProxyConn() (conn net.Conn) {
	conn, err := net.DialTimeout("tcp", config.Cfg.TLSListenAddr, time.Second*3)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (fb *FirstBlood) initHandler(conn net.Conn, isTLSConn bool) {
	var tlsProxyConn net.Conn

	defer func() {
		conn.Close()
		if tlsProxyConn != nil {
			tlsProxyConn.Close()
		}
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//str := "Q05YTgAAAAEAEAAAVgAAAOweAAC8saexZGV2aWNlOjpyby5wcm9kdWN0Lm5hbWU9aG0gbm90ZSAxcztyby5wcm9kdWN0Lm1vZGVsPWhtIG5vdGUgMXM7cm8ucHJvZHVjdC5kZXZpY2U9eDg2OwA="
	//defaultResponse, _ := base64.StdEncoding.DecodeString(str)

	response := []byte("\x00\x00")
	payloadBuf := bytes.Buffer{}

	var stageTls bool
	var firstPalyloadLen int

	for {
		conn.SetDeadline(time.Now().Add(5 * time.Second))
		buf := make([]byte, PAYLOAD_MAX_LEN)
		l, err := conn.Read(buf)
		if err != nil || l < 1 {
			break
		}

		if firstPalyloadLen == 0 {
			firstPalyloadLen = l
		}
		payload := buf[:l]

		//fmt.Println(payload)
		//TODO ssl protocol identify
		if !stageTls && !isTLSConn &&
			l >= 6 && payload[0] == TypeHandshake &&
			payload[1] == VersionSSLH &&
			payload[1] >= 0x00 && payload[1] <= 0x03 && //SSL/3.0 TLS/1.0/1.1/1.2
			payload[5] == TypeClientHello {
			stageTls = true
			tlsProxyConn = fb.getTLSProxyConn()
		}
		if stageTls {
			res := fb.tlsRedirect(payload, tlsProxyConn)
			if len(res) != 0 {
				conn.Write(res)
			}
			continue
		}

		payloadBuf.Write(payload)
		if payloadBuf.Len() > PAYLOAD_MAX_LEN {
			break
		}

		for _, disguiser := range DisguiserMap {
			identify, _ := disguiser.Fingerprint(payload)
			if identify {
				pkt := disguiser.Parser(conn.RemoteAddr().String(), conn.LocalAddr().String(), payload)
				/*
					if fb.sha1Filter.Hit(pkt.Psha1) {
						break
					}
				*/
				out, err := json.Marshal(pkt)
				if err == nil {
					fb.outputer.Output(out)
				}
				response = disguiser.DisguiserResponse(payload)
				conn.Write(response)
				break
			}
		}

	}
	if payloadBuf.Len() > 0 && payloadBuf.Len() != firstPalyloadLen {
		pkt, err := NewApplayer(conn.RemoteAddr().String(), conn.LocalAddr().String(), PtypeOther, TransportTCP, payloadBuf.Bytes())
		if err != nil {
			return
		}
		/*
			if fb.sha1Filter.Hit(pkt.Psha1) {
				return
			}
		*/
		out, err := json.Marshal(pkt)
		if err == nil {
			fb.outputer.Output(out)
		}
	}
}
