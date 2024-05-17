package main

import (
	"testing"
	"time"
)

// creating a test function for the timeoutFibonacci function
func TestTimeoutFibonacci(t *testing.T) {
	testCases := []struct {
		input    int
		timeout  time.Duration
		expected int
	}{
		{10, 1 * time.Second, 55},
		{40, 5 * time.Second, 102334155},
	}

	//iterating over every testcases
	for _, tc := range testCases {
		result, err := TimeoutFibonacci(tc.input, tc.timeout)
		if err != nil {
			t.Errorf("timeoutFibonacci(%d, %v) returned an error: %v", tc.input, tc.timeout, err)
		}
		if result != tc.expected {
			t.Errorf("timeoutFibonacci(%d, %v) = %d; want %d", tc.input, tc.timeout, result, tc.expected)
		}
	}

}
