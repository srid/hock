package main

import (
	"fmt"
	"github.com/heroku/drain"
)

type Drain struct {
	name        string
	broadcaster *Broadcaster
	*drain.Drain
}

func NewDrain(name string) *Drain {
	return &Drain{
		name,
		NewBroadcaster(1500),
		drain.NewDrain()}
}

func (d *Drain) Run() {
	go d.broadcaster.Run()
	for line := range d.Logs() {
		d.broadcaster.Broadcast(fmt.Sprintf("%+v", line))
	}
}
