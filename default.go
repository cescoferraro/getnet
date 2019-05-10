package getnet

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"time"
)

// API TODO: NEEDS COMMENT INFO
type API struct {
	URL          string
	ClientID     string
	ClientSecret string
	Verbose      bool
}

// NewAPI creates ane
func NewAPI(api API) *API {
	return &api
}

func (api *API) getURL() (*url.URL, error) {
	var URL *url.URL
	URL, err := url.ParseRequestURI(api.URL)
	if err != nil {
		return URL, err
	}
	return URL, nil
}

// Authentication string
func (api *API) Authentication() string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(api.ClientID+":"+api.ClientSecret))
}

// DefaultCode sdfkj
func (api *API) DefaultCode() int {
	return 400
}
func (api *API) defaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 400,
	}
}
