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

type Code struct {
	Code  string
	State string
}

var (
	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "start login process",
		Run: func(cmd *cobra.Command, args []string) {
			cid := viper.GetString("client_id")
			cs := viper.GetString("client_secret")

			log.Info("checking if creds exist",
				"cid", cid,
				"cs", cs,
			)

			credsExist := cid != "" && cs != ""
			useSavedCreds := "n"
			if credsExist {
				prompt("Use saved creds? [Y|n] ", &useSavedCreds, "Y")
			}

			switch strings.ToLower(useSavedCreds) {
			case "y":
				break
			case "n":
				// if they dont get back to 0 value then it will use the old creds if just
				// pressing enter with Scanln
				cid = ""
				cs = ""
				prompt("Client ID: ", &cid, "")
				prompt("Client Secret: ", &cs, "")

				viper.Set("client_id", cid)
				viper.Set("client_secret", cs)
				if err := viper.WriteConfig(); err != nil {
					log.Fatal("could not write to config file", "error", err)
				}
			default:
				prompt("Not a valid answer. Use saved creds? [Y|n] \n", &useSavedCreds, "Y")
			}

			codeChan := make(chan Code)
			go func() {
				startHTTPServer(codeChan)
			}()
			code := <-codeChan

			log.Info("received server response", "code", code)

			loginArgs := &api.LoginArgs{
				ClientID:     cid,
				ClientSecret: cs,
				RedirectURI:  viper.GetString("redirect_uri"),
				Code:         code.Code,
				State:        code.State,
			}
			reply := &api.CredentialsReply{}
			if err := client.Call("Auth.LoginWithCode", loginArgs, reply); err != nil {
				log.Error("unable to log in with code", "error", err)
			}
			log.Info("reply from authService.LoginWithcode", "reply", reply)

			viper.Set("access_token", reply.AccessToken)
			viper.Set("refresh_token", reply.RefreshToken)
			viper.Set("last_refresh", time.Now().Unix())

			fmt.Printf("Saving configuration...\n")
			viper.WriteConfig()
			fmt.Printf("Congratulations. Login successful!\n")
		},
	}
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

func prompt(q string, ans *string, defaultAns string) {
	fmt.Printf("%s", q)
	fmt.Scanln(ans)

	// if ans is empty then user chose default answer "Y"
	if *ans == "" {
		*ans = defaultAns
	}
}

func startHTTPServer(resChan chan Code) {
	redirectHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	serverPort := "8080"
	log.Info(
		"staring server",
		"host", serverHost,
		"port", serverPort,
	)

	fmt.Printf("Please go to %s:%s on your browser\n", serverHost, serverPort)
	if err := http.ListenAndServe(serverHost+":"+serverPort, nil); err != nil {
		log.Fatal("something unexpected happened with the server", "error", err)
	}
}
