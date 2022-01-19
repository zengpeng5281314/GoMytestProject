package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Fprintf(w, "%s  %s  %d years old!", user.FirstName, user.LastName, user.Age)
	})

	http.HandleFunc("/encode", func(writer http.ResponseWriter, request *http.Request) {
		peter := User{
			FirstName: "John",
			LastName:  "Doe",
			Age:       25,
		}

		json.NewEncoder(writer).Encode(peter)
	})

	http.ListenAndServe(":80", nil)
}
