package miscudp

import (
	"bytes"
)

type SSDP struct {
	request []byte
}

func (ssdp *SSDP) Response() (resp []byte) {
	keyIdx := bytes.Index(ssdp.request, []byte("MAN:\"ssdp:discover\""))
	if keyIdx == -1 {
		return
	}
	buf := bytes.Buffer{}
	buf.Write([]byte("HTTP/1.1 200 OK\r\n"))
	buf.Write([]byte("Server: Custom/1.0 UPnP/1.0 Proc/Ver\r\n"))
	buf.Write([]byte("EXT:\r\n"))
	buf.Write([]byte("Location: http://192.168.1.1:5431/dyndev/uuid:38e59522-dd48-48dd-2295-e538e522480000\r\n"))
	buf.Write([]byte("Cache-Control:max-age=1800\r\n"))
	buf.Write([]byte("ST:ssdp:all\r\n"))
	buf.Write([]byte("USN:uuid:38e59522-dd48-48dd-2295-e538e522480000::upnp:rootdevice\r\n"))
	buf.Write([]byte("\r\n"))
	resp = buf.Bytes()
	return
}
