package main

import (
	"encoding/json"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {
	var students []Student

	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(students); i++ {
		students[i].Grade++
	}

	var newJSONData []byte
	newJSONData, err = json.Marshal(&students)

	if err != nil {
		return nil, err
	}

	return newJSONData, nil
}
