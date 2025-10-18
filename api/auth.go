package api

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CodeResponse struct {
	Code string
	State string
}

func LoginWithCode(authRes CodeResponse) (LoginResponse, error) {
	q := url.Values{}
	q.Add("code", authRes.Code)	
	q.Add("redirect_uri", viper.GetString("redirect_uri")) 
	q.Add("grant_type", "authorization_code")

	url := viper.GetString("spotify_account_url") + "/api/token?" + q.Encode()

	client := &http.Client{}
	data := viper.GetString("client_id" )+ ":" + viper.GetString("client_secret") 
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	authKey := "Basic " + encodedData

	req, err := http.NewRequest("POST", url, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", authKey)		

	res, err := client.Do(req)
	if err != nil {
		log.Error("not able to request auth token", "error", err)
		return LoginResponse{}, nil
	}

	defer res.Body.Close()

	var loginRes LoginResponse
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &loginRes)

	log.Debug("auth token req received", "resStruct", loginRes)

	return loginRes, nil
}

