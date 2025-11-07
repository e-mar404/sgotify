package api

import (
	"github.com/charmbracelet/log"
)

func init() {
	server.Register(NewUserService())
}

type profileResponse struct {
	DisplayName string `json:"display_name"`
	Followers   struct {
		Total int `json:"total"`
	} `json:"followers"`
}

type topResponse struct {
	Items []struct {
		Name string `json:"name"`
	} `json:"items"`
}

type ProfileArgs struct {
	AccessToken string
	BaseUrl     string
}

type ProfileReply struct {
	Username  string
	Followers int
	TopArtist string
	TopTrack  string
}

type User struct {
	Client *UserClient
}

func NewUserService() *User {
	return &User{
		Client: NewUserClient(),
	}
}

func (u *User) Profile(args *ProfileArgs, reply *ProfileReply) error {
	log.Info("called User.Profile")

	u.Client.args = *args

	urlProfile := args.BaseUrl + "/me"
	urlTopArtists := args.BaseUrl + "/me/top/artists"
	urlTopTracks := args.BaseUrl + "/me/top/tracks"

	timeRange := map[string]string{
		"time_range": "short_term",
		"limit":      "1",
	}

	profileRes, err := do[profileResponse](u.Client, "GET", urlProfile, nil)
	if err != nil {
		return err
	}
	topArtists, err := do[topResponse](u.Client, "GET", urlTopArtists, timeRange)
	if err != nil {
		return err
	}
	topTracks, err := do[topResponse](u.Client, "GET", urlTopTracks, timeRange)
	if err != nil {
		return err
	}

	*reply = ProfileReply{
		Username:  profileRes.DisplayName,
		Followers: profileRes.Followers.Total,
		TopArtist: topArtists.Items[0].Name,
		TopTrack:  topTracks.Items[0].Name,
	}

	return nil
}
