# APacket

Sniffer syn and backscatter packets

It is built on the shoulders of [Beats](https://github.com/elastic/beats). A big thanks.

# Installation from source

```
go get github.com/Acey9/apacket
cd $GOPATH/src/github.com/Acey9/apacket
make
cp ./apacket $GOPATH/bin/
apacket -h

#install log server
go get github.com/Acey9/sapacket
cd $GOPATH/src/github.com/Acey9/sapacket
make
cp ./sapacket $GOPATH/bin/
sapacket -h
```

# Usage
```
Usage of ./apacket [option]
  -V  Version
  -a string
      Log server auth token
  -b int
      Interface buffer size.(MB) (default 30)
  -bs
      Sniffer syn/backscatter packets only
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
      Keep the number of log files (default 7)
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
      The size of each log file.(KB) (default 10240)
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
