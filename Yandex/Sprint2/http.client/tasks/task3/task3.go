package main

import (
	"context"
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("X-Request-ID")

	if reqId == nil {
		w.Write([]byte("Hello! RequestID not found"))
	} else {
		fmt.Fprintf(w, "Hello! RequestID: %s", reqId)
	}
}

func RequestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqId := r.Header.Get("X-Request-ID")

		if reqId != "" {
			ctx := context.WithValue(r.Context(), "X-Request-ID", reqId)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	}
}

func main() {
	http.Handle("/hello", RequestIDMiddleware(http.HandlerFunc(HelloHandler)))

	http.ListenAndServe(":8089", nil)
}
