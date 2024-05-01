package main

import (
	"sync"
	"time"
)

var (
	counter = 0
	mu      sync.Mutex
)

func main() {
	go incrementCounter()
	go incrementCounter()
	go incrementCounter()

	time.Sleep(1 * time.Second)
}

func incrementCounter() {
	mu.Lock()
	counter++
	mu.Unlock()
}
