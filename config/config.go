package config

import (
	"github.com/elastic/beats/libbeat/logp"
)

type Config struct {
	Iface   *InterfacesConfig
	Logging *logp.Logging
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
