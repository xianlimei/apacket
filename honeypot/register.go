package honeypot

import (
	"github.com/Acey9/apacket/honeypot/core"
	"github.com/Acey9/apacket/honeypot/http"
	"github.com/Acey9/apacket/honeypot/redis"
	"github.com/Acey9/apacket/honeypot/smtp"
)

var DisguiserMap []core.Disguiser
var serviceMap = map[string]core.Services{}

func init() {
	http := http.NewHTTP()
	DisguiserMap = append(DisguiserMap, http)

	redis := redis.NewRedis()
	DisguiserMap = append(DisguiserMap, redis)

	//services
	smtp := smtp.NewSmtp()
	serviceMap[smtp.Name()] = smtp
}
