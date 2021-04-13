package middleware

import (
	"log"
	"net/http"
)

func LogRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] %v", r.Method, r.RequestURI)
		next.ServeHTTP(rw, r)
	}
}
