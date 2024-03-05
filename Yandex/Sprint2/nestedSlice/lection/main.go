package main

import "fmt"

func main() {
	fmt.Println(MaxValueOfExpression([]int{3, 2, 6, 3, 4, 5}))
}

func MaxValueOfExpression(nums []int) int {
	first := make([]int, len(nums)+1)
	for i := len(nums) - 1; i >= 0; i-- {
		first[i] = max(first[i+1], nums[i])
	}
	second := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		second[i] = max(second[i+1], first[i+1]-nums[i])
	}
	return second[0]
}
