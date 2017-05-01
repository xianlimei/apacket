package main

import (
	"encoding/json"
	"github.com/Acey9/apacket/decoder"
	"github.com/elastic/beats/libbeat/logp"
)

type Outputer struct {
	pktQueue    chan *decoder.Packet
	filterQueue chan *decoder.Packet
}

func NewOutputer() *Outputer {
	o := &Outputer{pktQueue: make(chan *decoder.Packet),
		filterQueue: make(chan *decoder.Packet)}
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

func (out *Outputer) filterTCP(pkt *decoder.Packet) *decoder.Packet {
	if !pkt.Tcp.SYN && !(pkt.Tcp.SYN && pkt.Tcp.ACK) {
		return nil
	}
	//TODO filter our syn and syn_ack
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
		//TODO
	case decoder.PktTypeICMPv4:
		//TODO
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
