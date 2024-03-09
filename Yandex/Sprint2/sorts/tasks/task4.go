package main

import (
	"sort"
)

func SortByFreq(str string) string {
	freq := make(map[rune]int)
	var uniqueRunes []rune

	for _, char := range str {
		if _, ok := freq[char]; !ok {
			uniqueRunes = append(uniqueRunes, char)
		}
		freq[char]++
	}

	sort.Slice(uniqueRunes, func(i, j int) bool {
		if freq[uniqueRunes[i]] == freq[uniqueRunes[j]] {
			return uniqueRunes[i] < uniqueRunes[j]
		}
		return freq[uniqueRunes[i]] < freq[uniqueRunes[j]]
	})

	var result string
	for _, char := range uniqueRunes {
		for i := 0; i < freq[char]; i++ {
			result += string(char)
		}
	}

	return result
}
