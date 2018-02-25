package firstblood

import (
	//"fmt"
	"sync"
)

type Netflow struct {
	RemoteAddr string
	LocalAddr  string
}

type Session struct {
	tab      map[string]*Netflow
	cntMutex *sync.RWMutex
}

func NewSesson() *Session {
	s := &Session{tab: make(map[string]*Netflow),
		cntMutex: &sync.RWMutex{}}
	return s
}

func (s *Session) AddSession(sid string, netflow *Netflow) {
	//fmt.Println("AddSession:", sid, netflow.RemoteAddr, netflow.LocalAddr)
	s.cntMutex.Lock()
	defer s.cntMutex.Unlock()
	s.tab[sid] = netflow
}

func (s *Session) QuerySession(sid string) (netflow *Netflow, ok bool) {
	s.cntMutex.RLock()
	defer s.cntMutex.RUnlock()
	netflow, ok = s.tab[sid]
	return
}

func (s *Session) DeleteSession(sid string) {
	//fmt.Println("DeleteSession:", sid)
	s.cntMutex.Lock()
	defer s.cntMutex.Unlock()
	delete(s.tab, sid)
}
