package main

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		// имя теста
		a    int
		b    int
		want int
	}{
		// тестовые данные №1
		{
			a:    2,
			b:    1,
			want: 3,
		},
		// тестовые данные №2
		{
			a:    -3,
			b:    3,
			want: 0,
		},
	}
	// перебор всех тестов
	for _, tc := range cases {
		tc := tc
		// запуск отдельного теста
		t.Run("sum test", func(t *testing.T) {
			// тестируе,м функцию Sum
			got := Sum(tc.a, tc.b)
			// проверим полученное значение
			if got != tc.want {
				t.Errorf("Sum(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
