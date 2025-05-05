package model

import (
	"log"
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
	randHolder := randHoldersSelector()
	duration := float64(rand.Intn(randTime) + randTimeMin)

	log.Printf("New bomb created by %s with duration %.2fs", randHolder, duration)

	return &Bomb{
		mu:         sync.Mutex{},
		holder:     randHolder,
		isExploded: false,
		timeLeft:   duration,
	}

}

func randHoldersSelector() string {
	randValue := rand.Intn(1) + 1

	if randValue == 1 {
		return hostHolder
	}
	return clientHolder
}

func (b *Bomb) DecreaseTime(holdingTime float64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.timeLeft -= holdingTime

	if b.timeLeft <= 0 {
		b.isExploded = true
		log.Printf("Bomb exploded in hands of %s!", b.holder)
	}
}

func (b *Bomb) SwitchHolder() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.holder == hostHolder {
		b.holder = clientHolder
	} else {
		b.holder = hostHolder
	}
	log.Printf("Turn switched. New holder: %s", b.holder)
}

func (b *Bomb) IsExploded() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.isExploded {
		return true
	}
	return false
}

func (b *Bomb) SetExploded() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.isExploded = true

}

func (b *Bomb) WhoHold() string {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.holder == hostHolder {
		return hostHolder
	} else {
		return clientHolder
	}
}

func (b *Bomb) Snapshot() Bomb {
	b.mu.Lock()
	defer b.mu.Unlock()
	return Bomb{
		holder:     b.holder,
		isExploded: b.isExploded,
		timeLeft:   b.timeLeft,
	}

}
