package vnc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Acey9/apacket/honeypot/core"
	"github.com/Acey9/apacket/logp"
	"github.com/Acey9/apacket/outputs"
	"net"
	"time"
)

const (
	PtypeVNC       = "vnc"
	version        = "RFB 003.008\n"
	versionAncient = "RFB 003.003\n"
	challenge      = "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
)

type VNC struct {
	name       string
	listenAddr string
	outputer   outputs.Outputer
}

func NewVNC() *VNC {
	v := &VNC{
		name:       "vnc",
		listenAddr: "0.0.0.0:5900",
	}
	return v
}

func (v *VNC) Name() string {
	return v.name
}

func (v *VNC) Start(outputer outputs.Outputer) {
	v.outputer = outputer
	v.Listen()
}

func (v *VNC) Listen() error {
	srv, err := net.Listen("tcp", v.listenAddr)
	if err != nil {
		logp.Err("VNC.Listen: %v", err)
	}

	for {
		conn, err := srv.Accept()
		if err != nil {
			logp.Err("VNC.Listen.Accept %v", err)
			break
		}
		go v.initHandler(conn)
	}
	return nil

}

func (v *VNC) initHandler(conn net.Conn) {
	defer func() {
		conn.Close()
		if err := recover(); err != nil {
			msg := fmt.Sprintf("%s", err)
			logp.Err(msg)
		}
	}()

	payloadBuf := v.handler(conn)
	raddr := conn.RemoteAddr().String()
	laddr := conn.LocalAddr().String()
	if payloadBuf.Len() != 0 {
		v.sendLog(raddr, laddr, payloadBuf.Bytes())
	}

}
func (v *VNC) handler(conn net.Conn) (payloadBuf bytes.Buffer) {
	if _, err := conn.Write([]byte(version)); nil != err {
		logp.Err("%v disconnected before client version: %v", conn.RemoteAddr(), err)
		return
	}

	payloadBuf = bytes.Buffer{}

	SetReadTimeout(conn, 10)
	n, ver, err := v.ReadNBytesFromPipe(conn, len(version))
	if err != nil || n != len(version) {
		if n > 0 {
			payloadBuf.Write(ver)
		}
		return
	}

	payloadBuf.Write(ver)
	strVer := string(ver)

	switch strVer {
	case version:
		if _, err := conn.Write([]byte("\x01\x02")); err != nil {
			return
		}

		_, buf, err := v.ReadNBytesFromPipe(conn, 1)
		if err != nil {
			return
		}
		payloadBuf.Write(buf)
		if buf[0] != 0x02 {
			return
		}
	case versionAncient:
		if _, err := conn.Write([]byte{0, 0, 0, 2}); err != nil {
			return
		}
	default:
		return
	}

	if _, err := conn.Write([]byte(challenge)); nil != err {
		return
	}

	SetReadTimeout(conn, 10)
	n, res, err := v.ReadNBytesFromPipe(conn, 16)
	if err != nil || n != 16 {
		if n > 0 {
			payloadBuf.Write(res)
		}
		return
	}
	payloadBuf.Write(res)

	if payloadBuf.Len() > 524288 {
		return
	}

	conn.Write(append(
		[]byte{
			0, 0, 0, 0,
			0, 0, 0, 4,
		},
		[]byte("succ")...,
	))

	SetReadTimeout(conn, 10)
	n, res, err = v.ReadNBytesFromPipe(conn, 40960)
	if n != 0 {
		payloadBuf.Write(res)
	}

	return
}

func (v *VNC) ReadNBytesFromPipe(conn net.Conn, n int) (_len int, res []byte, err error) {
	var m int
	buf := make([]byte, n)
	bufPos := 0
	for {
		m, err = conn.Read(buf[bufPos : bufPos+1])
		if err != nil || m != 1 {
			break
		}
		_len += m
		bufPos++
		if bufPos >= n {
			break
		}

	}
	res = buf[:_len]
	return
}
func (v *VNC) sendLog(raddr, laddr string, payload []byte) {
	pkt, err := core.NewApplayer(raddr, laddr, PtypeVNC, core.TransportTCP, payload, false, nil)
	if err != nil {
		logp.Err("Smtp.sendLog:%v", err)
		return
	}

	out, err := json.Marshal(pkt)
	if err == nil {
		v.outputer.Output(out)
	}
}

func SetReadTimeout(c net.Conn, s int) {
	c.SetReadDeadline(time.Now().Add(time.Duration(s) * time.Second))
}
