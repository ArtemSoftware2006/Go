package main

func Map[T, U any](arr []T, transform func(T) U) []U {
	var result []U
	for _, item := range arr {
		result = append(result, transform(item))
	}

	return result
}
