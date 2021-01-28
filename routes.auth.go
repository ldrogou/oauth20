package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ldrogou/goauth20/model"
	templateoauth "github.com/ldrogou/goauth20/templateOAuth"
)

//JSONToken json token
type JSONToken struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RedirectURI  string `json:"redirect_uri"`
	Code         string `json:"code"`
}

//Token token
type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

//Claim claims to export
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
		var jwtKey = []byte(r.FormValue("secret"))

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

		// Declare the token with the algorithm used for signing, and the claims
		ts := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Create the JWT string
		at, err := ts.SignedString(jwtKey)
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
		rj := "http://localhost:8080/jwt?model=" + monID
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
			"&redirect_uri=http://localhost:8080/oauth/redirect%3Fstate=" + st +
			"&abort_uri=http://localhost:8080/index"
		http.Redirect(rw, r, rhttp, http.StatusMovedPermanently)

	}

}

func (s *server) handleRedirect() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		c := r.URL.Query().Get("code")
		st := r.URL.Query().Get("state")

		// ici jouter la récupération du param
		p, err := s.store.GetParam(st)
		if err != nil {
			fmt.Printf("erreur à la recupération des param (err=%v)", err)
		}
		jsonStr := constJSONToken(c, st, p)
		//log.Printf("jsonStr %v", jsonStr)
		apiURL := "https://api." + p.Domaine + "/auth/v1/oauth2.0/accessToken"
		data := url.Values{}
		log.Printf("data %v", data)
		data.Set("client_id", jsonStr.ClientID)
		data.Set("client_secret", jsonStr.ClientSecret)
		//"YNVZF88dD4vny59k")
		data.Set("grant_type", jsonStr.GrantType)
		data.Set("redirect_uri", jsonStr.RedirectURI)
		data.Set("code", jsonStr.Code)

		client := &http.Client{}
		req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
		if err != nil {
			log.Printf("erreur sur le post (err=%v)", err)
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("client erreur %v", err)
		}

		log.Printf("resp status %v", resp.StatusCode)
		var t map[string]interface{}
		// here's the trick
		json.NewDecoder(resp.Body).Decode(&t)

		if err != nil {
			log.Printf("Cannot parse token body err=%v", err)
			s.response(rw, r, nil, http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		if err != nil {
			log.Printf("Cannot parse token body err=%v", err)
			s.response(rw, r, nil, http.StatusBadGateway)
			return
		}
		// Insert en base de données
		o := &model.Oauth{
			ID:           0,
			AccessToken:  t["access_token"].(string),
			TokenType:    t["token_type"].(string),
			ExpiresIN:    t["expires_in"].(float64),
			RefreshToken: t["refresh_token"].(string),
		}
		err = s.store.CreateOauth(o)
		if err != nil {
			fmt.Printf("erreur suivante %v", err)
		}

		monID := strconv.Itoa(int(o.ID))
		// Puis redisrect vers page resultat
		rj := "http://localhost:8080/jwt?model=" + monID
		http.Redirect(rw, r, rj, http.StatusMovedPermanently)
	}
}

func (s *server) handleJSONWebToken() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		c := r.URL.Query().Get("model")

		rw.Header().Set("Content-Type", "text/html")
		rw.WriteHeader(http.StatusOK)

		t, err := template.New("test").Parse(templateoauth.TemplateIndex)
		if err != nil {
			fmt.Printf("erreur suivante %v", err)
		}

		oauthID, err := strconv.ParseInt(c, 10, 64)

		oauth, err := s.store.GetOauth(oauthID)
		if err != nil {
			log.Printf("erreur a la récupération oauth (err=%v)", err)
		}
		tokenVal := oauth.AccessToken

		fmt.Println("============")
		fmt.Println(tokenVal)
		fmt.Println("============")

		tableau := strings.Split(tokenVal, ".")
		header, err := jwt.DecodeSegment(tableau[0])
		if err != nil {
			fmt.Printf("Impossible de décoder le header. (err=%v)", err)
		}
		payload, err := jwt.DecodeSegment(tableau[1])
		if err != nil {
			fmt.Printf("Impossible de décoder le payload. (err=%v)", err)
		}

		//t := template.New("mon template")
		t, err = template.New("Resultat").Parse(templateoauth.Resultat)
		if err != nil {
			fmt.Printf("erreur suivante %v", err)
		}

		f := File{
			JwtProduce: tokenVal,
			Header:     string(header),
			Payload:    string(payload),
			Sign:       tableau[2],
		}

		err = t.Execute(rw, f)
		if err != nil {
			fmt.Printf("erreur suivante %v", err)
		}
	}
}

func constJSONToken(code, state string, param *model.Param) JSONToken {
	return JSONToken{
		ClientID:     param.ClientID,
		ClientSecret: param.ClientSecret,
		GrantType:    param.GrantType,
		RedirectURI:  "http://localhost:8080/oauth/redirect?state=" + state,
		Code:         code,
	}
}
