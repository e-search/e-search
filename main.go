package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		w.WriteHeader(http.StatusOK)
		encoder.Encode("Hello, World!")
	})
	http.ListenAndServe(port, nil)
}
