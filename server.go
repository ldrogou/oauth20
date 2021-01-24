package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/ldrogou/goauth20/model"
	"github.com/ldrogou/goauth20/store"
	templateoauth "github.com/ldrogou/goauth20/templateOAuth"
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

func (s *server) responseFile(rw http.ResponseWriter, _ *http.Request, data interface{}, status int) error {
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(status)

	tokenVal := data.(string)

	tableau := strings.Split(tokenVal, ".")
	header, err := jwt.DecodeSegment(tableau[0])
	if err != nil {
		return fmt.Errorf("Impossible de décoder le header. (err=%v)", err)
	}
	payload, err := jwt.DecodeSegment(tableau[1])
	if err != nil {
		return fmt.Errorf("Impossible de décoder le payload. (err=%v)", err)
	}

	//t := template.New("mon template")
	t, err := template.New("Resulta").Parse(templateoauth.Resultat)
	if err != nil {
		return fmt.Errorf("erreur suivante %v", err)
	}

	f := File{
		JwtProduce: tokenVal,
		Header:     string(header),
		Payload:    string(payload),
		Sign:       tableau[2],
	}

	o := &model.Oauth{
		ID:           0,
		AccessToken:  tokenVal,
		ExpireIN:     180,
		RefreshToken: "eeeee",
	}
	err = s.store.CreateOauth(o)
	if err != nil {
		fmt.Printf("erreur suivante %v", err)
	}

	log.Println("ezdzedezd")
	err = t.Execute(rw, f)
	if err != nil {
		return fmt.Errorf("erreur suivante %v", err)
	}

	return nil
}

func (s *server) decode(rw http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)

}
