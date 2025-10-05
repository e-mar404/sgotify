package main

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

const redirecURI = "http://127.0.0.1:8080/callback"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := url.Values{}
		query.Add("response_type", "code")
		query.Add("client_id", clientID)
		query.Add("redirect_uri", redirecURI)
		query.Add("scope", "streaming user-read-private user-read-email")
		query.Add("state", "1234567890")

		url := fmt.Sprintf("%s?%s", authURL, query.Encode())
		log.Printf("redirecting from / to %v\n", url)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	})

	// TODO: need to figure out how to do the call to get the auth token
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")
		log.Printf("[/callback] code: %s, state: %s\n", code, state)
		
		// body := struct {
		// 	GrantType string `json:"grant_type"`
		// 	Code string `json:"code"`
		// 	RedirectURI string `json:"redirect_uri"`
		// }{
		// 	GrantType: "authorization_code",
		// 	Code: code,
		// 	RedirectURI: redirecURI,
		// }
		//
		// jsonData, _ := json.Marshal(body)
		//
		// req, _ := http.NewRequest("POST", tokenURL, bytes.NewReader(jsonData))
		// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		// req.Header.Add("Authorization", "")
	})

	log.Println("listening on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
