package miscudp

import (
	"github.com/Acey9/apacket/honeypot/core"
	"time"
)

const (
	PtypeMisc = "misc"
)

type Misc struct {
}

func (m *Misc) Fingerprint(request []byte, tlsTag bool) (identify bool, ptype string, err error) {
	ptype = PtypeMisc
	identify = true
	return
}

func (m *Misc) Parser(remoteAddr, localAddr string, request []byte, ptype string, tls bool) (response *core.Applayer) {
	return
}

func (m *Misc) DisguiserResponse(request []byte) (response []byte) {
	netis := &Netis{request: request}
	response = netis.Response()
	sleep := time.Millisecond * time.Duration(100)
	time.Sleep(sleep)
	return
}

func NewMisc() (misc *Misc) {
	return &Misc{}
}
