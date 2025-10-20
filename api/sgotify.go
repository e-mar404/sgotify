package api

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type Client interface {
	authKeySet() bool
	header() *http.Header
	addHeader(string, string)
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

	data := viper.GetString("client_id" )+ ":" + viper.GetString("client_secret") 
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	authKey := "Basic " + encodedData
	c.addHeader("Authorization", authKey)

	req.Header = *c.header()

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
	if err := json.Unmarshal(body, &resStruct); err != nil {
		log.Error("could not unmarshal response", "error", err)
		return nil, err
	}
	log.Debug("returning raw body", "body", string(body))
	log.Debug("returning unmashaled res body", "resStruct", resStruct)

	return &resStruct, nil
}
