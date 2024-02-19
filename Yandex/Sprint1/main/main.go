package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	for i := a; i > 1; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < b; i++ {
		fmt.Println("*")
	}
}
