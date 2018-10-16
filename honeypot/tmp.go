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

const shellcode = "02d6202524020fa60101010c02b5202524020fa60101010c27a4ffcaafa4ffe0afa0ffe427a5ffe02806ffff24020fab0101010c2804ffff24020fa10101010c"

var cntMutexTmp *sync.RWMutex

func hex2bytes() []byte {
	src := []byte(shellcode)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, _ := hex.Decode(dst, src)
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
