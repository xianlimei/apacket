package main

import (
	"encoding/json"
	"github.com/Acey9/apacket/decoder"
	"github.com/Acey9/apacket/logp"
	"github.com/tsg/gopacket/layers"
	"net"
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
	return o
}

func (out *Outputer) PublishEvent(pkt *decoder.Packet) {
	if cfg.Backscatter {
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

func (out *Outputer) getSipDipProtocol(pkt *decoder.Packet) (s, d, proto string) {
	var sip, dip net.IP
	var p layers.IPProtocol

	if pkt.IPV == 4 {
		sip = pkt.Ip4.SrcIP
		dip = pkt.Ip4.DstIP
		p = pkt.Ip4.Protocol
	} else if pkt.IPV == 6 {
		sip = pkt.Ip6.SrcIP
		dip = pkt.Ip6.DstIP
		p = pkt.Ip6.NextHeader
	} else {
		return "", "", proto
	}
	return sip.String(), dip.String(), p.String()
}

func (out *Outputer) filterTCP(pkt *decoder.Packet) *decoder.Packet {

	//ignore not syn and not syn_ack
	if !pkt.Tcp.SYN && !(pkt.Tcp.SYN && pkt.Tcp.ACK) {
		return nil
	}

	if pkt.Tcp.SYN && pkt.Tcp.ACK { //syn_ack
		//ignore device sended syn_ack
		_, ok := cfg.IfaceAddrs[pkt.Flow.Sip.String()]
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
		_, ok := cfg.IfaceAddrs[pkt.Flow.Sip.String()]
		if ok {
			logp.Debug("filter", "device syn, flow id:%s", pkt.Flow.FlowID())
			out.session.AddSession(pkt.Flow.FlowID())
			return nil
		}
	}
	return pkt
}

func (out *Outputer) filterUDP(pkt *decoder.Packet) *decoder.Packet {
	//ignore device sended udp
	_, ok := cfg.IfaceAddrs[pkt.Flow.Sip.String()]
	if ok {
		logp.Debug("filter", "device udp, flow id:%s", pkt.Flow.FlowID())
		out.session.AddSession(pkt.Flow.FlowID())
		return nil
	} else {
		//ignore device sended udp response
		flowid := pkt.Flow.ToOutFlowID()
		if out.session.QuerySession(flowid) {
			out.session.DeleteSession(flowid)
			logp.Debug("filter", "device udp response, flow id:%s", pkt.Flow.ToOutFlowID())
			return nil
		}
	}
	return pkt
}

func (out *Outputer) filter(pkt *decoder.Packet) *decoder.Packet {
	switch pkt.PktType {
	case decoder.PktTypeTCP:
		p := out.filterTCP(pkt)
		if p == nil {
			return nil
		}
		return pkt
	case decoder.PktTypeUDP:
		p := out.filterUDP(pkt)
		if p == nil {
			return nil
		}
		return pkt
	case decoder.PktTypeICMPv4:
		//TODO
	case decoder.PktTypeDNS:
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
