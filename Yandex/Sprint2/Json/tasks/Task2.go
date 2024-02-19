package main

import (
	"encoding/json"
)

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {

	persons := make([]Person, 0)

	for _, jsonData := range jsonDataList {
		var person []Person
		err := json.Unmarshal(jsonData, &person)
		if err != nil {
			return nil, err
		}

		for _, item := range person {
			persons = append(persons, item)
		}
	}

	jsonNewData, err := json.Marshal(persons)
	if err != nil {
		return nil, err
	}

	return jsonNewData, nil
}
