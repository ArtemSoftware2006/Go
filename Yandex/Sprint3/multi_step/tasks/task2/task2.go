package main

import "fmt"

func ToString[T any](done <-chan struct{}, valueStream <-chan T) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for val := range valueStream {
			select {
			case <-done:
				return
			default:
				out <- fmt.Sprint(val)
			}
		}
	}()
	return out
}
