package main

import (
	"time"
)

type Yujin struct {
	lastCoffeeAt time.Time
}

func (y *Yujin) IsProductive() bool {
	return time.Since(y.lastCoffeeAt) < (time.Hour * 3)
}
