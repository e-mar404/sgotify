package api

// import (
//
// 	"github.com/charmbracelet/log"
// 	"github.com/spf13/viper"
// )
//
// type TopItemsResponse struct {
// 	Items []struct {
// 		Name string `json:"name"`
// 	} `json:"items"`
// }
//
// type ProfileResponse struct {
// 	DisplayName string `json:"display_name"`
// 	Followers   struct {
// 		Total int `json:"total"`
// 	} `json:"followers"`
// }
//
//
// func (uc *UserClient) UserProfile() (*ProfileResponse, error) {
// 	apiUrl := viper.GetString("spotify_api_url") + "/me"
// 	profileRes, err := do[ProfileResponse](uc, "GET", apiUrl, nil)
// 	if err != nil {
// 		log.Error("could not complete user profile req", "error", err)
// 		return nil, err
// 	}
// 	return profileRes, nil
// }
//
// func (uc *UserClient) TopArtist() (*TopItemsResponse, error) {
// 	q := map[string]string{
// 		"time_range": "short_term",
// 		"limit":      "1",
// 		"offset":     "0",
// 	}
// 	apiUrl := viper.GetString("spotify_api_url") + "/me/top/artists"
// 	topArtistRes, err := do[TopItemsResponse](uc, "GET", apiUrl, q)
// 	if err != nil {
// 		log.Error("could not complete top artist request", "error", err)
// 		return nil, err
// 	}
// 	return topArtistRes, nil
// }
//
// func (uc *UserClient) TopTrack() (*TopItemsResponse, error) {
// 	q := map[string]string{
// 		"time_range": "short_term",
// 		"limit":      "1",
// 		"offset":     "0",
// 	}
// 	apiUrl := viper.GetString("spotify_api_url") + "/me/top/tracks"
// 	topTrackRes, err := do[TopItemsResponse](uc, "GET", apiUrl, q)
// 	if err != nil {
// 		log.Error("could not complete top track request", "error", err)
// 		return nil, err
// 	}
// 	return topTrackRes, nil
// }
