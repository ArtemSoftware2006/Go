package main

import (
	"Sprint2Json/task4"
)

func main() {

	test4_data := []byte(`
	[
		{
			"name": "Анна",
			"admission_date": "2021-09-01T00:00:00Z"
		},
		{
			"name": "Иван",
			"admission_date": "2022-03-15T00:00:00Z"
		},
		{
			"name": "Мария",
			"admission_date": "2023-01-10T00:00:00Z"
		}
	]`)

	task4.ProcessJSON(test4_data)
}
