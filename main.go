package main

import (
	"flag"
	"fmt"
	"github.com/Acey9/apacket/config"
	"github.com/Acey9/apacket/decoder"
	"github.com/Acey9/apacket/logp"
	"github.com/Acey9/apacket/sniffer"
	"github.com/Acey9/apacket/utils"
	"github.com/tsg/gopacket"
	"github.com/tsg/gopacket/layers"
	"os"
	"runtime"
)

var cfg config.Config

var outputer *Outputer = NewOutputer()

type MainWorker struct {
}

func (this *MainWorker) OnPacket(data []byte, ci *gopacket.CaptureInfo) {
	d := &decoder.Decoder{}
	//go func() {
	pkt, _ := d.Process(data, ci)
	if pkt != nil {
		outputer.PublishEvent(pkt)
	}
	//}()
}

func createWorker(dl layers.LinkType) (sniffer.Worker, error) {
	w := &MainWorker{}
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

	flag.BoolVar(&ifaceConfig.WithVlans, "wl", false, "with vlans")

	flag.IntVar(&ifaceConfig.Snaplen, "s", 65535, "snap length")
	flag.IntVar(&ifaceConfig.BufferSizeMb, "b", 30, "buffer size mb")

	flag.StringVar(&logging.Level, "l", "info", "logging level")
	flag.StringVar(&fileRotator.Path, "p", "", "log path")
	flag.StringVar(&fileRotator.Name, "n", "apacket.log", "log name")

	flag.BoolVar(&cfg.Backscatter, "bs", false, "catch backscatter only")

	flag.Parse()

	cfg.Iface = &ifaceConfig

	logging.Files = &fileRotator
	if logging.Files.Path != "" {
		tofiles := true
		logging.ToFiles = &tofiles
	}
	cfg.Logging = &logging

	if ifaceConfig.Device == "" {
		flag.Usage()
		os.Exit(1)
	}

	ifaceAddrs, err := utils.InterfaceAddrsByName(cfg.Iface.Device)
	if err != nil {
		flag.Usage()
		fmt.Println("get interface addrs error.")
		os.Exit(1)
	}

	cfg.IfaceAddrs = make(map[string]bool)
	for _, addr := range ifaceAddrs {
		cfg.IfaceAddrs[addr] = true
	}
}

func init() {
	optParse()
	logp.Init("apacket", cfg.Logging)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go outputer.Start()

	sniff := &sniffer.SnifferSetup{}
	sniff.Init(false, cfg.Iface.BpfFilter, createWorker, cfg.Iface)
	defer sniff.Close()
	sniff.Run()
}
