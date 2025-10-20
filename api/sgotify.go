package api

import "net/url"

// TODO:  
// 1. make this func work for initial login
// 2. make this a generic function 
func do(ac *AuthClient, method string, url string, q url.Values) (*LoginResponse, error) {
	// build url with query
	// create request
	// add headers to request
	// do request
	// read res
	// unmarshal res
	// return loginRes, err
}
