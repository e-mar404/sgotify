package command

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"net/url"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/internal/config"
	constants "github.com/e-mar404/sgotify/internal/const"
)

func Login(state *constants.State) error {
	log.Info("Checking for client id & client secret")
	
	if state.Cfg.ClientSecret == "" || state.Cfg.ClientID == "" {
		log.Info("Configuration file does not contain necessary fields") 

		var cid, cs string
		fmt.Printf("What is your ClientID?\n")
		fmt.Print("> ")
		fmt.Scan(&cid)

		fmt.Printf("\nWhat is your ClientSecret?\n")
		fmt.Print("> ")
		fmt.Scan(&cs)

		state.Cfg.ClientID = cid
		state.Cfg.ClientSecret = cs

		err := state.Cfg.Save()

		if err != nil {
			log.Error("could not save new client configuration", "error", err)
			return err
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		stateCode := rand.Text()
		scope := "user-read-private user-read-email"

		q := url.Values{}
		q.Add("response_type", "code")
		q.Add("client_id", state.Cfg.ClientID)
		q.Add("scope", scope)
		q.Add("redirect_uri", config.RedirectURI)
		q.Add("state", stateCode)

		url := config.RedirectURI + "?" + q.Encode()
		
		log.Info("redirecting to spotify auth page", "redirect url", url)

		http.Redirect(w, r, url, http.StatusPermanentRedirect)
	})

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("doing callback stuff"))
		r.Close = true
	})

	log.Info("Starting auth server", "port", ":8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Error("Problem starting auth server", "error", err)
	}

	return nil
}

