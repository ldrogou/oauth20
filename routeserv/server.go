package routeserv

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ldrogou/goauth20/middleware"
	"github.com/ldrogou/goauth20/store"
)

type Server struct {
	Router *mux.Router
	Store  store.Store
}

//File structure du fichier
type File struct {
	JwtID      int64
	JwtProduce string
	Header     string
	Payload    string
	Sign       string
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	middleware.LogRequestMiddleware(s.Router.ServeHTTP).ServeHTTP(rw, r)
}

func (s *Server) response(rw http.ResponseWriter, _ *http.Request, data interface{}, status int) {
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

func (s *Server) decode(rw http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)

}
