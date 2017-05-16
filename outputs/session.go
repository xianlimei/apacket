package outputs

import (
	"github.com/Acey9/apacket/logp"
	"sync"
	"time"
)

const SessionExpired = 5

type Session struct {
	tab      map[string]time.Time
	cntMutex *sync.RWMutex
}

func NewSesson() *Session {
	s := &Session{tab: make(map[string]time.Time),
		cntMutex: &sync.RWMutex{}}
	go s.clean()
	return s
}

func (s *Session) AddSession(flowid string) {
	s.cntMutex.Lock()
	defer s.cntMutex.Unlock()
	s.tab[flowid] = time.Now()
}

func (s *Session) QuerySession(flowid string) bool {
	s.cntMutex.Lock()
	defer s.cntMutex.Unlock()
	_, ok := s.tab[flowid]
	if ok {
		return true
	}
	return false
}

func (s *Session) DeleteSession(flowid string) {
	s.cntMutex.Lock()
	defer s.cntMutex.Unlock()
	delete(s.tab, flowid)
}

func (s *Session) del() {
	defer func() {
		if err := recover(); err != nil {
			logp.Err("del session error:%v", err)
		}
	}()

	logp.Debug("session", "session map len:%d", len(s.tab))
	for k, v := range s.tab {
		if time.Since(v) > time.Second*SessionExpired {
			logp.Debug("session", "delete session id:%s", k)
			s.DeleteSession(k)
		}
	}
}

func (s *Session) clean() {
	sleep := time.Millisecond * time.Duration(1000)
	for {
		s.del()
		time.Sleep(sleep)
	}
}
