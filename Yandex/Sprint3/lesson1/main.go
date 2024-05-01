package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumGoroutine())
	chanel()
}

func chanel() {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)

		ch <- 123
	}()

	fmt.Println(<-ch)

	notMitex()
}

func notMitex() {
	data := make(map[string]int)

	go func() {
		for i := 0; i < 1000; i++ {
			data["key"] = i
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(data["key"])
		}
	}()

	time.Sleep(time.Second) // Ожидание завершения работы горутин
}
