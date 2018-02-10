package outputs

import (
	"bufio"
	"github.com/Acey9/apacket/logp"
	"os"
	"strings"
	"sync"
)

const (
	SHA1LIST = "APACKET_SHA1_LIST"
)

type ShaOneFilter struct {
	tab      map[string]bool
	sha1list string
	cntMutex *sync.RWMutex
}

func NewShaOneFilter() (shaone *ShaOneFilter) {
	sha1list := os.Getenv(SHA1LIST)
	shaone = &ShaOneFilter{
		tab:      make(map[string]bool),
		sha1list: sha1list,
		cntMutex: &sync.RWMutex{},
	}

	if sha1list == "" {
		return
	}

	targetFile, err := os.Open(sha1list)
	if err != nil {
		return
	}
	defer targetFile.Close()

	fielScanner := bufio.NewScanner(targetFile)
	for fielScanner.Scan() {
		sha1 := strings.Trim(fielScanner.Text(), " \n")
		if sha1 != "" {
			shaone.Add(sha1)
		}
	}

	return
}

func (shaone *ShaOneFilter) Hit(sha1 string) bool {
	shaone.cntMutex.Lock()
	defer shaone.cntMutex.Unlock()
	_, ok := shaone.tab[sha1]
	logp.Debug("sha1Filter", "filter:%s,hit:%v", sha1, ok)
	return ok
}

func (shaone *ShaOneFilter) Add(sha1 string) {
	shaone.cntMutex.Lock()
	defer shaone.cntMutex.Unlock()
	shaone.tab[sha1] = true
}
