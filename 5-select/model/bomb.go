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
	mu       sync.Mutex `json:"-"` // gob için encode edilmez
	TimeLeft float64
	Holder   string
	Exploded bool
}

func NewBomb() *Bomb {
	randHolder := randHoldersSelector()
	duration := float64(rand.Intn(randTime) + randTimeMin)

	log.Printf("New bomb created by %s with duration %.2fs", randHolder, duration)

	return &Bomb{
		mu:       sync.Mutex{},
		Holder:   randHolder,
		Exploded: false,
		TimeLeft: duration,
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

	b.TimeLeft -= holdingTime

	if b.TimeLeft <= 0 {
		b.Exploded = true
		log.Printf("Bomb exploded in hands of %s!", b.Holder)
	}
}

func (b *Bomb) SwitchHolder() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.Holder == hostHolder {
		b.Holder = clientHolder
	} else {
		b.Holder = hostHolder
	}
	log.Printf("Turn switched. New holder: %s", b.Holder)
}

func (b *Bomb) IsExploded() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.Exploded {
		return true
	}
	return false
}

func (b *Bomb) SetExploded() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Exploded = true

}

func (b *Bomb) WhoHold() string {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.Holder == hostHolder {
		return hostHolder
	} else {
		return clientHolder
	}
}

func (b *Bomb) Snapshot() Bomb {
	b.mu.Lock()
	defer b.mu.Unlock()
	return Bomb{
		Holder:   b.Holder,
		Exploded: b.Exploded,
		TimeLeft: b.TimeLeft,
	}

}
