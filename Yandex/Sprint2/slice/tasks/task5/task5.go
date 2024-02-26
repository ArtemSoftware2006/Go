package main

func Mix(nums []int) []int {
	result := make([]int, len(nums))

	i, j, t := 0, 0, 0
	for t < len(nums) {
		if t%2 == 0 {
			result[i] = nums[j]
			i += 2
		} else {
			result[i-1] = nums[len(nums)/2+j]
			j++
		}
		t++
	}

	return result
}
