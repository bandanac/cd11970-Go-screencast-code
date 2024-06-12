package main

import (
	"encoding/json"
	"fmt"

	// "io/ioutil" - deprecated
	"io"
	"net/http"

	// Import the third-party gorilla/mux package
	"github.com/gorilla/mux"
)

var dictionary = map[string]string{
	"Go":     "A programming language created by Google.",
	"Gopher": "A software engineer who builds with Go.",
	"Golang": "Another name for Go.",
}

func getDictionary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dictionary)
}

func create(w http.ResponseWriter, r *http.Request) {
	// 1. set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// 2. keep track of new entry
	var newEntry map[string]string

	// 3. read the request
	// ioutil.ReadAll(r.Body) - deprecated
	requestBody, _ := io.ReadAll(r.Body)

	// 4. parse JSON body
	json.Unmarshal(requestBody, &newEntry)

	// 5. add new entry to dictionary
	// - response with conflict if entry exists
	// - response with OK if entry does not exists
	for k, v := range newEntry {
		if _, ok := dictionary[k]; ok {
			w.WriteHeader(http.StatusConflict)
		} else {
			dictionary[k] = v
			w.WriteHeader(http.StatusCreated)
		}
	}

	// 6. return dictionary
	json.NewEncoder(w).Encode(dictionary)
}

func main() {
	// Instantiate a new router by invoking the "NewRouter" handler
	router := mux.NewRouter()

	// Rather calling http.HandleFunc, call the equivalent router.HandleFunc
	// This gives us access to method-based routing
	router.HandleFunc("/dictionary", getDictionary).Methods("GET")

	// using POST to add data into dictionary
	router.HandleFunc("/dictionary", create).Methods("POST")

	fmt.Println("Server is starting on port 3000...")

	// The second argument for "ListenAndServe" was previously nil
	/// Now that we are using our own custom router, we pass it along to "ListenAndServe" as its second argument
	http.ListenAndServe(":3000", router)
}
