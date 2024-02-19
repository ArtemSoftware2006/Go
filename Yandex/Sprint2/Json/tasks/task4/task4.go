package task4

import (
	"encoding/json"
	"fmt"
	"time"
)

type Student struct {
	Name          string    `json:"name"`
	AdmissionDate time.Time `json:"admission_date"`
	DaysOnCourse  int       `json:"-"`
}

func ProcessJSON(jsonData []byte) error {
	var students []Student

	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		return err
	}
	//создай даты 1 октября 2023 года
	endDate := time.Date(2023, time.October, 2, 0, 0, 0, 0, time.Local)

	for _, student := range students {
		student.DaysOnCourse = int(endDate.Sub(student.AdmissionDate) / (time.Hour * 24))
		fmt.Printf("%s: %d\n", student.Name, student.DaysOnCourse)
	}

	return nil
}
