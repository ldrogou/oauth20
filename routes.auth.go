package main

import (
	"fmt"
	"net/http"
)

func (s *server) handleIndex() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Welcome to Goflix")
	}

}

func (s *server) handleRedirect() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}
