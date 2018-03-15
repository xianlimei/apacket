package memcached

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Acey9/apacket/honeypot/core"
	"github.com/Acey9/apacket/logp"
	"github.com/Acey9/apacket/outputs"
	"net"
	"strings"
	"time"
)

const (
	PtypeMemcached    = "memcached"
	PtypeMemcachedUDP = "memcachedudp"
)

var commandMap = map[string]string{
	"stats":   commandStats,
	"get":     commandGet,
	"gets":    commandGets,
	"version": commandVersion,
}

type Memcached struct {
	name          string
	listenAddr    string
	udplistenAddr string
	outputer      outputs.Outputer
}

type MemcachedMsg struct {
	Payload     []byte            `json:"-"`
	ContentType string            `json:"ctype,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
}

func NewMemcached() *Memcached {
	m := &Memcached{
		name:          "memcached",
		listenAddr:    "0.0.0.0:11211",
		udplistenAddr: "0.0.0.0:11211",
		//listenAddr:    "127.0.0.1:11211",
		//udplistenAddr: "127.0.0.1:11211",
	}
	return m
}

func (m *Memcached) Start(outputer outputs.Outputer) {
	m.outputer = outputer
	go m.listenUDP()
	m.Listen()
}

func (m *Memcached) Name() string {
	return m.name
}

func (m *Memcached) listenUDP() error {
	udpAddr, err := net.ResolveUDPAddr("udp", m.udplistenAddr)
	if err != nil {
		logp.Err("ListenUDP.ResolveUDPAddr: %v", err)
		return err
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		logp.Err("ListenUDP.ResolveUDPAddr: %v", err)
		return err
	}
	defer conn.Close()
	for {
		m.handlerUDP(conn)
	}
	return nil

}

func (m *Memcached) handlerUDP(conn *net.UDPConn) {
	defer func() {
		if err := recover(); err != nil {
			logp.Err("handlerUDP err:%v", err)
		}
	}()

	for {
		payload := make([]byte, 4096)
		plen, remoteAddr, err := conn.ReadFromUDP(payload)
		if err != nil {
			break
		}
		if plen == 0 {
			break
		}
		payload = payload[:plen]
		m.sendLog(remoteAddr.String(), conn.LocalAddr().String(), PtypeMemcachedUDP, payload)

		response := m.parser(payload)
		if len(response) != 0 {
			logp.Debug("memcached", "memcached.udp.response:% 2x", response)
			_, err = conn.WriteToUDP(response, remoteAddr)
			if err != nil {
				logp.Err("WriteToUDP remoteaddr:%s err:%v", remoteAddr.String(), err)
				break
			}
		}
	}
}

func (m *Memcached) Listen() error {
	srv, err := net.Listen("tcp", m.listenAddr)
	if err != nil {
		logp.Err("Memcached.Listen: %v", err)
		return err
	}

	for {
		conn, err := srv.Accept()
		if err != nil {
			logp.Err("Memcached.Listen.Accept %v", err)
			break
		}
		go m.initHandler(conn)
	}
	return nil
}

func (m *Memcached) initHandler(conn net.Conn) {

	defer func() {
		conn.Close()
		if err := recover(); err != nil {
			logp.Err("Memcached.initHandler:%v", err)
		}
	}()

	raddr := conn.RemoteAddr().String()
	laddr := conn.LocalAddr().String()

	var payload []byte

	for {
		conn.SetDeadline(time.Now().Add(15 * time.Second))
		buf := make([]byte, 524288)
		l, err := conn.Read(buf)
		if err != nil || l < 1 {
			logp.Debug("memcached", "break...")
			break
		}
		payload = buf[:l]
		m.sendLog(raddr, laddr, PtypeMemcached, payload)

		response := m.parser(payload)
		if len(response) != 0 {

			logp.Debug("memcached", "memcached.udp.response:% 2x", response)
			conn.Write(response)
		}
	}
}

func (m *Memcached) parseCommand(payload []byte) (commands []string) {
	cmd := string(payload)
	commands = strings.Split(cmd, " ") //TODO
	return
}

func (m *Memcached) parser(payload []byte) (response []byte) {
	var cmd string
	var command []byte

	r := bytes.NewBuffer(payload)
	command, err := r.ReadBytes('\n')
	if err != nil {
		return
	}

	sz := len(command)
	if sz >= 2 {
		command = command[:sz-2]
	}

	spaceIdx := bytes.IndexByte(command, 0x20)
	if spaceIdx != -1 {
		cmd = string(command[:spaceIdx])
	} else {
		cmd = string(command)
	}

	logp.Debug("memcached", "memcached.parser.command:%s", cmd)
	res, ok := commandMap[cmd]
	if ok {
		if cmd == "get" || cmd == "gets" {
			cmds := m.parseCommand(command)
			if len(cmds) < 2 {
				response = []byte("ERROR\r\n")
			} else {
				response = []byte(fmt.Sprintf(res, cmds[1]))
			}

		} else {
			response = []byte(res)
		}
	} else {
		response = []byte("OK\r\n")
	}
	return
}

func (m *Memcached) sendLog(raddr, laddr, ptype string, payload []byte) {
	var proto uint16
	if ptype == PtypeMemcachedUDP {
		proto = core.TransportUDP
	} else {
		proto = core.TransportTCP
	}
	pkt, err := core.NewApplayer(raddr, laddr, ptype, proto, payload, false, nil)
	if err != nil {
		logp.Err("Memcached.sendLog:%v", err)
		return
	}

	out, err := json.Marshal(pkt)
	if err == nil {
		m.outputer.Output(out)
	}
}
