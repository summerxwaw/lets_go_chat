package main

import (
	"fmt"
	"net/http"

	"github.com/summerxwaw/lets_go_chat/user"
)

var store = make(map[string]string)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		key := r.URL.Path[1:]
		var value string

		fmt.Println(r.Body, "%s", &value)

		store[key] = value
	case "GET":
		fmt.Fprintf(w, store[r.URL.Path[1:]])
	}
}

func main() {
	http.HandleFunc("/v1/user", user.HandleUserCreate)
	http.HandleFunc("/v1/user/login", user.HandleUserLogin)

	fmt.Println("Server is listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
