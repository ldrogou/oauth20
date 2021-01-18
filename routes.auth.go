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

func (s *server) handleRedirect() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		codes, _ := r.URL.Query()["code"]
		jsonStr := constJsonToken(codes[0])

		apiURL := "https://api.XXX.XXX.XXX/auth/v1/oauth2.0/accessToken"
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

		fmt.Println(t.(interface{}).(map[string]interface{})["access_token"])

		if err != nil {
			log.Printf("Cannot parse token body err=%v", err)
			s.response(rw, r, nil, http.StatusBadGateway)
			return
		}

		s.response(rw, r, t, http.StatusOK)

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
