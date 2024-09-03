package main

import "testing"

func TestMultiply(t *testing.T) {
	cases := []struct {
		name string
		want int
		a    int
		b    int
	}{
		{
			name: "two positiv",
			want: 10,
			a:    1,
			b:    10,
		},
		{
			name: "two negative",
			want: 10,
			a:    -1,
			b:    -10,
		},
		{
			name: "one positiv, one negative",
			want: -10,
			a:    -1,
			b:    10,
		},
		{
			name: "one zero",
			want: 0,
			a:    0,
			b:    10,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			got := Multipy(tc.a, tc.b)

			if got != tc.want {
				t.Errorf("Multipy(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
