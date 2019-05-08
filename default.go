package getnet

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"time"
)

// API TODO: NEEDS COMMENT INFO
type API struct {
	ClientID     string
	ClientSecret string
	Verbose      bool
}

// New creates ane
func New(clientID, clientSecret string) *API {
	return &API{
		ClientSecret: clientSecret,
		ClientID:     clientID,
		Verbose:      true,
	}
}

func (api *API) getURL() (*url.URL, error) {
	var URL *url.URL
	URL, err := url.Parse("https://api-sandbox.getnet.com.br/")
	if err != nil {
		return URL, err
	}
	return URL, nil
}

// Authentication string
func (api *API) Authentication() string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(api.ClientID+":"+api.ClientSecret))
}
func (api *API) defaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 400,
	}
}
