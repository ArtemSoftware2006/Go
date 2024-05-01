package main

import "time"

func main() {
	go sleeper()
}

func sleeper() {
	time.Sleep(2 * time.Second)
}
