package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Acey9/apacket/config"
	"github.com/Acey9/apacket/decoder"
	"github.com/Acey9/apacket/sniffer"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/tsg/gopacket"
	"github.com/tsg/gopacket/layers"
	"os"
	"runtime"
)

const (
	NAME = "apacket"
)

var cfg config.Config

type MainWorker struct {
	pktQueue chan *decoder.Packet
}

func (this *MainWorker) output(pkt *decoder.Packet) {
	b, err := json.Marshal(pkt)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	os.Stdout.Write([]byte("\n"))
}

func (this *MainWorker) worker() {
	for {
		select {
		case pkt := <-this.pktQueue:
			this.output(pkt)
			break
		}
	}
}

func (this *MainWorker) OnPacket(data []byte, ci *gopacket.CaptureInfo) {
	d := &decoder.Decoder{}
	go func() {
		pkt, _ := d.Process(data, ci)
		if pkt != nil {
			this.pktQueue <- pkt
		}
	}()
}

func createWorker(dl layers.LinkType) (sniffer.Worker, error) {
	w := &MainWorker{
		make(chan *decoder.Packet)}
	go w.worker()
	return w, nil
}

func optParse() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [option]\n", os.Args[0])
		flag.PrintDefaults()
	}

	var ifaceConfig config.InterfacesConfig
	var logging logp.Logging
	var fileRotator logp.FileRotator

	flag.StringVar(&ifaceConfig.Device, "i", "", "interface")
	flag.StringVar(&ifaceConfig.Type, "t", "pcap", "type")
	flag.StringVar(&ifaceConfig.BpfFilter, "f", "", "BpfFilter")

	flag.BoolVar(&ifaceConfig.WithVlans, "wl", true, "with vlans")

	flag.IntVar(&ifaceConfig.Snaplen, "s", 65535, "snap length")
	flag.IntVar(&ifaceConfig.BufferSizeMb, "b", 30, "buffer size mb")

	flag.StringVar(&logging.Level, "l", "info", "logging level")
	flag.StringVar(&fileRotator.Path, "p", "", "log path")
	flag.StringVar(&fileRotator.Name, "n", NAME, "log name")

	flag.Parse()

	cfg.Iface = &ifaceConfig

	logging.Files = &fileRotator
	if logging.Files.Path != "" {
		tofiles := true
		logging.ToFiles = &tofiles
		fmt.Println(*logging.ToFiles)
	}
	cfg.Logging = &logging

	if ifaceConfig.Device == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func init() {
	optParse()
	logp.Init("apacket", cfg.Logging)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	sniff := &sniffer.SnifferSetup{}
	sniff.Init(false, cfg.Iface.BpfFilter, createWorker, cfg.Iface)
	defer sniff.Close()
	sniff.Run()
}
