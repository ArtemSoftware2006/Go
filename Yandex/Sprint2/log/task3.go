package main

import (
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Логгируем информацию о запросе
	log.Printf("%s %s %s %s %s", start.Format("2006/01/02 15:04:05"), r.Method, r.URL.Path, r.RemoteAddr, r.Header)

	// Имитация обработки запроса
	time.Sleep(10 * time.Millisecond)

	// Логгируем время выполнения запроса
	log.Printf("%s %s %s %s %v", start.Format("2006/01/02 15:04:05"), r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
}

func main() {
	http.HandleFunc("/", handler)

	port := ":8080"
	log.Printf("Starting server on port %s", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
