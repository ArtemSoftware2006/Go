package main

import "fmt"

func Rotate(data []int, pos int) []int {

	pos = pos % len(data)

	if pos < 0 {
		pos = len(data) - pos*(-1)
	}

	firstArr := data[len(data)-pos:]
	secondArr := data[:len(data)-pos]

	return append(firstArr, secondArr...)
}

func main() {
	fmt.Println(Rotate([]int{1, 2, 3, 4, 5, 6, 7}, -10))
}
