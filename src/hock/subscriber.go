package main

type Subscriber struct {
	ch    chan string
	drops int
}

func NewSubscriber() *Subscriber {
	return &Subscriber{
		make(chan string),
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
