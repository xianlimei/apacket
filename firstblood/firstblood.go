package firstblood

import (
	//"encoding/base64"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Acey9/apacket/config"
	"github.com/Acey9/apacket/outputs"
	"net"
	"time"
)

const PAYLOAD_MAX_LEN = 1048576 //1MB

type FirstBlood struct {
	ListenAddr string
	outputer   outputs.Outputer
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

	fb := &FirstBlood{
		ListenAddr: config.Cfg.ListenAddr,
		outputer:   o,
	}
	return fb
}

func (fb *FirstBlood) Start() {
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
		go fb.initHandler(conn)
	}
	return nil
}

func (fb *FirstBlood) initHandler(conn net.Conn) {

	defer func() {
		conn.Close()
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//str := "Q05YTgAAAAEAEAAAVgAAAOweAAC8saexZGV2aWNlOjpyby5wcm9kdWN0Lm5hbWU9aG0gbm90ZSAxcztyby5wcm9kdWN0Lm1vZGVsPWhtIG5vdGUgMXM7cm8ucHJvZHVjdC5kZXZpY2U9eDg2OwA="
	//defaultResponse, _ := base64.StdEncoding.DecodeString(str)

	response := []byte("\x00\x00")
	payloadBuf := bytes.Buffer{}

	for {
		conn.SetDeadline(time.Now().Add(5 * time.Second))
		buf := make([]byte, PAYLOAD_MAX_LEN)
		l, err := conn.Read(buf)
		if err != nil || l < 1 {
			break
		}

		payload := buf[:l]
		payloadBuf.Write(payload)
		if payloadBuf.Len() > PAYLOAD_MAX_LEN {
			break
		}

		for _, disguiser := range DisguiserMap {
			identify, _ := disguiser.Fingerprint(payload)
			if identify {
				pkt := disguiser.Parser(conn.RemoteAddr().String(), conn.LocalAddr().String(), payload)
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
	pkt, err := NewApplayer(conn.RemoteAddr().String(), conn.LocalAddr().String(), PtypeOther, TransportTCP, payloadBuf.Bytes())
	if err != nil {
		return
	}
	out, err := json.Marshal(pkt)
	if err == nil {
		fb.outputer.Output(out)
	}
}
