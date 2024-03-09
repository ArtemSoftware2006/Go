package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		data, err := ioutil.ReadFile("state.cfg")
		if err != nil {
			http.Error(w, "Failed to read state.cfg file", http.StatusInternalServerError)
			return
		}

		fill := strings.TrimSpace(string(data))

		response := fmt.Sprintf(`{"fill":%s}`, fill[0:len(fill)-1])
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})

	http.ListenAndServe(":8081", nil)
}
