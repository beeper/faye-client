package client

import (
	"sync"

	"github.com/beeper/faye-client/pkg/message"
)

type Subscriptions struct {
	lock *sync.RWMutex

	subscriptions map[string][]func(*message.Message)
}

func NewSubscriptions() *Subscriptions {
	return &Subscriptions{
		lock:          &sync.RWMutex{},
		subscriptions: make(map[string][]func(*message.Message)),
	}
}

func (s *Subscriptions) Add(channel string, callback func(*message.Message)) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.subscriptions[channel] = append(s.subscriptions[channel], callback)
}
