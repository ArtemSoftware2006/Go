package main

import (
	"errors"
	"time"
)

func GetFib(number int) int {
	var fib1, fib2, fib3 int
	fib1 = 1
	fib2 = 1
	fib3 = 2

	switch number {
	case 1:
		return fib1
	case 2:
		return fib2
	}

	for i := 3; i < number; i++ {
		fib1 = fib2
		fib2 = fib3
		fib3 = fib1 + fib2
	}

	return fib3
}

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	out := make(chan int)
	go func(n int) {
		out <- GetFib(n)

	}(n)

	select {
	case res := <-out:
		return res, nil
	case <-time.After(timeout):
		return 0, errors.New("Timeout")
	}
}
