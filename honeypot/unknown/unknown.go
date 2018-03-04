package unknown

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/Acey9/apacket/honeypot/core"
	"os/exec"
)

const CmdUnknownResponse = "fb_unknown_response"

type Unknown struct {
}

func NewUnknown() *Unknown {
	u := &Unknown{}
	return u
}

func (s *Unknown) Fingerprint(request []byte, tlsTag bool) (identify bool, ptype string, err error) {
	return
}

func (s *Unknown) Parser(remoteAddr, localAddr string, request []byte, ptype string, tls bool) (response *core.Applayer) {
	return
}

func (s *Unknown) DisguiserResponse(request []byte) (response []byte) {
	var out bytes.Buffer
	str := base64.StdEncoding.EncodeToString(request)
	cmd := exec.Command(CmdUnknownResponse, str)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	response = out.Bytes()
	return
}
