package main

import (
	"fmt"
	"math"
)

func task2() {
	var number int
	fmt.Scanln(&number)

	for i := 3; i < number; i += 5 {
		if isPrime(i) {
			fmt.Print("хоп ")
		} else {
			fmt.Print(i, " ")
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	maxDivisor := int(math.Sqrt(float64(n)))
	for i := 2; i <= maxDivisor; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
