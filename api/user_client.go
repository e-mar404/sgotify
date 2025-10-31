package api

import (
	"net/http"

	"github.com/charmbracelet/log"
)

type UserClient struct {
	HTTP     *http.Client
	args ProfileArgs 
	Query    map[string]string
}

func (u *UserClient) prep(req *http.Request) {
	log.Debug("prepping user client", "access_token", u.args.AccessToken)
	req.Header.Add("Authorization", "Bearer "+u.args.AccessToken)
}

func (u *UserClient) do(req *http.Request) (*http.Response, error) {
	return u.HTTP.Do(req)
}

func NewUserClient() *UserClient {
	return &UserClient{
		HTTP: &http.Client{},
	}
}
