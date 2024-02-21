package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var requestCount int = 0
var fibs = []int{0, 1, 1}

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	if requestCount < 3 {
		fmt.Fprint(w, fibs[requestCount])
	} else {
		tmp := fibs[1]
		fibs[1] = fibs[2]
		fibs[2] = tmp + fibs[1]
		fmt.Fprint(w, fibs[2])
	}
	requestCount++
}
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
}

func TimeMetricMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		log.Printf("Время выполнения запроса: %s", duration)
	})
}

func main() {
	mux := http.NewServeMux()

	fibonacciHandler := http.HandlerFunc(FibonacciHandler)

	mux.Handle("/", TimeMetricMiddleware(fibonacciHandler)) // вызываем TimeMetricMiddleware до вызова FibonacciHandler
	mux.Handle("/metrics", http.HandlerFunc(MetricsHandler))

	http.ListenAndServe(":9005", mux)
}
