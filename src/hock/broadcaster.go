package main

import (
	log "github.com/Sirupsen/logrus"
)

type Broadcaster struct {
	ch                  chan string
	subscribers         map[string]*Subscriber
	subscribeRequests   chan *Subscriber
	unsubscribeRequests chan string
}

func NewBroadcaster(bufferSize int) *Broadcaster {
	return &Broadcaster{
		make(chan string, bufferSize),
		make(map[string]*Subscriber),
		make(chan *Subscriber),
		make(chan string)}
}

func (b *Broadcaster) Subscribe() *Subscriber {
	sub := NewSubscriber()
	b.subscribeRequests <- sub
	return sub
}

func (b *Broadcaster) Unsubscribe(sub *Subscriber) {
	b.unsubscribeRequests <- sub.GetID()
}

func (b *Broadcaster) Broadcast(log string) {
	b.ch <- log
}

func (b *Broadcaster) Run() {
	for {
		select {
		case sub := <-b.subscribeRequests:
			b.subscribers[sub.GetID()] = sub
			log.WithFields(log.Fields{
				"count#hock.subscribers": len(b.subscribers),
			}).Info("Created new subscriber '%s'", sub.GetID())
		case id := <-b.unsubscribeRequests:
			b.subscribers[id].Close()
			delete(b.subscribers, id)
			log.WithFields(log.Fields{
				"count#hock.subscribers": len(b.subscribers),
			}).Info("Deleted subscriber '%s'", id)
		case log := <-b.ch:
			for _, sub := range b.subscribers {
				sub.send(log)
			}
		}
	}
}
