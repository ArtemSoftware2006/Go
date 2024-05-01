package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendDataToRoutine(ch)
	fmt.Println(<-ch)
}

func sendDataToRoutine(ch chan string) {
	ch <- "Hello from the channel!"
}
