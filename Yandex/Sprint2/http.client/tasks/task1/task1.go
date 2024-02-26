package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println(SendHTTPRequest("https://ya.ru"))
}

func SendHTTPRequest(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
