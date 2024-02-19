package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	//array := []string{"one", "two", "three"}
	var arrayInt = []int{12, 234, 12}

	nums := make([]int, 23)

	fmt.Println(nums)

	arrayInt = append(arrayInt, 1234)
	sort.Ints(arrayInt)

	for _, i := range arrayInt {
		fmt.Println(i)
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(100)
	}

	fmt.Println("\n Рандомный массив \n")

	for _, i := range nums {
		fmt.Print(i, " ")
	}

	fmt.Println("\n Рандомный массив (первые 5 элементов) \n")

	for _, item := range nums[:5] {
		fmt.Print(item, " ")
	}

	nums = append(nums, 12)

	nums = append(nums[:3], nums[5:]...) // Удаление элементов из массива

	fmt.Println()

	for _, item := range nums {
		fmt.Print(item, " ")
	}
}
