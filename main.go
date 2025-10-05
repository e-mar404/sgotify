package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const redirectURI = "http://127.0.0.1:8080/callback"
const authURL = "https://accounts.spotify.com/authorize"
const tokenURL = "https://accounts.spotify.com/api/token"

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	clientID, found := os.LookupEnv("CLIENT_ID")
	if !found {
		log.Fatal("client id not found after loading env file")	
	}

	clientSecret, found := os.LookupEnv("CLIENT_SECRET")
	if !found {
		log.Fatal("client secret not found after loading env file")	
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := url.Values{}
		query.Add("response_type", "code")
		query.Add("client_id", clientID)
		query.Add("redirect_uri", redirectURI)
		query.Add("scope", "streaming user-read-private user-read-email")
		query.Add("state", "1234567890")

		url := fmt.Sprintf("%s?%s", authURL, query.Encode())
		log.Printf("redirecting from / to %v\n", url)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")
		log.Printf("[/callback] code: %s, state: %s\n", code, state)

		client := &http.Client{}

		form := url.Values{
			"grant_type": []string{"authorization_code"},
			"redirect_uri": []string{redirectURI},
			"code": []string{code},
		}
		formReader := form.Encode()
		
		authEncoding := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret)) 

		req, _ := http.NewRequest("POST", tokenURL, strings.NewReader(formReader))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization", "Basic " + authEncoding)

		res, err := client.Do(req)
		if err != nil {
			log.Printf("auth token req err : %v\n", err)
			w.Write([]byte("could not complete auth token req"))
		}

		// TODO: get auth token out of res body
		defer res.Body.Close()

		log.Printf("tokenURL response: %v\n", res)
		w.Write([]byte("authtoken in logs"))
	})

	log.Println("listening on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
