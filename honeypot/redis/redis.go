package redis

import (
	"bytes"
	"encoding/base64"
	"github.com/Acey9/apacket/honeypot/core"
	"github.com/Acey9/apacket/logp"
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

const (
	CmdRedisResponse = "fb_redis_response"
	PtypeRedis       = "redis"
)

type Redis struct {
}

func NewRedis() *Redis {
	redis := &Redis{}
	return redis
}

func (s *Redis) Fingerprint(request []byte, tlsTag bool) (identify bool, ptype string, err error) {

	ptype = PtypeRedis

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
		return
	}
	return
}

func (s *Redis) Parser(remoteAddr, localAddr string, request []byte, ptype string, tls bool) (response *core.Applayer) {
	//msg := &RedisMsg{Payload: request}
	//response.Appl = msg
	response, err := core.NewApplayer(remoteAddr, localAddr, ptype, core.TransportTCP, request, tls, nil)
	if err != nil {
		return
	}
	return
}

func (s *Redis) DisguiserResponse(request []byte) (response []byte) {
	var out bytes.Buffer
	str := base64.StdEncoding.EncodeToString(request)
	cmd := exec.Command(CmdRedisResponse, str)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		logp.Err("Redis.DisguiserResponse.cmd:%v", err)
		return
	}
	response = out.Bytes()
	return
}
