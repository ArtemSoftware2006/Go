package main

import (
	"net/http"
	"regexp"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	w.Write([]byte("Hello, " + name + "!"))
}

func sanitizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		isValid := regexp.MustCompile("^[^!@#$%^&*()_+]+$")
		if isValid.MatchString(name) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid name", http.StatusInternalServerError)
		}
	})
}

func setDefaultNameMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if len(name) == 0 {
			params := r.URL.Query()
			params.Set("name", "stanger")
			r.URL.RawQuery = params.Encode()
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	greet := http.HandlerFunc(greetHandler)

	mux.HandleFunc("/", greet)

	handler := setDefaultNameMiddleware(sanitizeMiddleware(greet))

	http.ListenAndServe(":9005", handler)

}
