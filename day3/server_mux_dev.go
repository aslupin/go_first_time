package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var people Person

func getNums(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["id"])
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func createPeople(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&people)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people", createPeople).Methods("POST")
	router.HandleFunc("/nums/{id}", getNums).Methods("GET")
	http.ListenAndServe(":8080", router)
}
