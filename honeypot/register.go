package honeypot

import (
	"github.com/Acey9/apacket/honeypot/core"
	"github.com/Acey9/apacket/honeypot/dns"
	"github.com/Acey9/apacket/honeypot/http"
	"github.com/Acey9/apacket/honeypot/memcached"
	"github.com/Acey9/apacket/honeypot/miscudp"
	"github.com/Acey9/apacket/honeypot/redis"
	"github.com/Acey9/apacket/honeypot/smtp"
	"github.com/Acey9/apacket/honeypot/telnet"
	"github.com/Acey9/apacket/honeypot/vnc"
)

var DisguiserMap []core.Disguiser
var DisguiserMapUDP []core.Disguiser
var serviceMap = map[string]core.Services{}

func init() {
	//TCP
	http := http.NewHTTP()
	DisguiserMap = append(DisguiserMap, http)

	redis := redis.NewRedis()
	DisguiserMap = append(DisguiserMap, redis)
	//TCP END

	//UDP
	dns := dns.NewDNS()
	DisguiserMapUDP = append(DisguiserMapUDP, dns)

	//UDP END
	misc := miscudp.NewMisc()
	DisguiserMapUDP = append(DisguiserMapUDP, misc)

	//services
	smtp := smtp.NewSmtp()
	serviceMap[smtp.Name()] = smtp

	memcached := memcached.NewMemcached()
	serviceMap[memcached.Name()] = memcached

	telnet := telnet.NewTelnet()
	serviceMap[telnet.Name()] = telnet

	vnc := vnc.NewVNC()
	serviceMap[vnc.Name()] = vnc
}
