package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/charmbracelet/log"
)

type Client interface {
	prep(*http.Request)
	do(*http.Request) (*http.Response, error)
}

// Takes in a client and will Unmarshal the response that it gets into struct T
func do[T any](c Client, method string, urlPath string, q map[string]string) (*T, error) {
	query := url.Values{}
	for key, value := range q {
		query.Add(key, value)
	}
	 
	fullUrl := urlPath + "?" + query.Encode()
	log.Debug("created full url", "fullUrl", fullUrl)

	req, err := http.NewRequest(method, fullUrl, nil)
	if err != nil {
		log.Error("could not create request", "method", method, "fullUrl", fullUrl)
		return nil, err
	}
	c.prep(req)

	res, err := c.do(req)
	if err != nil {
		log.Error("could not complete request", "method", method, "fullUrl", fullUrl)
		return nil, err
	}
	log.Debug("completed request", "res", res)

	// TODO: check status code here

	defer res.Body.Close()	

	var resStruct T 
	body, _ := io.ReadAll(res.Body)
	log.Debug("raw res body", "body", string(body))
	if err := json.Unmarshal(body, &resStruct); err != nil {
		log.Error("could not unmarshal response", "error", err)
		return nil, err
	}
	log.Debug("returning unmashaled res body", "resStruct", resStruct)

	return &resStruct, nil
}
