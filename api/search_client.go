package api

import (
	"net/http"
	"time"
)

type SearchClient struct {
	HTTP  *http.Client
	args  SearchArgs
	Query map[string]string
}

func NewSearchClient() *SearchClient {
	return &SearchClient{
		HTTP: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *SearchClient) do(req *http.Request) (*http.Response, error) {
	return s.HTTP.Do(req)
}

func (s *SearchClient) prep(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+s.args.AccessToken)
}
