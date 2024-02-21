package main

import (
	"bufio"
	"fmt"
	"os"
)

func task3() {
	var letter1, letter2 string
	var str string

	fmt.Scanln(&letter1)
	fmt.Scanln(&letter2)

	in := bufio.NewReader(os.Stdin)
	str, _ = in.ReadString('\n')

	letters := make(map[string]int)

	for _, v := range str {
		if v != ' ' {
			letters[string(v)]++
		}
	}

	if letters[letter1] > letters[letter2] {
		fmt.Printf("%s %d\n", letter1, letters[letter1])
		fmt.Printf("%s %d\n", letter2, letters[letter2])
	} else {
		fmt.Printf("%s %d\n", letter2, letters[letter2])
		fmt.Printf("%s %d\n", letter1, letters[letter1])

	}
}
