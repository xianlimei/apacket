package redis

import (
	"bytes"
	"fmt"
	"github.com/Acey9/apacket/firstblood/core"
	"os/exec"
	"strconv"
	"strings"
)

var dataTypeMap = map[uint8]bool{
	0x2a: true, //*
	0x2b: true, //+
	0x24: true, //$
	0x3a: true, //:
}

const CmdRedisResponse = "fb_redis_response"

type Redis struct {
}

func NewRedis() *Redis {
	redis := &Redis{}
	return redis
}

func (s *Redis) Fingerprint(request []byte, tlsTag bool) (identify bool, ptype string, err error) {
	//fmt.Printf("request:% 2x\n", request)
	//fmt.Println(string(request))

	dataType := request[0]
	if dataType != 0x2a { //*
		return
	}

	i := bytes.Index(request, []byte("\x0d\x0a\x24")) //\r\n$
	if i < 2 {
		return
	}

	argNumStr := string(request[1:i])
	argNum, err := strconv.ParseUint(argNumStr, 10, 64)
	if err != nil {
		return
	}

	if argNum < 1 {
		return
	}

	request = request[i+3:]
	i = bytes.Index(request, []byte("\x0d\x0a")) //\r\n
	if i < 1 {
		return
	}

	arg1LenStr := string(request[:i])
	arg1Len, err := strconv.ParseUint(arg1LenStr, 10, 64)
	if err != nil {
		return
	}
	cmd := string(request[i+2 : uint64(i+2)+arg1Len])
	if request[uint64(i+2)+arg1Len] != 0x0d && request[uint64(i+3)+arg1Len] != 0x0a {
		return
	}
	cmd = strings.ToUpper(cmd)
	_, ok := allowCMDMap[cmd]
	if ok {
		identify = true
		ptype = core.PtypeRedis
		return
	}
	return
}

func (s *Redis) Parser(remoteAddr, localAddr string, request []byte, ptype string, tls bool) (response *core.Applayer) {
	response, err := core.NewApplayer(remoteAddr, localAddr, ptype, core.TransportTCP, nil, tls)
	if err != nil {
		return
	}
	msg := &RedisMsg{Payload: request}
	cPayload := response.Compress(request)
	msg.Payload = cPayload.Bytes()
	response.Appl = msg
	response.Psha1 = response.Sha1HexDigest(string(request))
	response.Plen = uint(len(request))
	return
}

func (s *Redis) DisguiserResponse(request []byte) (response []byte) {
	var out bytes.Buffer
	cmd := exec.Command(CmdRedisResponse, string(request))
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	response = out.Bytes()
	return
}
