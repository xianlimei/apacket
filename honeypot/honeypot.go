package honeypot

import (
	//"encoding/base64"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/Acey9/apacket/config"
	"github.com/Acey9/apacket/honeypot/core"
	"github.com/Acey9/apacket/honeypot/unknown"
	"github.com/Acey9/apacket/logp"
	"github.com/Acey9/apacket/outputs"
	"net"
	"time"
)

const (
	PAYLOAD_MAX_LEN = 524288 //512KB
	PtypeOther      = "other"
)

const (
	TypeHandshake   uint8 = 0x16
	TypeClientHello uint8 = 0x01
	VersionSSLH     uint8 = 0x03
)

type Honeypot struct {
	ListenAddr    string
	TLSListenAddr string
	outputer      outputs.Outputer
	session       *Session
	//sha1Filter *outputs.ShaOneFilter
}

func NewHoneypot() *Honeypot {

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

	hp := &Honeypot{
		ListenAddr:    config.Cfg.ListenAddr,
		TLSListenAddr: config.Cfg.TLSListenAddr,
		outputer:      o,
		session:       NewSesson(),
		//sha1Filter: shaone,
	}
	return hp
}

func (hp *Honeypot) Start() {

	hp.ServicesStart()

	if hp.TLSListenAddr != "" {
		go hp.TLSListen("tcp", hp.TLSListenAddr)
	}
	hp.Listen("tcp", hp.ListenAddr)
}

func (hp *Honeypot) ServicesStart() {
	for _, module := range config.Cfg.Args {
		svr, ok := serviceMap[module]
		if !ok {
			logp.Warn("service %s not found.", module)
			continue
		}
		go svr.Start(hp.outputer)
		logp.Info("%s service start.", module)
	}
}

func (hp *Honeypot) Listen(network, address string) error {
	srv, err := net.Listen(network, address)
	if err != nil {
		logp.Err("Listen: %v", err)
		return err
	}

	for {
		conn, err := srv.Accept()
		if err != nil {
			logp.Err("Listen.Accept %v", err)
			break
		}
		go hp.initHandler(conn, false)
	}
	return nil
}

