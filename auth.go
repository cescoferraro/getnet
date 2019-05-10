package getnet

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Auth dsjkfndsfkfjng
func (api *API) Auth(ctx context.Context) (AuthFullResponse, error) {
	result := AuthFullResponse{Code: api.DefaultCode()}
	URL, err := api.getURL()
	if err != nil {
		return result, err
	}
	URL.Path += "/auth/oauth/v2/token"
	parameters := url.Values{}
	parameters.Add("scope", "oob")
	parameters.Add("grant_type", "client_credentials")
	URL.RawQuery = parameters.Encode()
	req, err := http.NewRequest("POST", URL.String(), nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", api.Authentication())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := api.defaultHTTPClient().Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		err := errors.New("bad request getnet server")
		return result, err
	}
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	payload := AuthResponse{}
	err = json.Unmarshal(byt, &payload)
	if err != nil {
		return result, err
	}
	result.Response = payload
	return result, nil
}
