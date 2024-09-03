package main

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestSortIntegers(t *testing.T) {
	cases := []struct {
		name   string
		want   []int
		values []int
	}{
		{
			name:   "base integers",
			values: []int{4, 2, 1, -4},
			want:   []int{-4, 1, 2, 4},
		},
		{
			name:   "empty slice",
			values: []int{},
			want:   []int{},
		},
		{
			name:   "nil slice",
			values: nil,
			want:   nil,
		},
	}

	wg.Add(len(cases))

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			defer wg.Done()

			SortIntegers(tc.values)

			if !compareSlices(tc.want, tc.values) {
				t.Errorf("SortIntegers(%v) = %v; want %v", tc.values, tc.want, tc.want)
			}
		})
	}

	wg.Wait()
}

func compareSlices(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
