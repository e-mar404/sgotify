package api

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type AuthClient struct {
	HTTP *http.Client
	Header *http.Header
	Query map[string]string
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CodeResponse struct {
	Code string
	State string
}

func NewAuthClient() *AuthClient {
	return &AuthClient{
		HTTP: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (ac *AuthClient) do(req *http.Request) (*http.Response, error) {
	return ac.HTTP.Do(req)
}

func (ac *AuthClient) prep(req *http.Request) {
	data := viper.GetString("client_id" )+ ":" + viper.GetString("client_secret") 
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	authKey := "Basic " + encodedData

	req.Header.Add("Authorization", authKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}

func (ac *AuthClient) LoginWithCode(authRes CodeResponse) (*LoginResponse, error) {
	q := map[string]string{
		"code": authRes.Code,
		"redirect_uri": viper.GetString("redirect_uri"),
		"grant_type": "authorization_code",
	}
	url := viper.GetString("spotify_account_url") + "/api/token"
	loginRes, err := do[LoginResponse](ac, "POST", url, q)
	if err != nil {
		log.Error("could not login in with code", "error", err)
		return nil, err
	}

	return loginRes, nil
}

func (ac *AuthClient) RefreshAccessToken() (*LoginResponse, error) {
	q := map[string]string {
		"grant_type": "refresh_token",
		"refresh_token": viper.GetString("refresh_token"),
	}
	url := viper.GetString("spotify_account_url") + "/api/token"
	refreshRes, err := do[LoginResponse](ac, "POST", url, q)
	if err != nil {
		log.Error("could not refresh access token", "error", err)
		return nil, err
	}
	return refreshRes, nil
}

