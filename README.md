# apacket
Capture syn scan and backscatter packets

# Usage
```
Usage of apacket [option]
  -V    version
  -b int
    	interface buffer size mb (default 30)
  -bs
    	capture syn scan/backscatter packets only
  -d string
    	Enable certain debug selectors
  -df string
    	dump to file
  -e	Log to stderr and disable syslog/file output
  -f string
    	BpfFilter
  -i string
    	listen on interface
  -k int
    	number of keep files (default 7)
  -l string
    	logging level (default "info")
  -lp int
    	loop
  -n string
    	log name (default "apacket.log")
  -p string
    	log path
  -r uint
        rotate every KB (default 10240)
  -rf string
    	read packets from file
  -s int
    	snap length (default 65535)
  -t string
    	type (default "pcap")
  -v	Log at INFO level
  -wl
    	with vlans
```
