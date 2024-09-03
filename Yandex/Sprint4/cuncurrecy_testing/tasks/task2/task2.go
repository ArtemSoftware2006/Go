package main

import (
	"sync"
	"testing"
)

func Contains(numbers []int, target int) bool {
	for _, num := range numbers {
		if num == target {
			return true
		}
	}
	return false
}

func TestContains(t *testing.T) {
	cases := []struct {
		name  string
		want  bool
		value struct {
			target int
			values []int
		}
	}{
		{
			name: "base slice with target",
			want: true,
			value: struct {
				target int
				values []int
			}{
				target: 1,
				values: []int{1, 2, 3},
			},
		},
		{
			name: "base slice without target",
			want: false,
			value: struct {
				target int
				values []int
			}{
				target: 100,
				values: []int{1, 2, 3},
			},
		},
		{
			name: "slice is nil",
			want: true,
			value: struct {
				target int
				values []int
			}{
				target: 1,
				values: nil,
			},
		},
	}

	var wg sync.WaitGroup

	for _, tc := range cases {
		tc := tc

		wg.Add(len(cases))

		t.Run(tc.name, func(t *testing.T) {
			got := Contains(tc.value.values, tc.value.target)

			if got != tc.want {
				t.Errorf("Contains(%v, %v) = %v; want %v", tc.value.values, tc.value.target, got, tc.want)
			}
		})
	}
}
