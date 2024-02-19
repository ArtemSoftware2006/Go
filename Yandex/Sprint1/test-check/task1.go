package main

import (
	"fmt"
)

func task1() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	for i := b; i > 1; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < a; i++ {
		fmt.Println("*")
	}
}
