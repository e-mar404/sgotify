package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func newAuthRouter(cfg Config) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := url.Values{}
		query.Add("response_type", "code")
		query.Add("client_id", cfg.clientID)
		query.Add("redirect_uri", cfg.redirectURI)
		query.Add("scope", "streaming user-read-private user-read-email")
		query.Add("state", "1234567890")

		url := fmt.Sprintf("%s?%s", cfg.authURL, query.Encode())
		log.Printf("redirecting from / to %v\n", url)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	})

	router.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")
		log.Printf("[/callback] code: %s, state: %s\n", code, state)

		client := &http.Client{}

		form := url.Values{
			"grant_type": []string{"authorization_code"},
			"redirect_uri": []string{cfg.redirectURI},
			"code": []string{code},
		}
		formReader := form.Encode()

		authEncoding := base64.StdEncoding.EncodeToString([]byte(cfg.clientID + ":" + cfg.clientSecret)) 

		req, _ := http.NewRequest("POST", cfg.tokenURL, strings.NewReader(formReader))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization", "Basic " + authEncoding)

		res, err := client.Do(req)
		if err != nil {
			log.Printf("auth token req err : %v\n", err)
			w.Write([]byte("could not complete auth token req"))
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			w.Write([]byte("authtoken did not complete properly"))
		}

		body, _ := io.ReadAll(res.Body)
		authTokenRes := struct{
			AccessToken	string	`json:"access_token"`
			TokenType	string	`json:"token_type"`
			Scope	string	`json:"scope"`
			ExpiresIn	int	`json:"expires_in"`
			RefreshToken	string `json:"refresh_token"`
		}{}

		err = json.Unmarshal(body, &authTokenRes)

		log.Printf("tokenURL response: %v\n", authTokenRes)
		w.Write([]byte("authtoken in logs"))
	})

	return router
}

