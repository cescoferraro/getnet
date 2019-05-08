package getnet

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
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
	log.Println(URL.String())
	log.Println(URL.String())
	log.Println(api.Authentication())
	// body := strings.NewReader(`scope=oob&grant_type=client_credentialsn`)
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

// Basic ZjkyYjQzMDgtY2Y5Zi00NjZmLTgwNjYtYjIxZGE3NzQxMDRhOjMzNDU4ODE0LTZkYTUtNDM0OC05M2E2LTNmYWI1ZWJiODM2MA==
// Basic ZjkyYjQzMDgtY2Y5Zi00NjZmLTgwNjYtYjIxZGE3NzQxMDRhOjMzNDU4ODE0LTZkYTUtNDM0OC05M2E2LTNmYWI1ZWJiODM2MA==
