# APacket

Capture network packets

It is built on the shoulders of [Beats](https://github.com/elastic/beats). A big thanks.

# Requirements
* System tools
  * conntrack #Netfilter's connection tracking userspace tools

# Features

* Capture TCP/SYN and backscatter packet only.
* Capture TCP payload，reference [blackhole](https://github.com/dudeintheshell/blackhole).
* Capture all packets.

# Installation from source

```
go get github.com/Acey9/apacket
cd $GOPATH/src/github.com/Acey9/apacket
make install
apacket -h

#install log server
go get github.com/Acey9/sapacket
cd $GOPATH/src/github.com/Acey9/sapacket
make install
sapacket -h
```
