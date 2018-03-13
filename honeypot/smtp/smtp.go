package smtp

import (
	"bytes"
	"encoding/json"
	"github.com/Acey9/apacket/honeypot/core"
	"github.com/Acey9/apacket/logp"
	"github.com/Acey9/apacket/outputs"
	"net"
	"net/mail"
	"strings"
	"time"
	"unicode"
)

const PtypeSmtp = "smtp"

var ServerHello = []byte("220 www.github.com ESMTP postfix\r\n")

var ClientHello = []byte("250-mail.github.com\r\n250-PIPELINING\r\n250-SIZE 5242880\r\n250-VRFY\r\n250-ETRN\r\n250-AUTH CRAM-MD5 LOGIN DIGEST-MD5 PLAIN\r\n250-AUTH=CRAM-MD5 LOGIN DIGEST-MD5 PLAIN\r\n250-ENHANCEDSTATUSCODES\r\n250-8BITMIME\r\n250 DSN\r\n")

var authSuccessful = []byte("235 2.7.0 Authentication successful\r\n")

var OK = []byte("250 2.1.0 Ok\r\n")
var sendOK = []byte("250 2.0.0 Ok: queued as 8BCE210C\r\n")

var cmdMap = map[string][]byte{
	"hello": ClientHello,
	"ehlo":  ClientHello,
	"auth":  []byte("334 YWRtaW5AbWFpbC5naXRodWIuY29tCg==\r\n"),
	"data":  []byte("354 End data with <CR><LF>.<CR><LF>\r\n"),
}

type Smtp struct {
	name       string
	listenAddr string
	outputer   outputs.Outputer
}

type SmtpMsg struct {
	Payload     []byte            `json:"-"`
	ContentType string            `json:"ctype,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
}

func NewSmtp() *Smtp {
	m := &Smtp{
		name:       "smtp",
		listenAddr: "0.0.0.0:25",
		//listenAddr: "127.0.0.1:8902",
	}
	return m
}

func (m *Smtp) Start(outputer outputs.Outputer) {
	m.outputer = outputer
	m.Listen()
}

func (m *Smtp) Name() string {
	return m.name
}

func (m *Smtp) Listen() error {
	srv, err := net.Listen("tcp", m.listenAddr)
	if err != nil {
		logp.Err("Smtp.Listen: %v", err)
		return err
	}

	for {
		conn, err := srv.Accept()
		if err != nil {
			logp.Err("Smtp.Listen.Accept %v", err)
			break
		}
		go m.initHandler(conn)
	}
	return nil
}

func (m *Smtp) parserCmd(payload []byte) (cmd string, content []byte) {
	if len(payload) == 0 {
		return
	}

	idxCRLF := bytes.Index(payload, []byte("\r\n"))
	if idxCRLF == -1 {
		return
	}
	idxSpace := bytes.IndexFunc(payload, unicode.IsSpace)
	if idxSpace == -1 {
		cmd = string(payload[0:idxCRLF])
		return
	}

	if idxSpace < 1 {
		return
	}

	cmd = strings.ToLower(string(payload[0:idxSpace]))
	content = payload[idxSpace:]
	return
}

func (m *Smtp) initHandler(conn net.Conn) {

	defer func() {
		conn.Close()
		if err := recover(); err != nil {
			logp.Err("Smtp.initHandler:%v", err)
		}
	}()

	conn.Write(ServerHello)

	var stageAuth, stageSend bool

	raddr := conn.RemoteAddr().String()
	laddr := conn.LocalAddr().String()

	for {
		conn.SetDeadline(time.Now().Add(15 * time.Second))
		buf := make([]byte, 524288)
		l, err := conn.Read(buf)
		if err != nil || l < 1 {
			logp.Debug("smtp", "break...")
			break
		}
		payload := buf[:l]
		cmd, content := m.parserCmd(payload)
		logp.Debug("smtp", "cmd:%s, content:%s", cmd, content)
		resp, ok := cmdMap[cmd]
		if ok {
			logp.Debug("smtp", "response content:%s", string(resp))
			conn.Write(resp)
			if cmd == "auth" {
				stageAuth = true
			} else if cmd == "data" {
				stageSend = true
			}
		} else if stageAuth {
			logp.Debug("smtp", "response auth ok:%s", string(authSuccessful))
			stageAuth = false
			conn.Write(authSuccessful)
			m.sendLog(raddr, laddr, "auth", payload)
		} else if stageSend {
			logp.Debug("smtp", "response send ok:%s", string(sendOK))
			stageSend = false
			conn.Write(sendOK)
			m.sendLog(raddr, laddr, "content", payload)
		} else {
			conn.Write(OK)
		}
	}
}

func (m *Smtp) parser(ctype string, payload []byte) (msg *SmtpMsg) {
	var allowHeaderKey = map[string]bool{
		"From":     true,
		"To":       true,
		"Subject":  true,
		"Reply-to": true,
	}

	msg = &SmtpMsg{
		Payload:     payload,
		ContentType: ctype,
	}
	if ctype != "content" {
		return
	}

	r := bytes.NewReader(payload)
	message, err := mail.ReadMessage(r)
	if err != nil {
		return
	}

	header := message.Header
	msg.Headers = make(map[string]string)
	for key := range allowHeaderKey {
		value := header.Get(key)
		if value != "" {
			msg.Headers[strings.ToLower(key)] = value
		}
	}
	return
}

func (m *Smtp) sendLog(raddr, laddr, ctype string, payload []byte) {
	msg := m.parser(ctype, payload)
	pkt, err := core.NewApplayer(raddr, laddr, PtypeSmtp, core.TransportTCP, payload, false, msg)
	if err != nil {
		logp.Err("Smtp.sendLog:%v", err)
		return
	}

	out, err := json.Marshal(pkt)
	if err == nil {
		m.outputer.Output(out)
	}
}
