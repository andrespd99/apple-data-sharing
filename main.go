package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {

	r := http.NewServeMux()

	r.HandleFunc("/healtz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		enc := json.NewEncoder(os.Stdout)
		fmt.Println("HEADERS:")
		enc.Encode(r.Header)
		fmt.Println()
		enc.Encode(r.Body)
		fmt.Println()
	})

	http.ListenAndServe(":8080", r)
}
