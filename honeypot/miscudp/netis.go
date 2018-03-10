//Netis Routers Backdoor
//https://blog.trendmicro.com/trendlabs-security-intelligence/netis-routers-leave-wide-open-backdoor/
//https://raw.githubusercontent.com/h00die/MSF-Testing-Scripts/master/netis_backdoor.py

package miscudp

import (
	"bytes"
)

type Netis struct {
	request []byte
}

func (netis *Netis) Login() bool {
	//login message:	AAAAAAAAnetcore\x00
	a8idx := bytes.Index(netis.request, []byte("AAAAAAAA"))
	if a8idx != 0 {
		return false
	}

	lastByte := netis.request[len(netis.request)-1]
	if lastByte != 0x00 {
		return false
	}
	return true
}

func (netis *Netis) Command() bool {
	//cmd:	AA\x00\x00AAAA%s\x00
	prefixIdx := bytes.Index(netis.request, []byte("AA\x00\x00AAAA"))
	if prefixIdx != 0 {
		return false
	}

	lastByte := netis.request[len(netis.request)-1]
	if lastByte != 0x00 {
		return false
	}
	return true
}

func (netis *Netis) Response() (resp []byte) {
	if netis.Login() {
		resp = []byte("AA\x00\x05ABAA\x00\x00\x00\x00Login successed!\r\n")
		return
	} else if netis.Command() {
		resp = []byte("kcimu")
	}
	return
}
