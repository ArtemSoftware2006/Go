package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func decode() {
	file, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}

	var reader io.Reader = bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	var person Person
	err = decoder.Decode(&person)
	if err != nil {
		panic(err)
	}

	fmt.Println(person)
}
