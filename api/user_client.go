package api

import "net/http"

type userPrepArgs struct {
	AccessToken string
}

type UserClient struct {
	HTTP     *http.Client
	prepArgs userPrepArgs
	Query    map[string]string
}

func (u *UserClient) prep(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+u.prepArgs.AccessToken)
}

func (u *UserClient) do(req *http.Request) (*http.Response, error) {
	return u.HTTP.Do(req)
}

func NewUserClient() *UserClient {
	return &UserClient{
		HTTP: &http.Client{},
	}
}
