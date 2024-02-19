package main

import (
	"fmt"
	"log"
)

func main() {

	text := []string{"one", "two", "three"}
	for h, i := range text {
		fmt.Println(h, i)
	}
	i := 0
	for {
		if i == 2 {
			i++
			log.Fatal("fatal", i)
			continue
		}
		if i == 7 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
		i += 1
	}

}
