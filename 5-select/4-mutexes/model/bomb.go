package model

import "sync"

type Bomb struct {
	mu         sync.Mutex
	holder     string
	isExploded bool
	timeLeft   float64
}
