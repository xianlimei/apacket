package firstblood

import (
	"encoding/json"
	"fmt"
	"github.com/Acey9/apacket/config"
	"github.com/Acey9/apacket/outputs"
	"net"
	"time"
)

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

	for {
		conn.SetDeadline(time.Now().Add(10 * time.Second))
		buf := make([]byte, 4096)
		l, err := conn.Read(buf)
		if err != nil || l < 1 {
			return
		}

		payload := buf[:l]

		response := []byte("\x00")
		for _, disguiser := range DisguiserMap {
			identify, _ := disguiser.Fingerprint(payload)
			if identify {
				pkt := disguiser.Parser(conn.RemoteAddr().String(), conn.LocalAddr().String(), payload)
				out, err := json.Marshal(pkt)
				if err == nil {
					fb.outputer.Output(out)
				}
				response = disguiser.DisguiserResponse(payload)
				break
			}
		}
		conn.Write(response)
	}
}
