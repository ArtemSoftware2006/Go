package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("Request %s %s", r.Method, r.URL)

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		log.Printf("Time taken: %s", duration)
	})
}

func panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Panicln(err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	mux := http.NewServeMux()
	hello := http.HandlerFunc(helloHandler)
	mux.Handle("/", panicMiddleware(loggingMiddleware(hello)))

	if err := http.ListenAndServe(":9005", mux); err != nil {
		log.Fatal(err)
	}
}
