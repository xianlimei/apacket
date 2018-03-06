package telnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Acey9/apacket/honeypot/core"
	"github.com/Acey9/apacket/logp"
	"github.com/Acey9/apacket/outputs"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"net"
	"time"
)

const PtypeTelnet = "telnet"

var (
	motd   = "buildroot login: "
	prompt = `[root@bt101 /]# `
)

type Telnet struct {
	name       string
	listenAddr string
	outputer   outputs.Outputer
	prompt     string
	motd       string
}

func NewTelnet() *Telnet {
	t := &Telnet{
		name:       "telnet",
		listenAddr: "0.0.0.0:23",
		//listenAddr: "127.0.0.1:8902",
		prompt: prompt,
		motd:   motd,
	}
	return t
}

func (t *Telnet) Name() string {
	return t.name
}

func (t *Telnet) Start(outputer outputs.Outputer) {
	t.outputer = outputer
	t.Listen()
}

func (t *Telnet) Listen() error {
	srv, err := net.Listen("tcp", t.listenAddr)
	if err != nil {
		logp.Err("Telnet.Listen: %v", err)
	}

	for {
		conn, err := srv.Accept()
		if err != nil {
			logp.Err("Telnet.Listen.Accept %v", err)
			break
		}
		go t.initHandler(conn)
	}
	return nil

}

func (t *Telnet) initHandler(conn net.Conn) {
	defer func() {
		conn.Close()
		if err := recover(); err != nil {
			logp.Err("Telnet.initHandler:%v", err)
		}
	}()

	payloadBuf := bytes.Buffer{}

	conn.SetDeadline(time.Now().Add(50 * time.Second))
	term := terminal.NewTerminal(conn, prompt)
	user, err := term.ReadPassword(motd)
	if err != nil {
		return
	}
	pwd, err := term.ReadPassword("password:")
	if err != nil {
		return
	}
	payloadBuf.Write([]byte(user + "\n"))
	payloadBuf.Write([]byte(pwd + "\n"))
	pro := fmt.Sprintf("[%s@bt101 /]# ", string(user))
	term.SetPrompt(pro)
	for {
		conn.SetDeadline(time.Now().Add(5 * time.Second))
		line, err := term.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			logp.Err("Telnet.initHandler.ReadLine:%v", err)
			break
		}
		//TODO
		payloadBuf.Write([]byte(line + "\n"))

		if payloadBuf.Len() > 524288 {
			break
		}

		if line == "exit" || line == "quit" {
			break
		}

		value, ok := cmdMap[line]
		if ok {
			term.Write([]byte(value + "\n"))
		} else {
			term.Write([]byte(fmt.Sprintf("%s: command not found\n", line)))
		}
	}

	raddr := conn.RemoteAddr().String()
	laddr := conn.LocalAddr().String()
	if payloadBuf.Len() != 0 {
		t.sendLog(raddr, laddr, payloadBuf.Bytes())
	}
}

func (t *Telnet) sendLog(raddr, laddr string, payload []byte) {
	pkt, err := core.NewApplayer(raddr, laddr, PtypeTelnet, core.TransportTCP, payload, false, nil)
	if err != nil {
		logp.Err("Smtp.sendLog:%v", err)
		return
	}

	out, err := json.Marshal(pkt)
	if err == nil {
		t.outputer.Output(out)
	}
}
