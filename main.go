package main

import (
	"flag"
	"fmt"
	"github.com/Acey9/apacket/config"
	"github.com/Acey9/apacket/decoder"
	"github.com/Acey9/apacket/logp"
	"github.com/Acey9/apacket/outputs"
	"github.com/Acey9/apacket/sniffer"
	"github.com/Acey9/apacket/utils"
	"github.com/tsg/gopacket"
	"github.com/tsg/gopacket/layers"
	"os"
	"runtime"
)

const version = "1.0.2"

type MainWorker struct {
	outputer *outputs.Outputer
	decoder  *decoder.Decoder
}

func (this *MainWorker) OnPacket(data []byte, ci *gopacket.CaptureInfo) {
	//go func() {
	pkt, _ := this.decoder.Process(data, ci)
	if pkt != nil {
		this.outputer.PublishEvent(pkt)
	}
	//}()
}

func NewWorker(dl layers.LinkType) (sniffer.Worker, error) {
	o := outputs.NewOutputer()

	d := decoder.NewDecoder()

	w := &MainWorker{outputer: o,
		decoder: d}
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
	var rotateEveryKB uint64
	var keepFiles int

	flag.StringVar(&ifaceConfig.Device, "i", "", "listen on interface")
	flag.StringVar(&ifaceConfig.Type, "t", "pcap", "type")
	flag.StringVar(&ifaceConfig.BpfFilter, "f", "", "BpfFilter")
	flag.StringVar(&ifaceConfig.File, "rf", "", "read packets from file")
	flag.StringVar(&ifaceConfig.Dumpfile, "df", "", "dump to file")
	flag.IntVar(&ifaceConfig.Loop, "lp", 0, "loop")

	flag.BoolVar(&ifaceConfig.WithVlans, "wl", false, "with vlans")

	flag.IntVar(&ifaceConfig.Snaplen, "s", 65535, "snap length")
	flag.IntVar(&ifaceConfig.BufferSizeMb, "b", 30, "interface buffer size mb")

	flag.StringVar(&logging.Level, "l", "info", "logging level")
	flag.StringVar(&fileRotator.Path, "p", "", "log path")
	flag.StringVar(&fileRotator.Name, "n", "apacket.log", "log name")
	flag.Uint64Var(&rotateEveryKB, "r", 10240, "rotate every KB")
	flag.IntVar(&keepFiles, "k", 7, "number of keep files")

	flag.BoolVar(&config.Cfg.Backscatter, "bs", false, "capture syn scan/backscatter packets only")

	printVersion := flag.Bool("V", false, "version")

	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	config.Cfg.Iface = &ifaceConfig

	logging.Files = &fileRotator
	if logging.Files.Path != "" {
		tofiles := true
		logging.ToFiles = &tofiles

		rotateKB := rotateEveryKB * 1024
		logging.Files.RotateEveryBytes = &rotateKB
		logging.Files.KeepFiles = &keepFiles
	}
	config.Cfg.Logging = &logging

	if ifaceConfig.Device == "" && ifaceConfig.File == "" {
		flag.Usage()
		os.Exit(1)
	}

	if ifaceConfig.Device != "" {
		ifaceAddrs, err := utils.InterfaceAddrsByName(config.Cfg.Iface.Device)
		if err != nil {
			flag.Usage()
			fmt.Println("get interface addrs error.")
			os.Exit(1)
		}

		config.Cfg.IfaceAddrs = make(map[string]bool)
		for _, addr := range ifaceAddrs {
			config.Cfg.IfaceAddrs[addr] = true
		}
	}
}

func init() {
	optParse()
	logp.Init("apacket", config.Cfg.Logging)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	sniff := &sniffer.SnifferSetup{}
	sniff.Init(false, config.Cfg.Iface.BpfFilter, NewWorker, config.Cfg.Iface)
	defer sniff.Close()
	sniff.Run()
}
