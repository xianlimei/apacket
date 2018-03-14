package outputs

import (
	"github.com/Acey9/apacket/logp"
)

type FileOutputer struct {
}

func (this *FileOutputer) Output(msg []byte) {
	if len(msg) < 5 {
		return
	}
	logp.Info("pkt %s", msg)
}

func NewFileOutputer() (*FileOutputer, error) {
	fo := &FileOutputer{}
	return fo, nil
}
