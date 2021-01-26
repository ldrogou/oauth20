package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ldrogou/goauth20/store"
)

type server struct {
	router *mux.Router
	store  store.Store
}

//File structure du fichier
type File struct {
	JwtProduce string
	Header     string
	Payload    string
	Sign       string
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *server) serveHTTP(rw http.ResponseWriter, r *http.Request) {
	logRequestMiddleware(s.router.ServeHTTP).ServeHTTP(rw, r)
}

func (s *server) response(rw http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	rw.Header().Add("Content-type", "application/json")
	rw.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		log.Printf("Cannot encode to json (err=%v)\n", err)
	}

}

func (s *server) decode(rw http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)

}
