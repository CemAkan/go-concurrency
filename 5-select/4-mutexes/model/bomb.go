package model

import (
	"math/rand"
	"sync"
)

var (
	hostHolder   = "host"
	clientHolder = "client"
	randTimeMin  = 20
	randTime     = 60
)

type Bomb struct {
	mu         sync.Mutex
	holder     string
	isExploded bool
	timeLeft   float64
}

func NewBomb() *Bomb {
	return &Bomb{
		mu:         sync.Mutex{},
		holder:     randHoldersSelector(),
		isExploded: false,
		timeLeft:   float64(rand.Intn(randTime) + randTimeMin),
	}
}

func randHoldersSelector() string {
	randValue := rand.Intn(1) + 1

	if randValue == 1 {
		return hostHolder
	}
	return clientHolder
}
