package main

func Filter[T any](arr []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range arr {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
