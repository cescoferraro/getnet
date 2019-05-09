package getnet

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Auth dsjkfndsfkfjng
func (api *API) Auth(ctx context.Context) (AuthResponse, error) {
	result := AuthResponse{}
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
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(byt, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
