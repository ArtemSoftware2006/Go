package main

import (
	"fmt"
	"net/http"
)

func queryParamHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	path := r.URL.Path

	fmt.Printf("Path: %s\n", path)
	fmt.Printf("Name: %s, Age: %s\n", name, age)

	fmt.Fprintf(w, "Name: %s, Age: %s\n", name, age)
}

func main() {
	fmt.Println("Hello World!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	http.HandleFunc("/query", queryParamHandler)

	http.ListenAndServe(":9005", nil)
}
