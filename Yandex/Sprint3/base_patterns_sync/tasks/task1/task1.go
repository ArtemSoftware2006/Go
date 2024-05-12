package task1

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
		_, err := fmt.Fscanf(resp.Body, "%d", &mark)
		if err != nil {
			return 0, err
		}
		return mark, nil
	case http.StatusNotFound:
		return 0, errors.New("student not found")
	case http.StatusInternalServerError:
		return 0, errors.New("internal server error")
	default:
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}

func Compare(name1, name2 string) (string, error) {
	var wg sync.WaitGroup
	var mark1, mark2 int
	var err1, err2 error

	wg.Add(2)
	go func() {
		defer wg.Done()
		mark1, err1 = getMark(name1)
	}()

	go func() {
		defer wg.Done()
		mark2, err2 = getMark(name2)
	}()

	wg.Wait()

	if err1 != nil {
		return "", err1
	}
	if err2 != nil {
		return "", err2
	}

	switch {
	case mark1 > mark2:
		return ">", nil
	case mark1 < mark2:
		return "<", nil
	default:
		return "=", nil
	}
}
