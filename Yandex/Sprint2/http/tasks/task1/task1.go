package main

import (
	"fmt"
	"net/http"
	"regexp"
)

var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			fmt.Fprintf(w, "hello stranger")
		} else {
			if IsLetter(name) {
				fmt.Fprintf(w, "Hello %s", name)
			} else {
				fmt.Fprintf(w, "hello dirty hacker")
			}
		}
	})

	http.ListenAndServe(":9005", nil)
}
