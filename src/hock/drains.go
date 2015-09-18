package main

import (
	"sync"
)

var drains map[string]*Drain
var drainsMutex sync.Mutex

func getOrCreateDrain(key string) *Drain {
	if drn, ok := drains[key]; ok {
		return drn
	} else {
		// Delegate to the more expensive function
		return getOrCreateDrainSafe(key)
	}
}

func getOrCreateDrainSafe(key string) *Drain {
	drainsMutex.Lock()
	defer drainsMutex.Unlock()

	if drn, ok := drains[key]; ok {
		return drn
	} else {
		drn := NewDrain(key)
		drains[key] = drn
		go drn.Run()
		return drn
	}
}

func init() {
	drains = make(map[string]*Drain)
}
