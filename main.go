package main

import (
	"fmt"
	"log/slog"
	"net/http"

	middleware "github.com/summerxwaw/lets_go_chat/middlewares"
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
	mux := http.NewServeMux()

	mux.Handle("/v1/user", middleware.LoggingMiddleware(http.HandlerFunc(user.HandleUserCreate)))
	mux.Handle("/v1/user/login", middleware.LoggingMiddleware(http.HandlerFunc(user.HandleUserLogin)))

	slog.Info("Server is listening on localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)

	slog.Error(fmt.Sprint(err))
}
