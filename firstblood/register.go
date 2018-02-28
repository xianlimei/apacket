package firstblood

import (
	"github.com/Acey9/apacket/firstblood/core"
	"github.com/Acey9/apacket/firstblood/http"
)

var DisguiserMap []core.Disguiser

func init() {
	http := http.NewHTTP()
	DisguiserMap = append(DisguiserMap, http)
}
