package cmd

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type authRes struct {
	code string
	state string
}

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
				cobra.CheckErr(err)
			}

			log.Debug("client id and secret saved to config", "ClientID", cid, "ClientSecret", cs)
		}

		resChan := make(chan authRes)
		go func() {
			startHTTPServer(resChan)
		}()

		authRes := <- resChan

		log.Debug("response from spotify auth", "res", authRes) 
		
		// call api.LoginWithCode(code, state) and return creds
		// creds {
		// 	auth_token
		// 	refresh_token
		// }

		return nil
	},
}

func prompt(q string, a *string) {
	// could use bufio.NewReader(os.Stdin) here but i think it should be fine with fmt.Scan since it is only asking for values that will be a copy paste from the spotify dashboard
	fmt.Printf("%s\n> ", q)
	fmt.Scan(a)
}

func startHTTPServer(resChan chan authRes) {
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
		
		res := authRes {
			code: c,
			state: s,
		}

		resChan <- res 

		w.Write([]byte("You can close the tab now"))
	})

	http.Handle("/", redirectHandler)
	http.Handle("/callback", callbackHandler)

	serverUrl := "127.0.0.1:8080"
	log.Info("Starting server", "url", serverUrl)
	if err := http.ListenAndServe(serverUrl, nil); err != nil {
		log.Fatal("something went wrong with the http server", "error", err)
	}
}

// 	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
// 			q := url.Values{}
// 			q.Add("code", c)	
// 			q.Add("redirect_uri", config.RedirectURI) 
// 			q.Add("grant_type", "authorization_code")
//
// 			url := config.TokenURL + "?" + q.Encode()
//
// 			client := &http.Client{}
// 			data := state.Cfg.ClientID + ":" + state.Cfg.ClientSecret
// 			encodedData := base64.StdEncoding.EncodeToString([]byte(data))
// 			authKey := "Basic " + encodedData
//
// 			req, err := http.NewRequest("POST", url, nil)
// 			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 			req.Header.Add("Authorization", authKey)		
//
// 			res, err := client.Do(req)
// 			if err != nil {
// 				w.Write([]byte("could not properly get auth token"))
// 			}
//
// 			defer res.Body.Close()
//
// 			resStruct := struct {
// 				AccessToken string `json:"access_token"`
// 				RefreshToken string `json:"refresh_token"`
// 				Scope string `json:"scope"`
// 			}{}
//
// 			body, _ := io.ReadAll(res.Body)
// 			json.Unmarshal(body, &resStruct)
//
// 			log.Info("auth token req received", "resStruct", resStruct)
//
// 			state.Cfg.AuthToken = resStruct.AccessToken
// 			state.Cfg.RefreshToken = resStruct.RefreshToken
//
// 			if err := state.Cfg.Save(); err != nil {
// 				log.Error("could not save updated cfg", "error", err)
// 			}
// 		}
//
// 		r.Close = true
// 	})
//
// 	url := "http://127.0.0.1:8080/"
// 	log.Info("Starting auth server", "port", ":8080")
// 	cmd := exec.Command("xdg-open", url)
// 	err := cmd.Run()
// 	log.Info("if website does not open automatically you can go to " + url + " on your browser", "error", err)
//
// 	if err := http.ListenAndServe(":8080", mux); err != nil {
// 		log.Error("Problem starting auth server", "error", err)
// 	}
//
// 	return nil
// }
//
