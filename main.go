package main

import (
	"flag"
	"fmt"
	"github.com/Acey9/apacket/config"
	"github.com/Acey9/apacket/decoder"
	"github.com/Acey9/apacket/firstblood"
	"github.com/Acey9/apacket/logp"
	"github.com/Acey9/apacket/outputs"
	"github.com/Acey9/apacket/sniffer"
	"github.com/Acey9/apacket/utils"
	"github.com/tsg/gopacket"
	"github.com/tsg/gopacket/layers"
	"os"
	"runtime"
)

const version = "apacket 3.0.5"

type MainWorker struct {
	publisher *outputs.Publisher
	decoder   *decoder.Decoder
}

func (this *MainWorker) OnPacket(data []byte, ci *gopacket.CaptureInfo) {
	//go func() {
	pkt, _ := this.decoder.Process(data, ci)
	if pkt != nil {
		this.publisher.PublishEvent(pkt)
	}
	//}()
}

func NewWorker(dl layers.LinkType) (sniffer.Worker, error) {
	var o outputs.Outputer
	var err error

	if config.Cfg.LogServer != "" {
		o, err = outputs.NewSapacketOutputer(config.Cfg.LogServer, config.Cfg.Token)
	} else if len(config.Cfg.NsqdTCPAddress) != 0 {
		o, err = outputs.NewNSQOutputer(config.Cfg.NsqdTCPAddress, config.Cfg.NsqdTopic)
	} else {
		o, err = outputs.NewFileOutputer()
	}
	if err != nil {
		panic(err)
	}

	p := outputs.NewPublisher(o)

	d := decoder.NewDecoder()

	w := &MainWorker{
		publisher: p,
		decoder:   d,
	}
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
	destNsqdTCPAddrs := utils.StringArray{}

	flag.StringVar(&ifaceConfig.Device, "i", "", "Listen on interface")
	flag.StringVar(&ifaceConfig.Type, "t", "pcap", "Sniffer type.Possible case like pcap,af_packet,pfring,pf_ring")
	flag.StringVar(&ifaceConfig.BpfFilter, "f", "", "BPF filter")
	flag.StringVar(&ifaceConfig.File, "rf", "", "Read packets from file")
	flag.StringVar(&ifaceConfig.Dumpfile, "df", "", "Dump to file")
	flag.IntVar(&ifaceConfig.Loop, "lp", 0, "Loop")

	flag.BoolVar(&ifaceConfig.WithVlans, "wl", false, "With vlans")

	flag.IntVar(&ifaceConfig.Snaplen, "s", 65535, "Snap length")
	flag.IntVar(&ifaceConfig.BufferSizeMb, "b", 30, "Interface buffer size.(MB)")

	flag.StringVar(&logging.Level, "l", "info", "Logging level")
	flag.StringVar(&fileRotator.Path, "p", "", "Log path")
	flag.StringVar(&fileRotator.Name, "n", "apacket.log", "Log filename")
	flag.Uint64Var(&rotateEveryKB, "r", 10240, "The size of each log file.(KB)")
	flag.IntVar(&keepFiles, "k", 7, "Keep the number of log files")

	flag.BoolVar(&config.Cfg.Backscatter, "bs", false, "Sniffer syn/backscatter packets only")

	flag.StringVar(&config.Cfg.LogServer, "ls", "", "Log server address.The log will send to this server")

	flag.Var(&destNsqdTCPAddrs, "nsqd-tcp-address", "destination nsqd TCP address (may be given multiple times)")
	flag.StringVar(&config.Cfg.NsqdTopic, "nsqd-topic", "t.apacket", "NSQ topic to publish to")

	flag.StringVar(&config.Cfg.Token, "a", "", "Log server auth token")

	flag.BoolVar(&config.Cfg.FirstBloodDisable, "dfb", false, "Disable firstblood")
	flag.StringVar(&config.Cfg.ListenAddr, "listen", "", "Listen address")

	printVersion := flag.Bool("V", false, "Version")

	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	config.Cfg.Iface = &ifaceConfig
	config.Cfg.NsqdTCPAddress = destNsqdTCPAddrs

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

func sayHi() {
	fmt.Println("apacket version: ", version)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	sayHi()
	if !config.Cfg.FirstBloodDisable && config.Cfg.ListenAddr != "" {
		fb := firstblood.NewFirstBlood()
		go fb.Start()
	}
	sniff := &sniffer.SnifferSetup{}
	sniff.Init(false, config.Cfg.Iface.BpfFilter, NewWorker, config.Cfg.Iface)
	defer sniff.Close()
	err := sniff.Run()
	if err != nil {
		logp.Err("main %v", err)
	}
}
