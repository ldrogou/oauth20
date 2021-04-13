package routeserv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/ldrogou/goauth20/model"
	templateoauth "github.com/ldrogou/goauth20/templateOAuth"
)

//"YNVZF88dD4vny59k")
//JSONToken json token
type JSONToken struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RedirectURI  string `json:"redirect_uri"`
	Code         string `json:"code"`
}

func (s *Server) handleRedirect() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		c := r.URL.Query().Get("code")
		st := r.URL.Query().Get("state")

		// ici jouter la récupération du param
		p, err := s.Store.GetParam(st)
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

		if resp.StatusCode != 200 {
			log.Printf("Problème dans la requete retour http %v", resp.StatusCode)
			s.response(rw, r, nil, http.StatusBadGateway)
			return
		}
		var t map[string]interface{}
		// here's the trick
		err = json.NewDecoder(resp.Body).Decode(&t)
		if err != nil {
			log.Printf("Cannot parse token body err=%v", err)
			s.response(rw, r, nil, http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// Insert en base de données
		o := &model.Oauth{
			ID:           0,
			AccessToken:  t["access_token"].(string),
			TokenType:    t["token_type"].(string),
			ExpiresIN:    t["expires_in"].(float64),
			RefreshToken: t["refresh_token"].(string),
		}
		err = s.Store.CreateOauth(o)
		if err != nil {
			fmt.Printf("erreur suivante %v", err)
		}

		monID := strconv.Itoa(int(o.ID))
		// Puis redisrect vers page resultat
		rj := "http://localhost:8090/jwt/" + monID
		http.Redirect(rw, r, rj, http.StatusMovedPermanently)
	}
}

func (s *Server) handleJSONWebToken() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		vars, _ := mux.Vars(r)["id"]
		jwtID, err := strconv.ParseInt(vars, 10, 64)
		if err != nil {
			log.Printf("erreur a la récupération id jwt (err=%v)", err)
		}

		rw.Header().Set("Content-Type", "text/html")
		rw.WriteHeader(http.StatusOK)

		t, err := template.New("test").Parse(templateoauth.TemplateIndex)
		if err != nil {
			fmt.Printf("erreur suivante %v", err)
		}

		oauth, err := s.Store.GetOauth(jwtID)
		if err != nil {
			log.Printf("erreur a la récupération oauth (err=%v)", err)
		}
		tokenVal := oauth.AccessToken

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
			JwtID:      jwtID,
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
		RedirectURI:  "http://localhost:8090/oauth/redirect%3Fstate=" + state,
		Code:         code,
	}
}
