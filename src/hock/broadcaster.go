package main

import (
	"github.com/pborman/uuid"
)

type Broadcaster struct {
	ch              chan string
	subscribers     map[string]*Subscriber
	outgressAdds    chan *Subscriber
	outgressAddRets chan string
	outgressDels    chan string
}

func NewBroadcaster(bufferSize int) *Broadcaster {
	return &Broadcaster{
		make(chan string, bufferSize),
		make(map[string]*Subscriber),
		make(chan *Subscriber),
		make(chan string),
		make(chan string)}
}

func (b *Broadcaster) Subscribe() (string, *Subscriber) {
	sub := NewSubscriber()
	b.outgressAdds <- sub
	id := <-b.outgressAddRets
	return id, sub
}

func (b *Broadcaster) Unsubscribe(id string) {
	b.outgressDels <- id
}

func (b *Broadcaster) Broadcast(log string) {
	b.ch <- log
}

func (b *Broadcaster) Run() {
	for {
		select {
		case sub := <-b.outgressAdds:
			id := uuid.New()
			b.subscribers[id] = sub
			b.outgressAddRets <- id
		case id := <-b.outgressDels:
			sub := b.subscribers[id]
			sub.Close()
			delete(b.subscribers, id)
		case log := <-b.ch:
			for _, sub := range b.subscribers {
				sub.send(log)
			}
		}
	}
}
