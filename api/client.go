package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"regexp"

	"github.com/charmbracelet/log"
)

type Client interface {
	prep(*http.Request)
	do(*http.Request) (*http.Response, error)
}

// Takes in a client and will Unmarshal the response that it gets into struct T
func do[T any](c Client, method string, urlPath string, q map[string]string, body io.Reader) (reply *T, err error) {
	query := url.Values{}
	for key, value := range q {
		query.Add(key, value)
	}

	fullUrl := urlPath + "?" + query.Encode()
	log.Debug("created full url", "fullUrl", fullUrl)

	req, err := http.NewRequest(method, fullUrl, body)
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
	log.Debug("completed request", "status", res.Status)

	// TODO: check status code here

	defer res.Body.Close()

	rawBody, _ := io.ReadAll(res.Body)
	log.Debug("raw res body", "body", string(rawBody))

	// The player endpoints should return 204 with no content but instead it
	// returns 200 with random response when sending an empty body, did not find a
	// fix or any mention of it from spotify
	noContentEndpoint, _ := regexp.Compile("/me/player/[play|pause|next]")

	if noContentEndpoint.Match([]byte(fullUrl)) &&
		res.StatusCode == http.StatusOK ||
		res.StatusCode == http.StatusNoContent {
		log.Info("endpoint does not have meaningful return")
		return reply, nil
	}

	if err := json.Unmarshal(rawBody, &reply); err != nil {
		log.Error("could not unmarshal response", "error", err)
		return nil, err
	}

	log.Debug("returning unmashaled res body", "resStruct", *reply)
	return reply, nil
}
