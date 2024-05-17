package task4

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
)

func getMark(name string) (int, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8082/mark?name=%s", name))
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		var mark int
		_, err = fmt.Fscanf(resp.Body, "%d", &mark)
		return mark, nil
	case http.StatusNotFound:
		return 0, errors.New("Student not found")
	case http.StatusInternalServerError:
		return 0, errors.New("Internal server error")
	default:
		return 0, fmt.Errorf("Unexpected status code %d", resp.StatusCode)
	}
}

func getMarks(names []string) ([]int, error) {
	var wg sync.WaitGroup
	var marks = make([]int, len(names))
	var errors = make([]error, len((names)))
	for i, name := range names {
		wg.Add(1)
		go func(i int, name string) {
			marks[i], errors[i] = getMark(name)
			defer wg.Done()
		}(i, name)
	}

	wg.Wait()

	for _, err := range errors {
		if err != nil {
			return nil, err
		}
	}
	return marks, nil
}

func CompareList(names []string) (map[string]string, error) {
	var marks = make([]int, len(names))

	marks, err := getMarks(names)
	if err != nil {
		return nil, err
	}

	average := Average(marks)
	compareList := make(map[string]string)

	for i, name := range names {
		if marks[i] > average {
			compareList[name] = ">"
		} else if marks[i] == average {
			compareList[name] = "="
		} else {
			compareList[name] = "<"
		}
	}

	return compareList, nil
}

func Average(marks []int) int {
	sum := 0

	for _, mark := range marks {
		sum += mark
	}

	return sum / len(marks)
}
