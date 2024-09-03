package main

import (
	"sync"
	"testing"
)

func ReverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func TestReverseString(t *testing.T) {
	cases := []struct {
		name  string
		want  string
		value string
	}{
		{
			name:  "base str",
			value: "abc",
			want:  "cba",
		},
		{
			name:  "empty str",
			value: "",
			want:  "",
		},
	}

	var wg sync.WaitGroup

	wg.Add(len(cases))

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			defer wg.Done()

			got := ReverseString(tc.value)

			if got != tc.want {
				t.Errorf("ReverseString(%v) = %v; want %v", tc.value, got, tc.want)
			}
		})
	}

	wg.Wait()
}
