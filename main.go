package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	clientID, found := os.LookupEnv("CLIENT_ID")
	if !found {
		log.Fatal("client id not found after loading env file")	
	}

	// make this into const vars
	redirecURI := "https://127.0.0.1:8080/callback"
	authURL := "https://accounts.spotify.com/authorize"
	tokenURL := "https://accounts.spotify.com/api/token"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%v?response_type=code&client_id=%v&scope=%v&redirect_uri=%v&state=1234567890", authURL, clientID, "streaming user-read-private user-read-email", redirecURI)

		http.Redirect(w, r, url, http.StatusMovedPermanently)
	})

	// TODO: need to figure out how to do the call to get the auth token
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		// get code & state params
		code := r.URL.Query().Get("code")
		// state := r.URL.Query().Get("state")
		
		body := struct {
			GrantType string `json:"grant_type"`
			Code string `json:"code"`
			RedirectURI string `json:"redirect_uri"`
		}{
			GrantType: "authorization_code",
			Code: code,
			RedirectURI: redirecURI,
		}

		jsonData, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", tokenURL, bytes.NewReader(jsonData))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization", "")
	})

	log.Println("listening on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
