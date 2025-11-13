package api

import (
	"net/url"
	"strings"
)

type SearchArgs struct {
	AccessToken string
	Track       string
	Artist      string
	Album       string
	Type        []string
}

type ArtistItem struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

type TrackItem struct {
	Name    string       `json:"name"`
	Artists []ArtistItem `json:"artists"`
	URI     string       `json:"uri"`
}

type AlbumItem struct {
	Type    string       `json:"album_type"`
	Name    string       `json:"name"`
	Artists []ArtistItem `json:"artists"`
	URI     string       `json:"uri"`
}

type PlaylistItem struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

type TrackObject struct {
	Items []TrackItem `json:"items"`
}

type ArtistObject struct {
	Items []ArtistItem `json:"items"`
}

type AlbumObject struct {
	Items []AlbumItem `json:"items"`
}

type PlaylistObject struct {
	Items []PlaylistItem `json:"items"`
}

type SearchReply struct {
	Tracks   TrackObject    `json:"tracks"`
	Artists  ArtistObject   `json:"artists"`
	Albums   AlbumObject    `json:"albums"`
	Playlist PlaylistObject `json:"playlists"`
}

type Search struct {
	Client *SearchClient
}

func init() {
	server.Register(NewSearchService())
}

func NewSearchService() *Search {
	return &Search{
		Client: NewSearchClient(),
	}
}

func (s *Search) Catalog(args *SearchArgs, reply *SearchReply) error {
	u := apiBaseURL + "/search"
	// TODO: should convert the url encoding into a function or something since im
	// manually url encoding for poc
	subQuery := url.Values{}
	subQuery.Add("track", args.Track)
	subQuery.Add("artist", args.Artist)
	subQuery.Add("album", args.Album)

	q := map[string]string{
		"q":    subQuery.Encode(),
		"type": strings.Join(args.Type, ","),
	}

	s.Client.args = *args

	searchReply, err := do[SearchReply](s.Client, "GET", u, q, nil)
	if err != nil {
		return err
	}

	*reply = *searchReply

	return nil
}
