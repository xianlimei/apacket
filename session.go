package main

import (
	"github.com/Acey9/apacket/logp"
	"time"
)

type Session struct {
	tab map[string]time.Time
}

func NewSesson() *Session {
	s := &Session{tab: make(map[string]time.Time)}
	go s.clean()
	return s
}

func (s *Session) AddSession(flowid string) {
	s.tab[flowid] = time.Now()
}

func (s *Session) QuerySession(flowid string) bool {
	_, ok := s.tab[flowid]
	if ok {
		return true
	}
	return false
}

func (s *Session) DeleteSession(flowid string) {
	delete(s.tab, flowid)
}

func (s *Session) clean() {
	sleep := time.Millisecond * time.Duration(1000)
	for {
		logp.Debug("session", "session map len:%d", len(s.tab))
		for k, v := range s.tab {
			if time.Since(v) > time.Second*5 {
				logp.Debug("session", "delete session id:%s", k)
				s.DeleteSession(k)
			}
		}
		time.Sleep(sleep)
	}
}
