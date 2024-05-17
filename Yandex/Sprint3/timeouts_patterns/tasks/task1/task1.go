package main

import "time"

func SieveOfEratosthenes(max int) []int {

	isPrime := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= max; p++ {

		if isPrime[p] {

			for i := p * p; i <= max; i += p {
				isPrime[i] = false
			}
		}
	}

	var primes []int
	for p := 2; p <= max; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}

	return primes
}

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	primeNums := SieveOfEratosthenes(N)
	defer close(prime_nums)

	for _, n := range primeNums {
		select {
		case <-stop:
			return
		default:
			time.Sleep(time.Millisecond)
			prime_nums <- n
		}
	}
}
