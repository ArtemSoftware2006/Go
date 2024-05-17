package task3

import (
	"bufio"
	"os"
	"strconv"
)

func SumValuesPipeline(filename string) int {
	fileChan := NumbersGen(filename)
	filterChan := Filter(fileChan)
	return Sum(filterChan)
}
func NumbersGen(filename string) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		file, err := os.Open(filename)
		if err != nil {
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			num, err := strconv.Atoi(scanner.Text())
			if err == nil {
				out <- num
			}
		}
		if err := scanner.Err(); err != nil {
			return
		}
	}()

	return out
}
func Filter(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for num := range in {
			if num%2 == 0 {
				out <- num
			}
		}
	}()

	return out
}
func Sum(in <-chan int) int {
	sum := 0
	for num := range in {
		sum += num
	}

	return sum
}
