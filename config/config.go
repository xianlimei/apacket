package config

import (
	"github.com/Acey9/apacket/logp"
)

var Cfg Config

type Config struct {
	Iface       *InterfacesConfig
	IfaceAddrs  map[string]bool
	Logging     *logp.Logging
	Backscatter bool
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
