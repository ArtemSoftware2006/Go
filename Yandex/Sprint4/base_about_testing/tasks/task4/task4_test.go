package main

import (
	"testing"
)

func TestDeleteVowels(t *testing.T) {
	cases := []struct {
		name  string
		want  string
		value string
	}{
		{
			name:  "word with vowels",
			value: "AaOoEeUuIi",
			want:  "",
		},
		{
			name:  "word with Yy",
			value: "Yy",
			want:  "Yy",
		},
		{
			name:  "word with consonants",
			value: "QqWwRrTtPpSsDdFfGgHhJjKkLlZzXxCcVvBbNnMm",
			want:  "QqWwRrTtPpSsDdFfGgHhJjKkLlZzXxCcVvBbNnMm",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := DeleteVowels(tc.value)
			if got != tc.want {
				t.Errorf("DeleteVowels(%v) = %v; want %v", tc.value, got, tc.want)
			}
		})
	}
}
