package main

import (
	"encoding/json"
	"reflect"
)

func CompareJSON(json1, json2 []byte) (bool, error) {
	var data1 interface{}
	var data2 interface{}

	err := json.Unmarshal(json1, &data1)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(json2, &data2)
	if err != nil {
		return false, err
	}

	if reflect.DeepEqual(data1, data2) {
		return true, nil
	} else {
		return false, nil
	}
}
