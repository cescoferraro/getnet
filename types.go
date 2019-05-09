package getnet

// AuthResponse
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIN   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

// Bearer sdkfj
func (auth AuthResponse) Bearer() string {
	return "Bearer " + auth.AccessToken
}
