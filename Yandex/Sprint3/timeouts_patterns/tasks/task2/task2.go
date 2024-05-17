package main

import (
	"context"
	"fmt"
	"io"
	"os"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	defer close(result)
	out := make(chan []byte)
	go work(path, out)
	for data := range out {
		select {
		case <-ctx.Done():
			fmt.Println("Время операции истекло или было отменено.")
			return
		default:
			result <- data
		}
	}
}

func work(path string, subOut chan []byte) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return
	}

	defer close(subOut)
	defer file.Close()
	subOut <- data
}
