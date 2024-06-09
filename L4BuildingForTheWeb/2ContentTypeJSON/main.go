// Demo: Content Type JSON

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var dictionary = map[string]string{
	"Go":     "Programming language created by Google",
	"Gopher": "A software engineer who builds with Go",
	"Golang": "Another name for Go",
}

func getdictionary(w http.ResponseWriter, r *http.Request) {
	// build bussiness logic
	// we used the following line to set the Content-Type header to let the client know to expect JSON in the response:
	w.Header().Set("Content-Type", "application/json")

	// HTTP Response Codes
	w.WriteHeader(http.StatusOK)

	// If our handler function is meant to return HTML, we can actually use the same syntax to do so!
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")

	json.NewEncoder(w).Encode(dictionary)
}

func main() {
	http.HandleFunc("/", getdictionary)

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", nil)
}
