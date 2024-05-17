package main

import (
	"errors"
	"sync"
)

type Summer interface {
	ProcessSum(
		summer func(arr []int, result chan<- int),
		nums []int,
		сhunkSize int) (int, error)
}

func SumChunk(arr []int, result chan<- int) {
	var sum int
	for _, item := range arr {
		sum += item
	}
	result <- sum
}

func ProcessSum(
	summer func(arr []int, result chan<- int),
	nums []int,
	сhunkSize int) (int, error) {

	if сhunkSize <= 0 {
		return 0, errors.New("chunkSize <= 0 (It must be > 0)")
	}

	sumChan := make(chan int)
	var wg sync.WaitGroup

	var i int
	for i = сhunkSize; i < len(nums); i = i + сhunkSize {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			summer(nums[i-сhunkSize:i], sumChan)
		}(i)
	}

	if i-сhunkSize < len(nums) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			summer(nums[i-сhunkSize:], sumChan)
		}()
	}

	go func() {
		wg.Wait()
		close(sumChan)
	}()

	var sum int
	for s := range sumChan {
		sum += s
	}

	return sum, nil
}
