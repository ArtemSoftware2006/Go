package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var password string = "password"
var login string = "userid"

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if authHeader := r.Header.Get("Authorization"); authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusForbidden)
			return
		}

		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if pair[0] != login || pair[1] != password {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/answer/", Authorization(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The answer is 42"))
	})))

	http.ListenAndServe(":9005", mux)
}
