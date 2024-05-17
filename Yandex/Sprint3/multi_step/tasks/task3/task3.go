package main

import (
	"encoding/csv"
	"os"
)

func ReadCSV(file string) (<-chan []string, error) {

	fileDisript, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fileDisript.Close()

	reader := csv.NewReader(fileDisript)

	strs, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	out := make(chan []string)

	go func() {
		defer close(out)
		for _, str := range strs {
			out <- str
		}
	}()

	return out, nil
}
