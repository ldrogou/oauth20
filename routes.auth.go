package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//File structure du fichier
type File struct {
	Name  string
	Other string
}

type JsonToken struct {
	clientID     string `json:"client_id"`
	clientSecret string `json:"client_secret"`
	grantType    string `json:"grant_type"`
	redirectURI  string `json:"redirect_uri"`
	code         string `json:"code"`
}

type token struct {
	accessToken  string `json:"access_token"`
	tokenType    string `json:"token_type"`
	expiresIn    int    `json:"expires_in"`
	refreshToken string `json:"refresh_token"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Sub          string   `json:"sub"`
	IDEntreprise string   `json:"idEntreprise"`
	RcaPartnerID string   `json:"rcaPartnerId"`
	Roles        []string `json:"roles"`
	jwt.StandardClaims
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html")
		rw.WriteHeader(http.StatusOK)

		t, err := template.ParseFiles("template/jwt.html")
		if err != nil {
			fmt.Errorf("erreur suivante %v", err)
		}

		f := File{Name: "Drogou", Other: "Dans le fichier"}

		err = t.Execute(rw, f)
		if err != nil {
			fmt.Errorf("erreur suivante %v", err)
		}
	}

}
func (s *server) handleTest() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		fmt.Println("sub")

		sub := r.FormValue("sub")
		fmt.Printf("sub %v", sub)
		idEntreprise := r.FormValue("id_entreprise")
		fmt.Printf("idEntreprise %v", idEntreprise)
		rcaPartnerID := r.FormValue("rcaPartnerId")
		fmt.Printf("rcaPartnerID %v", rcaPartnerID)
		var jwtKey = []byte(r.FormValue("secret"))
		fmt.Printf("secret %v", jwtKey)

		// Declare the expiration time of the token
		// here, we have kept it as 5 minutes
		expirationTime := time.Now().Add(5 * time.Hour)
		roles := []string{"RCA_CLOUD_EXPERT_COMPTABLE",
			"E_COLLECTE_BO_CREA",
			"E_CREATION_CREA",
			"E_QUESTIONNAIRE_CREA"}
		// Create the JWT claims, which includes the username and expiry time
		claims := &Claims{
			Sub:          sub,
			IDEntreprise: idEntreprise,
			RcaPartnerID: rcaPartnerID,
			Roles:        roles,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}
		fmt.Printf("claims %v", claims)

		// Declare the token with the algorithm used for signing, and the claims
		tokenstr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		fmt.Printf("token %v", tokenstr)

		// Create the JWT string
		tokenString, err := tokenstr.SignedString(jwtKey)
		fmt.Printf("tokenString %v", tokenString)
		if err != nil {
			log.Printf("erreur %v", err)
			// If there is an error in creating the JWT return an internal server error
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		s.response(rw, r, tokenString, http.StatusOK)
	}

}
func (s *server) handleRedirect() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		codes, _ := r.URL.Query()["code"]
		jsonStr := constJsonToken(codes[0])

		apiURL := "https://api.captation.beta.rca.fr/auth/v1/oauth2.0/accessToken"
		data := url.Values{}
		data.Set("client_id", jsonStr.clientID)
		data.Set("client_secret", jsonStr.clientSecret)
		data.Set("grant_type", jsonStr.grantType)
		data.Set("redirect_uri", jsonStr.redirectURI)
		data.Set("code", jsonStr.code)

		client := &http.Client{}
		req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		var t interface{}
		// here's the trick
		json.NewDecoder(resp.Body).Decode(&t)

		if err != nil {
			log.Printf("Cannot parse token body err=%v", err)
			s.response(rw, r, nil, http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		tokenVal := t.(interface{}).(map[string]interface{})

		if err != nil {
			log.Printf("Cannot parse token body err=%v", err)
			s.response(rw, r, nil, http.StatusBadGateway)
			return
		}

		s.response(rw, r, tokenVal["access_token"], http.StatusOK)

	}
}

func constJsonToken(code string) JsonToken {
	return JsonToken{
		clientID:     "meg-test-interne",
		clientSecret: "YNVZF88dD4vny59k",
		grantType:    "authorization_code",
		redirectURI:  "http://localhost:8080/callback",
		code:         code,
	}
}
