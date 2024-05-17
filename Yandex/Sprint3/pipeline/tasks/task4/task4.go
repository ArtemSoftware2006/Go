package main

import (
	"fmt"
	"sync"
)

func MultiplyPipeline(inputNums ...[]int) int {
	var wg sync.WaitGroup
	multies := make([]int, len(inputNums))

	for i, nums := range inputNums {
		wg.Add(1)
		go func(nums []int, i int) {
			defer wg.Done()
			numsChan := NumbersGen(nums...)
			filterChan := Filter(numsChan)
			multies[i] = Multiply(filterChan)
		}(nums, i)
	}

	wg.Wait()

	totalMulty := 1
	for _, multy := range multies {
		totalMulty *= multy
	}
	return totalMulty
}
func NumbersGen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()
	return out
}
func Filter(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for num := range in {
			if num > 0 {
				out <- num
			}
		}
	}()

	return out
}
func Multiply(in <-chan int) int {
	multi := 1
	for num := range in {
		multi *= num
	}
	return multi
}

func main() {
	result := MultiplyPipeline([]int{1, 2, 3}, []int{-1, -2, -3})
	fmt.Println(result)
}
