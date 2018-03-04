package core

import (
	"github.com/Acey9/apacket/outputs"
)

type Services interface {
	Start(outputer outputs.Outputer)
	Name() string
}
