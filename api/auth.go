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

type AuthClient struct {
	Headers http.Header
	Query url.Values
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
	data := viper.GetString("client_id" )+ ":" + viper.GetString("client_secret") 
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	authKey := "Basic " + encodedData

	return &AuthClient{
		Headers: http.Header{
			"Content-Type": []string{"application/x-www-form-urlencoded"},
			"Authorization": []string{authKey},
		},
	}
}


func (ac *AuthClient) LoginWithCode(authRes CodeResponse) (*LoginResponse, error) {
	q := url.Values{
		"code": []string{authRes.Code},
		"redirect_uri": []string{viper.GetString("redirect_uri")},
		"grant_type": []string{"authorization_code"},
	}
	url := viper.GetString("spotify_account_url") + "/api/token"
	loginRes, err := do(ac, "POST", url, q)
	if err != nil {
		log.Error("could not login in with code", "error", err)
		return nil, err
	}

	return loginRes, nil
	// client := &http.Client{}
	// res, err := client.Do(req)
	// if err != nil {
	// 	log.Error("not able to request auth token", "error", err)
	// 	return LoginResponse{}, nil
	// }
	//
	// defer res.Body.Close()
	//
	// var loginRes LoginResponse
	// body, _ := io.ReadAll(res.Body)
	// json.Unmarshal(body, &loginRes)
	//
	// log.Debug("auth token req received", "resStruct", loginRes)
	//
	// return loginRes, nil
}

func RefreshAccessToken() (*LoginResponse, error) {
	q := url.Values{}
	q.Add("grant_type", "refresh_token")
	q.Add("refresh_token", viper.GetString("refresh_token")) 
	
	url := viper.GetString("spotify_account_url") + "/api/token?" + q.Encode()
	
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Error("could not create refresh req", "error", err)
		return nil, err
	}

	data := viper.GetString("client_id" )+ ":" + viper.GetString("client_secret") 
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	authKey := "Basic " + encodedData
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", authKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Error("could not complete refresh req", "error", err)
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var refreshRes LoginResponse
	json.Unmarshal(body, &refreshRes)

	log.Debug("refresh res received", "refreshRes", refreshRes)

	return &refreshRes, nil
}

