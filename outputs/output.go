package outputs

import (
	"bytes"
	"encoding/json"
	"github.com/Acey9/apacket/config"
	"github.com/Acey9/apacket/decoder"
	"github.com/Acey9/apacket/logp"
	"github.com/tsg/gopacket/layers"
	"strconv"
)

type Outputer struct {
	pktQueue    chan *decoder.Packet
	filterQueue chan *decoder.Packet
	session     *Session
}

func NewOutputer() *Outputer {
	s := NewSesson()
	o := &Outputer{pktQueue: make(chan *decoder.Packet),
		filterQueue: make(chan *decoder.Packet),
		session:     s}
	go o.Start()
	return o
}

func (out *Outputer) PublishEvent(pkt *decoder.Packet) {
	if config.Cfg.Backscatter {
		out.filterQueue <- pkt
	} else {
		out.pktQueue <- pkt
	}
}

func (out *Outputer) output(pkt *decoder.Packet) {
	b, err := json.Marshal(pkt)
	if err != nil {
		logp.Err("%s", err)
	}
	logp.Info("pkt %s", b)
}

func (out *Outputer) filterTCP(pkt *decoder.Packet) *decoder.Packet {

	//ignore not syn and not syn_ack
	if !pkt.Tcp.SYN && !(pkt.Tcp.SYN && pkt.Tcp.ACK) {
		return nil
	}

	if pkt.Tcp.SYN && pkt.Tcp.ACK { //syn_ack
		//ignore device sended syn_ack
		_, ok := config.Cfg.IfaceAddrs[pkt.Flow.Sip.String()]
		if ok {
			return nil
		}

		//ignore device syn response(syn_ack)
		flowid := pkt.Flow.ToOutFlowID()
		if out.session.QuerySession(flowid) {
			out.session.DeleteSession(flowid)
			logp.Debug("filter", "device syn_ack, flow id:%s", pkt.Flow.ToOutFlowID())
			return nil
		}
	} else if pkt.Tcp.SYN { //syn

		//ignore device syn
		_, ok := config.Cfg.IfaceAddrs[pkt.Flow.Sip.String()]
		if ok {
			logp.Debug("filter", "device syn, flow id:%s", pkt.Flow.FlowID())
			out.session.AddSession(pkt.Flow.FlowID())
			return nil
		}
	}
	return pkt
}

func (out *Outputer) filterUDP(pkt *decoder.Packet) *decoder.Packet {
	if pkt.PktType == decoder.PktTypeDNS && len(pkt.Dns.Questions) > 0 && pkt.Dns.Questions[0].Type == layers.DNSTypePTR {
		return nil
	}
	//ignore device sended udp
	_, ok := config.Cfg.IfaceAddrs[pkt.Flow.Sip.String()]
	if ok {
		flowid := pkt.Flow.FlowID()
		if pkt.PktType == decoder.PktTypeDNS {
			id := bytes.Buffer{}
			id.WriteString(flowid)
			id.WriteString(strconv.Itoa(int(pkt.Dns.ID)))
			flowid = id.String()
		}
		logp.Debug("filter", "device udp, flow id:%s", flowid)
		out.session.AddSession(flowid)
		return nil
	} else {
		//ignore device udp response
		flowid := pkt.Flow.ToOutFlowID()
		if pkt.PktType == decoder.PktTypeDNS {
			id := bytes.Buffer{}
			id.WriteString(flowid)
			id.WriteString(strconv.Itoa(int(pkt.Dns.ID)))
			flowid = id.String()
		}
		if out.session.QuerySession(flowid) {
			out.session.DeleteSession(flowid)
			logp.Debug("filter", "device udp response, flow id:%s", flowid)
			return nil
		}
	}
	return pkt
}

func (out *Outputer) filter(pkt *decoder.Packet) *decoder.Packet {
	switch pkt.Flow.Protocol {
	case layers.IPProtocolTCP:
		p := out.filterTCP(pkt)
		if p == nil {
			return nil
		}
		return pkt
	case layers.IPProtocolUDP:
		p := out.filterUDP(pkt)
		if p == nil {
			return nil
		}
		return pkt
	}
	return pkt
}

func (out *Outputer) Start() {
	for {
		select {
		case pkt := <-out.pktQueue:
			out.output(pkt)
		case pkt := <-out.filterQueue:
			p := out.filter(pkt)
			if p != nil {
				out.output(pkt)
			}
		}
	}
}
