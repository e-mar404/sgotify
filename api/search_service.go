package api

import (
	"strings"
)

type SearchArgs struct {
	BaseURL     string
	AccessToken string
	Track       string
	Type        []string
}

type SearchReply struct {
	Tracks struct {
		Items []struct {
			Name    string `json:"name"`
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
			URI string `json:"uri"`
		} `json:"items"`
	} `json:"tracks"`
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

func (s *Search) Search(args *SearchArgs, reply *SearchReply) error {
	u := args.BaseURL + "/search"
	q := map[string]string{
		"q":    "q=track%3D" + args.Track,
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
