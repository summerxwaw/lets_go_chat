package middleware

import (
	"log"
	"net/http"
)

type HttpError struct {
	statusCode uint16
	body       string
}

type handlerError func(w http.ResponseWriter, r *http.Request) *HttpError

func ErrorMiddleware(next handlerError) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			log.Printf(err.body, "statusCode", err.statusCode)

			w.WriteHeader(int(err.statusCode))
		}
	})
}
