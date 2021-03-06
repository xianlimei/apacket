# APacket

Capture malicious payload

It is built on the shoulders of [Beats](https://github.com/elastic/beats) and [blackhole](https://github.com/dudeintheshell/blackhole). A big thanks.

# Requirements
* System tools
  * conntrack #Netfilter's connection tracking userspace tools

# Features

* Low Interaction Honeypot
* Capture TCP/SYN and backscatter packet only.
* Capture Malicious payload，reference [blackhole](https://github.com/dudeintheshell/blackhole).
* Capture all packets.

# Installation from source

```
go get github.com/Acey9/apacket
cd $GOPATH/src/github.com/Acey9/apacket
make install
apacket -i eth1 -bs -f "not tcp port 22 and not arp" -r 102400 -k 10 -p ./apacket.logs -n apacket.log  -listen 0.0.0.0:54321 -tlslisten 0.0.0.0:54322 -crt ./localhost.crt -key ./ocalhost.key smtp memcached

#install log server
go get github.com/Acey9/sapacket
cd $GOPATH/src/github.com/Acey9/sapacket
make install
sapacket -h
```

# Framework
![apacket](https://github.com/Acey9/apacket/raw/master/doc/images/apacket.png)
