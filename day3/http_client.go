package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rs struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func f(c chan int) {

	var user string

	// fmt.Scanf("%v", &user)
	user = "aslupin"
	resp, _ := http.Get("https://api.github.com/users/" + user)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode, resp.Status)
	fmt.Println(string(body))
	data := Rs{}
	json.Unmarshal(body, &data)
	fmt.Println(data)
	c <- 0
}
func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go f(c)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}
