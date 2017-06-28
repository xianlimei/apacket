package firstblood

import (
	"fmt"
	"github.com/Acey9/apacket/config"
	"net"
	"time"
)

type FirstBlood struct {
	ListenAddr string
}

func NewFirstBlood() *FirstBlood {
	fb := &FirstBlood{
		ListenAddr: config.Cfg.ListenAddr,
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
		go initHandler(conn)
	}
	return nil
}

func initHandler(conn net.Conn) {
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(10 * time.Second))

	buf := make([]byte, 2048)
	l, err := conn.Read(buf)
	if err != nil || l <= 0 {
		return
	}
	conn.Write(buf[:l])
}
