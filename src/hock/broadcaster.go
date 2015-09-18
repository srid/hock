package main

import ()

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
		case id := <-b.unsubscribeRequests:
			b.subscribers[id].Close()
			delete(b.subscribers, id)
		case log := <-b.ch:
			for _, sub := range b.subscribers {
				sub.send(log)
			}
		}
	}
}
