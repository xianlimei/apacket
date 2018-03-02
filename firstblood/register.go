package firstblood

import (
	"github.com/Acey9/apacket/firstblood/core"
	"github.com/Acey9/apacket/firstblood/http"
	"github.com/Acey9/apacket/firstblood/redis"
	"github.com/Acey9/apacket/firstblood/smtp"
)

var DisguiserMap []core.Disguiser
var serviceMap = map[string]core.Services{}

const (
	nameServiceMail = "mail"
)

func init() {
	http := http.NewHTTP()
	DisguiserMap = append(DisguiserMap, http)

	redis := redis.NewRedis()
	DisguiserMap = append(DisguiserMap, redis)

	//services
	smtp := smtp.NewSmtp()
	serviceMap[smtp.Name()] = smtp
}
