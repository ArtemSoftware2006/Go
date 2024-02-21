package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	var length int
	fmt.Scan(&length)

	names := make([]string, length+1)

	scan := bufio.NewReader(os.Stdin)
	for i := 0; i < length; i++ {
		names[i], _ = scan.ReadString('\n')
	}

	fmt.Println(names)

	for {
		str, _ := scan.ReadString('\n')

		if len(str) == 2 {
			break
		}

		letter, _ := utf8.DecodeRuneInString(str)
		fmt.Println(letter)

		res := BinarySearch2(names, 0, letter)

		for i := 0; i < len(res); i++ {
			fmt.Print(res[i])
		}
	}

}

func BinarySearch2(names []string, letterNumber int, target rune) []string {
	left, rigth := 0, len(names)-1

	for left <= rigth {
		middle := (left + rigth) / 2

		letter, _ := utf8.DecodeRuneInString(names[middle])

		if letter == target {
			res := make([]string, rigth-left+1)
			for i := left; i < rigth; i++ {
				res[i] = names[i]
				fmt.Println(i)
			}

			return res
		} else if letter < target {
			left = middle + 1
		} else {
			rigth = middle - 1
		}
	}

	return []string{}
}
