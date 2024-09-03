package main

import (
	"testing"
)

func TestLength(t *testing.T) {
	cases := []struct {
		name  string
		want  string
		value int
	}{
		{
			name:  "less 0",
			value: -1,
			want:  "negative",
		},
		{
			name:  "equal 0",
			value: 0,
			want:  "zero",
		},
		{
			name:  "less 10",
			value: 9,
			want:  "short",
		},
		{
			name:  "less 100",
			value: 99,
			want:  "long",
		},
		{
			name:  "more 100",
			value: 101,
			want:  "very long",
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := Length(tc.value)

			if got != tc.want {
				t.Errorf("Length(%v) = %v; want %v", tc.value, got, tc.want)
			}
		})
	}
}

func Length(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "short"
	case a < 100:
		return "long"
	}
	return "very long"
}
