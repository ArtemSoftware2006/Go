package main

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[T Number](arr []T) T {
	var sum T
	for _, v := range arr {
		sum += v
	}
	return sum
}
