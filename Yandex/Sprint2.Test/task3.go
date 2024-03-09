package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func SumUp(filepath, colname string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		return 0, err
	}

	colIndex := -1
	for i, name := range header {
		if name == colname {
			colIndex = i
			break
		}
	}

	if colIndex == -1 {
		return 0, nil
	}

	sum := 0

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		val, err := strconv.Atoi(record[colIndex])
		if err != nil {
			return 0, err
		}
		sum += val
	}

	return sum, nil
}
