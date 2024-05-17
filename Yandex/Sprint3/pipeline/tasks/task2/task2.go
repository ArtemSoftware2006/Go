package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	// Пример использования функции NumbersGen
	filename := "numbers.txt"
	numbersChan := NumbersGen(filename)
	for num := range numbersChan {
		fmt.Println("Прочитано число из файла:", num)
	}
}
