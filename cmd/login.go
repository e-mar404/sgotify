package cmd

// package command
//
// import (
// 	"crypto/rand"
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// 	"os/exec"
//
// 	"github.com/charmbracelet/log"
// 	"github.com/e-mar404/sgotify/internal/config"
// 	constants "github.com/e-mar404/sgotify/internal/const"
// )
//
// func Login(state *constants.State) error {
// 	log.Info("Checking for client id & client secret")
//
// 	if state.Cfg.ClientSecret == "" || state.Cfg.ClientID == "" {
// 		log.Info("Configuration file does not contain necessary fields") 
//
// 		var cid, cs string
// 		fmt.Printf("What is your ClientID?\n")
// 		fmt.Print("> ")
// 		fmt.Scan(&cid)
//
// 		fmt.Printf("\nWhat is your ClientSecret?\n")
// 		fmt.Print("> ")
// 		fmt.Scan(&cs)
//
// 		state.Cfg.ClientID = cid
// 		state.Cfg.ClientSecret = cs
//
// 		err := state.Cfg.Save()
//
// 		if err != nil {
// 			log.Error("could not save new client configuration", "error", err)
// 			return err
// 		}
// 	}
//
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//
// 		stateCode := rand.Text()
// 		scope := "user-read-private user-read-email"
//
// 		q := url.Values{}
// 		q.Add("response_type", "code")
// 		q.Add("client_id", state.Cfg.ClientID)
// 		q.Add("scope", scope)
// 		q.Add("redirect_uri", config.RedirectURI)
// 		q.Add("state", stateCode)
//
// 		url := config.AuthorizeURL + "?" + q.Encode()
//
// 		log.Info("redirecting to spotify auth page", "redirect url", url)
//
// 		http.Redirect(w, r, url, http.StatusPermanentRedirect)
// 	})
//
// 	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
// 		c := r.URL.Query().Get("code")
// 		s := r.URL.Query().Get("state")
//
// 		if c == "" || s == "" {
// 			w.Write([]byte("Could not properly verify account"))
// 		} else {
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
