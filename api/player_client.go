package api

import (
	"net/http"
	"time"
)

type PlayerClient struct {
	HTTP  *http.Client
	args  PlayerArgs
	Query map[string]string
}

func NewPlayerClient() *PlayerClient {
	return &PlayerClient{
		HTTP: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (p *PlayerClient) prep(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+p.args.AccessToken)
	req.Header.Add("Content-Type", "application/json")
}

func (p *PlayerClient) do(req *http.Request) (*http.Response, error) {
	return p.HTTP.Do(req)
}
