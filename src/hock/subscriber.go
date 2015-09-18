package main

import (
	"github.com/pborman/uuid"
)

const SUBSCRIBER_BUFFER_SIZE = 100

type Subscriber struct {
	id    string
	ch    chan string
	drops int
}

func NewSubscriber() *Subscriber {
	return &Subscriber{
		uuid.New(),
		make(chan string, SUBSCRIBER_BUFFER_SIZE),
		0}
}

func (s *Subscriber) send(log string) {
	if s.drops > 0 {
		panic("not implemented")
		// How to send the drop record and `log` at the same time?
		// ...
		// s.drops = 0
	}
	select {
	case s.ch <- log:
	default:
		// Slow subscriber detected.
		s.drops += 1
	}
}

func (s *Subscriber) Close() {
	close(s.ch)
}

func (s *Subscriber) GetID() string {
	return s.id
}

func (s *Subscriber) Logs() chan string {
	return s.ch
}
