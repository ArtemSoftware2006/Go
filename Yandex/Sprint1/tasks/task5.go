package main

import (
	"bufio"
	"fmt"
	"os"
)

func task5() {

	var length int
	fmt.Scan(&length)

	names := make([]string, length+1)

	scan := bufio.NewReader(os.Stdin)
	for i := 0; i < length; i++ {
		fmt.Println(names)
		names[i], _ = scan.ReadString('\n')
	}

	for {
		str, _ := scan.ReadString('\n')
		if str == "" {
			break
		}

		res := BinarySearch(names, 0, rune(str[0]))

		for i := 1; i < len(str); i++ {
			res = BinarySearch(names, i, rune(str[i]))
		}

		for i := 0; i < len(res); i++ {
			fmt.Println(res[i])
		}
	}

}

func BinarySearch(names []string, letterNumber int, target rune) []string {
	left, rigth := 0, len(names)-1

	for left <= rigth {
		middle := (left + rigth) / 2

		if names[letterNumber][middle] == byte(target) {
			res := make([]string, rigth-left+1)
			for i := left; i <= rigth; i++ {
				res = append(res, names[i])
			}
			return res
		} else if names[middle][letterNumber] < byte(target) {
			left = middle + 1
		} else {
			rigth = middle - 1
		}
	}

	return []string{}
}
