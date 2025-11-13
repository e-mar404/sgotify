package api

import (
	"encoding/base64"
	"net/http"
	"time"
)

type authPrepArgs struct {
	ClientID     string
	ClientSecret string
}

type AuthClient struct {
	HTTP     *http.Client
	prepArgs authPrepArgs
	query    map[string]string
}

func NewAuthClient() *AuthClient {
	return &AuthClient{
		HTTP: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (ac *AuthClient) do(req *http.Request) (*http.Response, error) {
	res, err := ac.HTTP.Do(req)
	return res, err
}

func (ac *AuthClient) prep(req *http.Request) {
	data := ac.prepArgs.ClientID + ":" + ac.prepArgs.ClientSecret
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	authKey := "Basic " + encodedData

	req.Header.Add("Authorization", authKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}
