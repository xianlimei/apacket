package config

import (
	"github.com/Acey9/apacket/logp"
	"github.com/Acey9/apacket/utils"
)

var Cfg Config

type Config struct {
	Iface             *InterfacesConfig
	IfaceAddrs        map[string]bool
	Logging           *logp.Logging
	Backscatter       bool
	FirstBloodDisable bool
	LogServer         string
	NsqdTCPAddress    utils.StringArray
	NsqdTopic         string
	Token             string
	ListenAddr        string
	TLSListenAddr     string
	ServerCrt         string
	ServerKey         string
	Args              []string
}

type InterfacesConfig struct {
	Device       string
	Type         string
	File         string
	WithVlans    bool
	BpfFilter    string
	Snaplen      int
	BufferSizeMb int
	TopSpeed     bool
	Dumpfile     string
	OneAtATime   bool
	Loop         int
}
