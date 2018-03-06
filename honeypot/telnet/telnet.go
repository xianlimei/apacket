package telnet

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

func (t *Telnet) setPrompt(prompt string) {
	t.prompt = prompt
}

func (t *Telnet) initHandler(conn net.Conn) {

	defer func() {
		conn.Close()
		if err := recover(); err != nil {
			msg := fmt.Sprintf("%s", err)
			logp.Err(msg)
		}
	}()

	payloadBuf := bytes.Buffer{}

	conn.Write([]byte("\033[?1049h"))
	conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))
	conn.Write([]byte(strings.Replace(strings.Replace("spring\n", "\r\n", "\n", -1), "\n", "\r\n", -1)))

	conn.SetDeadline(time.Now().Add(6 * time.Second))
	conn.Write([]byte(t.motd))
	username, err := t.ReadLine(conn, false)
	if err != nil {
		//logp.Err("username: %s", err)
		return
	}

	conn.SetDeadline(time.Now().Add(6 * time.Second))
	conn.Write([]byte("password: "))
	password, err := t.ReadLine(conn, true)
	if err != nil {
		logp.Err("password: %s", err)
		return
	}

	payloadBuf.WriteString(username)
	payloadBuf.Write([]byte("\n"))
	payloadBuf.WriteString(password)
	payloadBuf.Write([]byte("\n"))

	logp.Debug("telnet", "login success: %s@%s", username, password)
	t.setPrompt(fmt.Sprintf("[%s@bt101 /]# ", username))

	conn.Write([]byte(t.prompt))

	for {
		conn.SetDeadline(time.Now().Add(6 * time.Second))
		line, err := t.ReadLine(conn, false)
		if err != nil {
			//logp.Err("read line: %s", err)
			break
		}

		if line == "exit" || line == "quit" {
			break
		}

		if line != "" {
			logp.Debug("telnet", "input line: %s", line)
		}
		value, ok := cmdMap[line]
		if ok {
			conn.Write([]byte(fmt.Sprintf("%s%s\r\n%s", t.prompt, value, t.prompt)))
		} else {
			conn.Write([]byte(fmt.Sprintf("%s%s: command not found\r\n%s", t.prompt, line, t.prompt)))
		}
		//conn.Write([]byte(t.prompt))
		payloadBuf.WriteString(line)
		payloadBuf.Write([]byte("\n"))
		if payloadBuf.Len() > 524288 {
			break
		}
	}
	raddr := conn.RemoteAddr().String()
	laddr := conn.LocalAddr().String()
	if payloadBuf.Len() != 0 {
		t.sendLog(raddr, laddr, payloadBuf.Bytes())
	}

}

func (t *Telnet) ReadLine(conn net.Conn, masked bool) (string, error) {

	buf := make([]byte, 2048)
	bufPos := 0

	for {
		n, err := conn.Read(buf[bufPos : bufPos+1])
		if err != nil || n != 1 {
			return "", err
		}
		if buf[bufPos] == '\xFF' {
			n, err := conn.Read(buf[bufPos : bufPos+2])
			if err != nil || n != 2 {
				return "", err
			}
			bufPos--
		} else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
			if bufPos > 0 {
				conn.Write([]byte(string(buf[bufPos])))
				bufPos--
			}
			bufPos--
		} else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
			bufPos--
		} else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
			conn.Write([]byte("\r\n"))
			return string(buf[:bufPos]), nil
		} else if buf[bufPos] == 0x03 {
			conn.Write([]byte("^C\r\n"))
			return "", nil
		} else {
			if buf[bufPos] == '\x1B' {
				buf[bufPos] = '^'
				conn.Write([]byte(string(buf[bufPos])))
				bufPos++
				buf[bufPos] = '['
				conn.Write([]byte(string(buf[bufPos])))
			} else if masked {
				conn.Write([]byte("*"))
			} else {
				conn.Write([]byte(string(buf[bufPos])))
			}
		}
		bufPos++
	}
	return string(buf), nil
}
