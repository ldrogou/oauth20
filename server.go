package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  Store
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

func (s *server) responseFile(rw http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(status)

	tokenVal := data.(interface{}).(map[string]interface{})

	//t := template.New("mon template")
	tem, err := template.ParseFiles("template/resultat.html")
	if err != nil {
		fmt.Errorf("erreur suivante %v", err)
	}

	sssss := tokenVal["access_token"].(string)
	header := tokenVal["header"].(string)
	payload := tokenVal["payload"].(string)
	//sssss := "erer"
	log.Println(sssss)

	f := File{
		JwtProduce: sssss,
		Header:     header,
		Payload:    payload,
	}

	err = tem.Execute(rw, f)
	if err != nil {
		fmt.Errorf("erreur suivante %v", err)
	}

}

func (s *server) decode(rw http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)

}
