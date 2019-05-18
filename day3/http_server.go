package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

var user User

func mainHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		encoder := json.NewEncoder(w)
		encoder.Encode(user)

	} else if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&user)
		fmt.Fprintln(w, "OK")
	} else {
		fmt.Fprintln(w, "unknow method")
	}

}
func main() {
	http.HandleFunc("/", mainHandle)
	http.ListenAndServe(":8000", nil)

}
