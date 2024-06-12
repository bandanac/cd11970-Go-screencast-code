package main

import (
	"encoding/json"
	"fmt"

	// "io/ioutil" - deprecated

	"net/http"

	// Import the third-party gorilla/mux package
	"github.com/gorilla/mux"
)

var members = map[string]string{
	"1": "Andy",
	"2": "Peter",
	"3": "Gabriella",
	"4": "Jordy",
}

// Create a handler function, getMembers(), that returns the members map in JSON format as a response
func getMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(members)
}

// Create another handler function, deleteMember().
func deleteMember(w http.ResponseWriter, r *http.Request) {
	// 1. set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	if _, ok := members[id]; ok {
		delete(members, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(members)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(members)
	}
}

func main() {
	// Instantiate a new router by invoking the "NewRouter" handler
	router := mux.NewRouter()
	router.HandleFunc("/members", getMembers).Methods("GET")
	router.HandleFunc("/deleteMember/{id}", deleteMember).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
