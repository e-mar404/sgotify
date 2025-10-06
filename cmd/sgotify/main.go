package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	clientID string
	clientSecret string
	redirectURI string 
	authURL string
	tokenURL string
}

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

	app := App {
		cfg: Config{
			clientID: clientID,
			clientSecret: clientSecret,
			redirectURI: "http://127.0.0.1:8080/callback",
			authURL: "https://accounts.spotify.com/authorize",
			tokenURL: "https://accounts.spotify.com/api/token",
		},
	}
	
	app.Start()
}
