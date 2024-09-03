package main

import (
	"testing"
)

func TestAreAnagrams(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Anagrams",
			args: args{"listen", "silent"},
			want: true,
		},
		{
			name: "Not Anagrams",
			args: args{"hello", "world"},
			want: false,
		},
		{
			name: "Different Length",
			args: args{"abc", "abcd"},
			want: false,
		},
		{
			name: "Case Insensitive",
			args: args{"Listen", "Silent"},
			want: true,
		},
		{
			name: "Empty Strings",
			args: args{"", ""},
			want: true,
		},
		{
			name: "Single Character",
			args: args{"a", "a"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreAnagrams(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("AreAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
