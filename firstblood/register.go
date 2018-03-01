package firstblood

import (
	"github.com/Acey9/apacket/firstblood/core"
	"github.com/Acey9/apacket/firstblood/http"
	"github.com/Acey9/apacket/firstblood/redis"
)

var DisguiserMap []core.Disguiser

func init() {
	http := http.NewHTTP()
	DisguiserMap = append(DisguiserMap, http)

	redis := redis.NewRedis()
	DisguiserMap = append(DisguiserMap, redis)
}
