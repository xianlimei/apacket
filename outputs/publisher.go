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

type Publisher struct {
	pktQueue    chan *decoder.Packet
	filterQueue chan *decoder.Packet
	outputer    Outputer
	session     *Session
}

func NewPublisher(o Outputer) *Publisher {
	s := NewSesson()
	p := &Publisher{
		pktQueue:    make(chan *decoder.Packet),
		filterQueue: make(chan *decoder.Packet),
		session:     s,
		outputer:    o,
	}
	go p.Start()
	return p
}

func (pub *Publisher) PublishEvent(pkt *decoder.Packet) {
	if config.Cfg.Backscatter {
		pub.filterQueue <- pkt
	} else {
		pub.pktQueue <- pkt
	}
}

func (pub *Publisher) output(pkt *decoder.Packet) {
	if config.Cfg.Backscatter {
		pkt.PayloadSha1 = pkt.CalPayloadSha1()
	}
	b, err := json.Marshal(pkt)
	if err != nil {
		logp.Err("%s", err)
		return
	}
	pub.outputer.Output(b)
}

func (pub *Publisher) getSynID(flowID string, seq uint32) string {
	synID := bytes.Buffer{}
	synID.WriteString(flowID)
	synID.WriteString(strconv.Itoa(int(seq)))
	return synID.String()
}

func (pub *Publisher) filterTCP(pkt *decoder.Packet) *decoder.Packet {

	//ignore not syn and not syn_ack
	//if !pkt.Tcp.SYN && !(pkt.Tcp.SYN && pkt.Tcp.ACK) {
	//	return nil
	//}

	if pkt.Tcp.SYN && pkt.Tcp.ACK { //syn_ack
		//ignore device sended syn_ack
		_, ok := config.Cfg.IfaceAddrs[pkt.Flow.Sip.String()]
		if ok {
			return nil
		}

		//ignore device syn response(syn_ack)
		flowid := pkt.Flow.ToOutFlowID()
		if pub.session.QuerySession(flowid) {
			pub.session.DeleteSession(flowid)
			logp.Debug("filter", "device syn_ack, flow id:%s", pkt.Flow.ToOutFlowID())
			return nil
		}
		return pkt
	} else if pkt.Tcp.SYN { //syn

		//ignore device syn
		_, ok := config.Cfg.IfaceAddrs[pkt.Flow.Sip.String()]
		if ok {
			logp.Debug("filter", "device syn, flow id:%s", pkt.Flow.FlowID())
			pub.session.AddSession(pkt.Flow.FlowID())
			return nil
		} else if !config.Cfg.FirstBloodDisable {
			logp.Debug("filter", "add syn, flow id:%s%v", pkt.Flow.FlowID(), pkt.Tcp.Seq)
			synID := pub.getSynID(pkt.Flow.FlowID(), pkt.Tcp.Seq)
			pub.session.AddSession(synID)
		}
		return pkt
	} else if !config.Cfg.FirstBloodDisable && pkt.Tcp.ACK {
		_, ok := config.Cfg.IfaceAddrs[pkt.Flow.Sip.String()]
		if ok {
			return nil
		}
		synID := pub.getSynID(pkt.Flow.FlowID(), pkt.Tcp.Seq-1)
		if pub.session.QuerySession(synID) && len(pkt.Tcp.Payload) != 0 { //First payload
			return pkt
		}
	}
	return nil
}

func (pub *Publisher) filterUDP(pkt *decoder.Packet) *decoder.Packet {
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
		pub.session.AddSession(flowid)
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
		if pub.session.QuerySession(flowid) {
			pub.session.DeleteSession(flowid)
			logp.Debug("filter", "device udp response, flow id:%s", flowid)
			return nil
		}
	}
	return pkt
}

func (pub *Publisher) filter(pkt *decoder.Packet) *decoder.Packet {
	switch pkt.Flow.Protocol {
	case layers.IPProtocolTCP:
		p := pub.filterTCP(pkt)
		if p == nil {
			return nil
		}
		return pkt
	case layers.IPProtocolUDP:
		p := pub.filterUDP(pkt)
		if p == nil {
			return nil
		}
		return pkt
	}
	return pkt
}

func (pub *Publisher) Start() {
	for {
		select {
		case pkt := <-pub.pktQueue:
			pub.output(pkt)
		case pkt := <-pub.filterQueue:
			p := pub.filter(pkt)
			if p != nil {
				pub.output(pkt)
			}
		}
	}
}
