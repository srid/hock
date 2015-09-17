package main

import (
	"fmt"
	"github.com/heroku/drain"
)

type Drain struct {
	name string
	*drain.Drain
}

func NewDrain(name string) *Drain {
	return &Drain{
		name,
		drain.NewDrain()}
}

func (d *Drain) Process() {
	for line := range d.Logs() {
		fmt.Println(line)
	}
}
