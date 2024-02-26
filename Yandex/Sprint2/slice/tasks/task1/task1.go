package main

import (
	"errors"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("n не может быть отрицательным")
	}

	var result []int
	for _, num := range nums {
		if num < limit {
			result = append(result, num)
			if len(result) == n {
				break
			}
		}
	}

	return result, nil
}
