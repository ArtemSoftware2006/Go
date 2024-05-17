package main

import (
	"testing"
)

func TestProcessSum(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		chunkSize int
		expected  int
	}{
		{
			name:      "EmptySlice",
			nums:      []int{},
			chunkSize: 5,
			expected:  0,
		},
		{
			name:      "SingleChunk",
			nums:      []int{1, 2, 3, 4, 5},
			chunkSize: 5,
			expected:  15,
		},
		{
			name:      "MultipleChunks",
			nums:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			chunkSize: 3,
			expected:  55,
		},
		{
			name:      "UnevenChunks",
			nums:      []int{1, 2, 3, 4, 5},
			chunkSize: 2,
			expected:  15,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := ProcessSum(SumChunk, test.nums, test.chunkSize)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
