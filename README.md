# APacket

Sniffer syn and backscatter packets

# Usage
```
Usage of ./apacket [option]
  -V  Version
  -a string
      Auth token
  -b int
      Interface buffer size mb (default 30)
  -bs
      Sniffer syn scan/backscatter packets only
  -d string
      Enable certain debug selectors
  -df string
      Dump to file
  -e  Log to stderr and disable syslog/file output
  -f string
      BPF filter
  -i string
      Listen on interface
  -k int
      Number of keep files (default 7)
  -l string
      Logging level (default "info")
  -lp int
      Loop
  -ls string
      Log server address.The log will send to this server
  -n string
      Log filename (default "apacket.log")
  -p string
      Log path
  -r uint
      Rotate every KB (default 10240)
  -rf string
      Read packets from file
  -s int
      Snap length (default 65535)
  -t string
      Sniffer type.Possible case like pcap,af_packet,pfring,pf_ring (default "pcap")
  -v  Log at INFO level
  -wl
      With vlans
```