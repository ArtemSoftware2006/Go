package main

import (
	"fmt"
	"net/http"
)

func main() {
	fibs := []int{0, 1, 1}
	countRequest := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if countRequest < 3 {
			fmt.Fprint(w, fibs[countRequest])
			countRequest++
		} else {
			tmp := fibs[1]
			fibs[1] = fibs[2]
			fibs[2] = tmp + fibs[1]
			fmt.Fprint(w, fibs[2])
		}
	})
	http.ListenAndServe(":9005", nil)
}
