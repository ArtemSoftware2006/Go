package main

import (
	"context"
	"sync"
)

func FanIn[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	outChan := make(chan T)
	wg := sync.WaitGroup{}
	for _, ch := range channels {
		wg.Add(1)
		go func(inChan <-chan T) {
			defer wg.Done()
			for {
				select {
				case data, ok := <-inChan:
					if !ok {
						return
					}
					outChan <- data
				case <-ctx.Done():
					return
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(outChan)
	}()

	return outChan
}
