package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ldrogou/goauth20/model"
	templateoauth "github.com/ldrogou/goauth20/templateOAuth"
)

//Claim claims to export
type Claims struct {
	Sub          string   `json:"sub"`
	IDEntreprise string   `json:"idEntreprise"`
	RcaPartnerID string   `json:"rcaPartnerId"`
	Scopes       []string `json:"scopes"`
	Roles        []string `json:"roles"`
	jwt.StandardClaims
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		rw.Header().Set("Content-Type", "text/html")
		rw.WriteHeader(http.StatusOK)

		t, err := template.New("test").Parse(templateoauth.TemplateIndex)
		if err != nil {
			fmt.Printf("erreur suivante %v", err)
		}

		err = t.Execute(rw, nil)
		if err != nil {
			fmt.Printf("erreur suivante %v", err)
		}
	}

}
func (s *server) handleLocal() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		sub := r.FormValue("sub")
		idEntreprise := r.FormValue("id_entreprise")
		rcaPartnerID := r.FormValue("rcaPartnerId")
		jwtKey := r.FormValue("secret")
		scopes := r.FormValue("scopes")
		roles := r.FormValue("roles")

		var sc []string
		sc = append(sc, scopes)

		rs := strings.Fields(roles)

		// Declare the expiration time of the token
		// here, we have kept it as 5 minutes
		expirationTime := time.Now().Add(5 * time.Hour)
		// Create the JWT claims, which includes the username and expiry time
		claims := &Claims{
			Sub:          sub,
			IDEntreprise: idEntreprise,
			RcaPartnerID: rcaPartnerID,
			Roles:        rs,
			Scopes:       sc,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		secretBase64, err := jwt.DecodeSegment(jwtKey)
		// Declare the token with the algorithm used for signing, and the claims
		ts := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

		at, err := ts.SignedString(secretBase64)
		// Create the JWT string
		if err != nil {
			log.Printf("erreur %v", err)
			// If there is an error in creating the JWT return an internal server error
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Puis redisrect vers page resultat
		o := &model.Oauth{
			ID:           0,
			AccessToken:  at,
			TokenType:    "bearer",
			ExpiresIN:    -1,
			RefreshToken: "refresh",
		}
		err = s.store.CreateOauth(o)
		if err != nil {
			fmt.Printf("erreur suivante %v", err)
		}

		monID := strconv.Itoa(int(o.ID))
		// Puis redisrect vers page resultat
		rj := "http://localhost:8090/jwt?model=" + monID
		http.Redirect(rw, r, rj, http.StatusMovedPermanently)

	}

}

func (s *server) handleOAuth20() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		d := r.FormValue("domain")
		ci := r.FormValue("clientId")
		cs := r.FormValue("clientSecret")
		sc := r.FormValue("scopes")
		cc := r.FormValue("currentCompany")
		if len(cc) == 0 {
			cc = "false"
		} else {
			cc = "true"
		}

		// Création du nombre aléatoire pour la state
		nr := rand.NewSource(time.Now().UnixNano())
		rand := rand.New(nr)
		st := strconv.Itoa(rand.Intn(10000000000))

		// Insert en base de données
		p := &model.Param{
			ID:           0,
			State:        st,
			Domaine:      d,
			ClientID:     ci,
			ClientSecret: cs,
			GrantType:    "authorization_code",
		}

		err := s.store.CreateParam(p)
		if err != nil {
			fmt.Printf("erreur suivante %v", err)
		}

		// on appelle les méthodes de l'instance de `rand.Rand` obtenue comme les autres méthodes du package.
		//fmt.Print(r1.Intn(100), ",")

		rhttp := "https://" + d + "/entreprise-partenaire/authorize?client_id=" + ci +
			"&scope=" + sc +
			"&current_company=" + cc +
			"&redirect_uri=http://localhost:8090/oauth/redirect%3Fstate=" + st +
			"&abort_uri=http://localhost:8090/index"
		http.Redirect(rw, r, rhttp, http.StatusMovedPermanently)

	}

}
