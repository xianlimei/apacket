package decoder

import (
	"bytes"
	"github.com/tsg/gopacket/layers"
	"net"
	"strconv"
)

type Flow struct {
	Sip, Dip     net.IP
	Sport, Dport uint16
	Protocol     layers.IPProtocol
}

func (f *Flow) FlowID() string {
	id := bytes.Buffer{}
	id.WriteString(f.Protocol.String())
	id.WriteString(f.Sip.String())
	id.WriteString(f.Dip.String())
	id.WriteString(strconv.Itoa(int(f.Sport)))
	id.WriteString(strconv.Itoa(int(f.Dport)))
	return id.String()
}

func (f *Flow) ToOutFlowID() string {
	id := bytes.Buffer{}
	id.WriteString(f.Protocol.String())
	id.WriteString(f.Dip.String())
	id.WriteString(f.Sip.String())
	id.WriteString(strconv.Itoa(int(f.Dport)))
	id.WriteString(strconv.Itoa(int(f.Sport)))
	return id.String()
}
