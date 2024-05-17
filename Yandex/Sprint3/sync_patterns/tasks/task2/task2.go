package main

import "context"

type sequenced interface {
	getSequence() int
}

func getSequence[T sequenced](num T) int {
	return num.getSequence()
}

func EvenNumbersGen[T sequenced](ctx context.Context, numbers ...T) <-chan T {
	outChan := make(chan T)

	go func() {
		defer close(outChan)
		for _, num := range numbers {
			select {
			case <-ctx.Done():
				return
			default:
				if getSequence(num)%2 == 0 {
					outChan <- num
				}
			}
		}
	}()

	return outChan
}
