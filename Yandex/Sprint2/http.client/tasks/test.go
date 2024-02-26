package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8089/hello", nil)
	req.Header.Set("X-Request-ID", "123456789")

	resp, err := http.DefaultClient.Do(req)

	body, err := io.ReadAll(resp.Body)

	fmt.Println(string(body), err)
}