func (hp *Honeypot) TLSListen(network, address string) error {
	cer, err := tls.LoadX509KeyPair(config.Cfg.ServerCrt, config.Cfg.ServerKey)
	if err != nil {
		logp.Err("TLSListen.tls.config:%v", err)
		return err
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	srv, err := tls.Listen(network, address, config)
	if err != nil {
		logp.Err("TLSListen:%v", err)
		return err
	}

	for {
		conn, err := srv.Accept()
		if err != nil {
			logp.Err("TLSListen.Accept:%v", err)
			break
		}
		go hp.initHandler(conn, true)
	}
	return nil
}

func (hp *Honeypot) tlsRedirect(payload []byte, conn net.Conn) (response []byte) {
	l, err := conn.Write(payload)
	if err != nil {
		return
	}
	buf := make([]byte, PAYLOAD_MAX_LEN)
	l, err = conn.Read(buf)
	if err != nil {
		return
	}
	response = buf[:l]
	return
}

func (hp *Honeypot) getTLSProxyConn() (conn net.Conn, tlsProxyLocalAddr string) {
	conn, err := net.DialTimeout("tcp", config.Cfg.TLSListenAddr, time.Second*3)
	if err != nil {
		logp.Err("getTLSProxyConn:%v", err)
		return
	}
	tlsProxyLocalAddr = conn.LocalAddr().String()
	return
}

func (hp *Honeypot) initHandler(conn net.Conn, isTLSConn bool) {
	var tlsProxyConn net.Conn
	var tlsProxyLocalAddr, remoteAddr, localAddr string
	var identify bool

	defer func() {
		conn.Close()
		if tlsProxyConn != nil {
			hp.session.DeleteSession(tlsProxyLocalAddr)
			tlsProxyConn.Close()
		}
		if err := recover(); err != nil {
			logp.Err("initHandler remote:%s local:%s info:%v", remoteAddr, localAddr, err)
		}
	}()

	//str := "Q05YTgAAAAEAEAAAVgAAAOweAAC8saexZGV2aWNlOjpyby5wcm9kdWN0Lm5hbWU9aG0gbm90ZSAxcztyby5wcm9kdWN0Lm1vZGVsPWhtIG5vdGUgMXM7cm8ucHJvZHVjdC5kZXZpY2U9eDg2OwA="
	//defaultResponse, _ := base64.StdEncoding.DecodeString(str)

	response := []byte("\x00\x00")
	payloadBuf := bytes.Buffer{}

	var stageTls, tlsTag bool
	var firstPalyloadLen int

	remoteAddr = conn.RemoteAddr().String()
	localAddr = conn.LocalAddr().String()

	unk := unknown.NewUnknown()

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

		//TODO ssl protocol identify
		if !stageTls && !isTLSConn &&
			l >= 6 && payload[0] == TypeHandshake &&
			payload[1] == VersionSSLH &&
			payload[1] >= 0x00 && payload[1] <= 0x03 && //SSL/3.0 TLS/1.0/1.1/1.2
			payload[5] == TypeClientHello {
			stageTls = true
			tlsProxyConn, tlsProxyLocalAddr = hp.getTLSProxyConn()
			nf := &Netflow{remoteAddr, localAddr}
			hp.session.AddSession(tlsProxyLocalAddr, nf)
		}
		if stageTls {
			res := hp.tlsRedirect(payload, tlsProxyConn)
			if len(res) != 0 {
				conn.Write(res)
			}
			continue
		}

		payloadBuf.Write(payload)
		if payloadBuf.Len() > PAYLOAD_MAX_LEN {
			break
		}

		if isTLSConn {
			tlsTag = true
		}

		netflow, ok := hp.session.QuerySession(conn.RemoteAddr().String())
		if ok {
			remoteAddr = netflow.RemoteAddr
			localAddr = netflow.LocalAddr
			tlsTag = true
		}

		for _, disguiser := range DisguiserMap {
			identify, response = hp.identifyProto(disguiser, payload, remoteAddr, localAddr, tlsTag)
			if identify {
				if len(response) != 0 {
					conn.Write(response)
				}
				break
			}
		}
		if !identify {
			response = unk.DisguiserResponse(payload)
			//response = []byte("\x00\x00")
			if len(response) != 0 {
				conn.Write(response)
			}
		}

	}
	if payloadBuf.Len() > 0 && payloadBuf.Len() != firstPalyloadLen {
		pkt, err := core.NewApplayer(remoteAddr, localAddr, PtypeOther, core.TransportTCP, payloadBuf.Bytes(), tlsTag, nil)
		if err != nil {
			return
		}
		/*
			if hp.sha1Filter.Hit(pkt.Psha1) {
				return
			}
		*/
		out, err := json.Marshal(pkt)
		if err == nil {
			hp.outputer.Output(out)
		}
	}
}

func (hp *Honeypot) identifyProto(disguiser core.Disguiser, payload []byte, remoteAddr, localAddr string, tlsTag bool) (identify bool, response []byte) {
	defer func() {
		if err := recover(); err != nil {
			logp.Err("identifyProto remote:%s local:%s info:%v", remoteAddr, localAddr, err)
		}
	}()

	identify, ptype, _ := disguiser.Fingerprint(payload, tlsTag)
	logp.Debug("honeypot", "disguiser.Fingerprint identify:%v, ptype:%v", identify, ptype)
	if identify {
		pkt := disguiser.Parser(remoteAddr, localAddr, payload, ptype, tlsTag)
		/*
			if hp.sha1Filter.Hit(pkt.Psha1) {
				break
			}
		*/
		out, err := json.Marshal(pkt)
		if err == nil {
			hp.outputer.Output(out)
		}
		response = disguiser.DisguiserResponse(payload)

	}
	return
}