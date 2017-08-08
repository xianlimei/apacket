package outputs

import (
	"bytes"
	"compress/zlib"
	"crypto/tls"
	"github.com/Acey9/apacket/logp"
	"github.com/Acey9/sapacket/packet"
	"net"
	"time"
)

type SapacketOutputer struct {
	Addr     string
	Token    string
	Conn     net.Conn
	msgQueue chan []byte
}

func NewSapacketOutputer(serverAddr, token string) (*SapacketOutputer, error) {
	so := &SapacketOutputer{
		Addr:     serverAddr,
		Token:    token,
		msgQueue: make(chan []byte, 1024),
	}
	err := so.Init()
	if err != nil {
		return nil, err
	}
	go so.Start()
	return so, nil
}

func (this *SapacketOutputer) Init() error {
	conn, err := this.ConnectServer(this.Addr)
	if err != nil {
		logp.Err("connect server error: %v", err)
		return err
	}
	this.Conn = conn
	return nil
}

func (this *SapacketOutputer) Close() {
	logp.Info("sapacket connect close.")
	this.Conn.Close()
}

func (this *SapacketOutputer) ReConnect() error {
	logp.Warn("reconnect server.")
	conn, err := this.ConnectServer(this.Addr)
	if err != nil {
		logp.Err("reconnect server error: %v", err)
		return err
	}
	this.Conn = conn
	return nil
}

func (this *SapacketOutputer) ConnectServer(addr string) (conn net.Conn, err error) {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err = tls.Dial("tcp", addr, conf)
	if err != nil {
		return nil, err
	}

	conn.SetDeadline(time.Now().Add(30 * time.Second))
	login, err := packet.Pack(packet.LOGIN, []byte(this.Token))
	if err != nil {
		return nil, err
	}
	err = packet.WritePacket(conn, login)
	if err != nil {
		//logp.Err("login faield. %v", err)
		return nil, err
	}

	conn.SetDeadline(time.Now().Add(30 * time.Second))
	pkt, err := packet.ReadPacket(conn)
	if err != nil {
		//logp.Err("login faield. %v", err)
		return nil, err
	}

	if pkt.Type != packet.LOGINSUCC {
		//logp.Err("login faield. %v", err)
		return nil, err
	}

	return conn, nil
}

func (this *SapacketOutputer) Output(msg []byte) {
	logp.Info("pkt %s", msg)

	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	w.Write(msg)
	w.Close()
	pkt, err := packet.Pack(packet.PACKET, buf.Bytes())
	if err != nil {
		logp.Err("output error:%v", err)
		return
	}
	this.msgQueue <- pkt
}

func (this *SapacketOutputer) Send(msg []byte) {
	this.Conn.SetDeadline(time.Now().Add(10 * time.Second))
	err := packet.WritePacket(this.Conn, msg)
	if err != nil {
		this.Conn.Close()
		err = this.ReConnect()
		if err != nil {
			logp.Err("send to server error: %v", err)
			return
		}
		logp.Debug("reconnect", "succ")
		err = packet.WritePacket(this.Conn, msg)
		if err != nil {
			logp.Err("resend to server error: %v", err)
		}
		return
	}
}

func (this *SapacketOutputer) Start() {
	counter := 0

	defer func() {
		if err := recover(); err != nil {
			logp.Err("sapacket error:: %v", err)
		}
		this.Close()
	}()

	for {
		select {
		case msg := <-this.msgQueue:
			counter++
			this.Send(msg)
			if counter%1024 == 0 {
				logp.Debug("logserver", "logserver packet number: %d", counter)
			}
		}
	}
}
