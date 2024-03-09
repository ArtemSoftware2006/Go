package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type State struct {
	Fill int `json:"fill"`
}

func setStateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var state State
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = os.WriteFile("state.cfg", []byte(strconv.Itoa(state.Fill)+"%"), 0644)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "State updated successfully")
}

func main() {
	http.HandleFunc("/setstate", setStateHandler)

	fmt.Println("Server running on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
