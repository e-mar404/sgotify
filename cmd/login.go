package cmd

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


// TODO: pretty print questions with lipgloss later
var loginCmd = &cobra.Command {
	Use: "login", 
	RunE: func(cmd *cobra.Command, args []string) error {
		var cid, cs string
		cid = viper.GetString("client_id")	
		cs = viper.GetString("client_secret")

		askForClientCreds := true 
		if cid != "" && cs != "" {
			log.Debug("using saved creds", "client_id", cid, "client_secret", cs)

			var response string
			prompt("Do you want to used the saved client id & secret [Y|n]", &response)

			switch strings.ToLower(response) {
			case "y":
				askForClientCreds = false 
			case "n":
				askForClientCreds = true
			default:
				prompt("Not valid input.\nDo you want to used the saved client id & secret [Y|n]", &response)
			}
		}

		if askForClientCreds {
			prompt("What is your ClientID?", &cid)
			prompt("What is your ClientSecret?", &cs)

			viper.Set("client_id", cid)
			viper.Set("client_secret", cs)
			if err := viper.WriteConfig(); err != nil {
				log.Error("could not save to configuration", "error", err)
				return err
			}

			log.Debug("client id and secret saved to config", "ClientID", cid, "ClientSecret", cs)
		}

		resChan := make(chan api.CodeResponse)
		go func() {
			startHTTPServer(resChan)
		}()

		authRes := <- resChan

		log.Debug("response from spotify auth", "res", authRes) 
		
		creds, err := api.LoginWithCode(authRes)
		if err != nil {
			log.Error("could not retrieve auth & refresh tokens", "error", err)
			return err
		}
		viper.Set("access_token", creds.AccessToken)
		viper.Set("refresh_token", creds.RefreshToken)
		viper.Set("last_refresh", time.Now().Unix())

		if err := viper.WriteConfig(); err != nil {
			log.Error("could not save to configuration", "error", err)
			return err
		}

		return nil
	},
}

func prompt(q string, a *string) {
	// could use bufio.NewReader(os.Stdin) here but i think it should be fine with fmt.Scan since it is only asking for values that will be a copy paste from the spotify dashboard
	fmt.Printf("%s\n> ", q)
	fmt.Scan(a)
}

func startHTTPServer(resChan chan api.CodeResponse) {
	redirectHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID := viper.GetString("client_id")
		redirecURI := viper.GetString("redirect_uri")
		stateCode := rand.Text()
		scope := "user-read-private user-read-email"

		q := url.Values{}
		q.Add("response_type", "code")
		q.Add("client_id", clientID)
		q.Add("scope", scope)
		q.Add("redirect_uri", redirecURI) 
		q.Add("state", stateCode)

		spotifyURL := viper.GetString("spotify_api_url")
		url := spotifyURL + "/authorize?" + q.Encode()

		log.Info("redirecting to spotify auth page", "redirect url", url)

		http.Redirect(w, r, url, http.StatusPermanentRedirect)
	})

	callbackHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.URL.Query().Get("code")
		s := r.URL.Query().Get("state")

		if c == "" || s == "" {
			log.Fatal("either code or state is malformed", "code", c, "state", s)
		}
		
		res := api.CodeResponse {
			Code: c,
			State: s,
		}

		resChan <- res 

		w.Write([]byte("You can close the tab now"))
	})

	http.Handle("/", redirectHandler)
	http.Handle("/callback", callbackHandler)

	serverUrl := "127.0.0.1:8080"
	log.Info("Starting server", "url", serverUrl)
	fmt.Printf("Starting server on http://%s\n", serverUrl)
	if err := http.ListenAndServe(serverUrl, nil); err != nil {
		log.Fatal("something went wrong with the http server", "error", err)
	}
}

