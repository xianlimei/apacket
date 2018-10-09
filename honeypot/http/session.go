package http

import (
	"github.com/Acey9/apacket/logp"
	"sync"
	"time"
)

const SessionExpired = 3600 //second
const interval = 100000     //Millisecond

type FmtString struct {
	Ts  time.Time
	Str string
}

type Session struct {
	tab      map[string]*FmtString
	cntMutex *sync.RWMutex
}

func NewSesson() *Session {
	s := &Session{tab: make(map[string]*FmtString),
		cntMutex: &sync.RWMutex{}}
	go s.clean()
	return s
}

func NewFmtString(str string) (fs *FmtString) {
	fs = &FmtString{
		Ts:  time.Now(),
		Str: str,
	}
	return fs
}

func (s *Session) AddSession(rip, fmtstring string) {
	s.cntMutex.Lock()
	defer s.cntMutex.Unlock()
	fs := NewFmtString(fmtstring)
	logp.Debug("http", "add new session:%s - %s", rip, fmtstring)
	s.tab[rip] = fs
}

func (s *Session) QuerySession(rip string) *FmtString {
	s.cntMutex.RLock()
	defer s.cntMutex.RUnlock()
	res, ok := s.tab[rip]
	if ok {
		return res
	}
	return nil
}

func (s *Session) DeleteSession(rip string) {
	s.cntMutex.Lock()
	defer s.cntMutex.Unlock()
	delete(s.tab, rip)
}

func (s *Session) del() {
	s.cntMutex.Lock()
	defer s.cntMutex.Unlock()

	logp.Debug("session", "session map len:%d", len(s.tab))
	for k, v := range s.tab { //TODO fatal error: concurrent map iteration and map write
		if time.Since(v.Ts) > time.Second*SessionExpired {
			logp.Debug("session", "clean session id:%s", k)
			delete(s.tab, k)
		}
	}
}

func (s *Session) clean() {
	sleep := time.Millisecond * time.Duration(interval)
	for {
		s.del()
		time.Sleep(sleep)
	}
}
