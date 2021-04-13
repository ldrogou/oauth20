package routeserv

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) handleRefreshToken() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		vars, _ := mux.Vars(r)["id"]
		jwtID, err := strconv.ParseInt(vars, 10, 64)
		if err != nil {
			log.Printf("erreur a la récupération id jwt (err=%v)", err)
		}

		fmt.Printf("le jwtID : %v", jwtID)

		// Puis redisrect vers page resultat
		//s.response(rw, r, resp, http.StatusOK)
	}
}
