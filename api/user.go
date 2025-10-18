package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type ProfileResponse struct {
	DisplayName string `json:"display_name"`
	Email string `json:"email"`
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
