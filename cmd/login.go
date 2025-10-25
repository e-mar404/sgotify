package cmd

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type Code struct {
	Code  string
	State string
}

func init() {
	availableCommands.AddCommand("login", loginHandler)
}

func loginHandler(_ command) error {
	cid := viper.GetString("client_id")
	cs := viper.GetString("client_secret")

	credsExist := cid != "" || cs != ""
	useSavedCreds := "n" // defaults to always asking for creds
	if credsExist {
		prompt("Use saved creds? [Y|n] ", &useSavedCreds)
	}

	switch strings.ToLower(useSavedCreds) {
	case "y":
		break
	case "n":
		// if they dont get back to 0 value then it will use the old creds if just
		// pressing enter with Scanln
		cid = ""
		cs = ""
		prompt("Client ID: ", &cid)
		prompt("Client Secret: ", &cs)

		viper.Set("client_id", cid)
		viper.Set("client_secret", cs)
		if err := viper.WriteConfig(); err != nil {
			log.Fatal("could not write to config file", "error", err)
		}
	default:
		prompt("Not a valid answer. Use saved creds? [Y|n] \n", &useSavedCreds)
	}

	codeChan := make(chan Code)
	go func() {
		startHTTPServer(codeChan)
	}()
	code := <-codeChan

	log.Info("received server response", "code", code)

	// get access + refresh token
	// save everything config to ~/.config/sgotigy/conf.json

	return nil
}

func prompt(q string, ans *string) {
	fmt.Printf("%s", q)
	fmt.Scanln(ans)

}

func startHTTPServer(resChan chan Code) {
	redirectHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create spotify link
		clientID := viper.GetString("client_id")
		redirectURI := viper.GetString("redirect_uri")
		state := rand.Text()
		scope := "user-read-playback-state user-modify-playback-state user-top-read user-read-private user-read-email"

		q := url.Values{}
		q.Add("response_type", "code")
		q.Add("client_id", clientID)
		q.Add("scope", scope)
		q.Add("state", state)
		q.Add("redirect_uri", redirectURI)

		spotifyURL := viper.GetString("spotify_account_url")
		redirectURL := spotifyURL + "/authorize?" + q.Encode()
		http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
	})

	callbackHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.URL.Query().Get("code")
		s := r.URL.Query().Get("state")

		if c == "" || s == "" {
			log.Fatal(
				"either code or state is malformed",
				"code", c,
				"state", s,
			)
		}

		res := Code{
			Code:  c,
			State: s,
		}

		resChan <- res

		w.Write([]byte("you can close the tab now"))
	})

	http.Handle("/", redirectHandler)
	http.Handle("/callback", callbackHandler)

	serverHost := "127.0.0.1"
	serverPort := ":8080"
	log.Info(
		"staring server",
		"host", serverHost,
		"port", serverPort,
	)

	if err := http.ListenAndServe(serverHost+serverPort, nil); err != nil {
		log.Fatal("something unexpected happened with the server", "error", err)
	}
}
