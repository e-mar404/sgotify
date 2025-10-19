package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

// TODO: still needs an api call for this
type TopItemsResponse struct {
	Items []struct {
		Name string `json:"name"`	
	} `json:"items"`
}

type ProfileResponse struct {
	DisplayName string `json:"display_name"`
	Followers struct {
		Total int `json:"total"`
	} `json:"followers"`
}

func UserProfile() (*ProfileResponse, error) {
	client := &http.Client{}
	apiUrl := viper.GetString("spotify_api_url") + "/me"

	log.Debug("creating request", "url", apiUrl)
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Error("could not create request", "error", err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer " + viper.GetString("access_token"))
	
	res, err := client.Do(req)
	if err != nil {
		log.Error("could not get response", "error", err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("could not read response body", "error", err)
		return nil, err
	}
	log.Debug("profile response body", "content", string(body))

	var profile ProfileResponse
	if err := json.Unmarshal(body, &profile); err != nil {
		log.Error("could not unmarshal response", "error", err)
		return nil, err 
	}
	
	return &profile, nil
}

func TopArtist() (*TopItemsResponse, error) {
	q := url.Values{}

	q.Add("time_range", "short_term")
	q.Add("limit", "1")
	q.Add("offset", "0")

	client := &http.Client{}
	apiUrl := viper.GetString("spotify_api_url") + "/me/top/artists?" + q.Encode()

	log.Debug("creating request", "url", apiUrl)
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Error("could not create request", "error", err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer " + viper.GetString("access_token"))
	
	res, err := client.Do(req)
	if err != nil {
		log.Error("could not get response", "error", err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("could not read response body", "error", err)
		return nil, err
	}
	log.Debug("profile response body", "content", string(body))

	var topArtist TopItemsResponse
	if err := json.Unmarshal(body, &topArtist); err != nil {
		log.Error("could not unmarshal response", "error", err)
		return nil, err 
	}
	
	return &topArtist, nil
}

func TopTrack() (*TopItemsResponse, error) {
	q := url.Values{}

	q.Add("time_range", "short_term")
	q.Add("limit", "1")
	q.Add("offset", "0")

	client := &http.Client{}
	apiUrl := viper.GetString("spotify_api_url") + "/me/top/tracks?" + q.Encode()

	log.Debug("creating request", "url", apiUrl)
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Error("could not create request", "error", err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer " + viper.GetString("access_token"))
	
	res, err := client.Do(req)
	if err != nil {
		log.Error("could not get response", "error", err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("could not read response body", "error", err)
		return nil, err
	}
	log.Debug("profile response body", "content", string(body))

	var topTrack TopItemsResponse
	if err := json.Unmarshal(body, &topTrack); err != nil {
		log.Error("could not unmarshal response", "error", err)
		return nil, err 
	}
	
	return &topTrack, nil
}
