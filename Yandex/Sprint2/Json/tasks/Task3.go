package main

import "encoding/json"

type Person struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func SplitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	var persons []Person

	err := json.Unmarshal(jsonData, &persons)
	if err != nil {
		return nil, err
	}

	personsByClasses := make(map[string][]byte)

	for _, person := range persons {
		jsonPerson, err := json.Marshal(person)
		if err != nil {
			return nil, err
		}
		personsByClasses[person.Class] = append(personsByClasses[person.Class], jsonPerson...)
	}

	return personsByClasses, nil
}
