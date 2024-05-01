package main

import (
	"sync"
)

var (
	mu    sync.Mutex
	myMap = make(map[int]int)
)

func main() {

}

func writeToMap(key, value int) {
	mu.Lock()
	myMap[key] = value
	mu.Unlock()
}

func readFromMap(key int) {
	mu.Lock()

	mu.Unlock()
}
