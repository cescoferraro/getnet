package getnet

import (
	"context"
	"net/http"
	"strings"
)

// RecipientBalance dsjkfndsfkfjng
func (api *API) Auth(ctx context.Context, recipientID string) (interface{}, error) {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	URL, err := api.getURL()
	if err != nil {
		return nil, err
	}
	URL.Path += "/auth/oauth/v2/token"
	body := strings.NewReader(`scope=oob&grant_type=client_credentials`)
	req, err := http.NewRequest("POST", URL.String(), body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authorization", "Basic ZjkyYjQzMDgtY2Y5Zi00NjZmLTgwNjYtYjIxZGE3NzQxMDRhOjMzNDU4ODE0LTZkYTUtNDM0OC05M2E2LTNmYWI1ZWJiODM2MA==")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	// var balance types.PagarMeTransactionBalance
	// URL.Path += "/1/recipients/" + recipientID + "/balance"
	// parameters := url.Values{}
	// parameters.Add("api_key", api.Key)
	// URL.RawQuery = parameters.Encode()
	// req, err := http.NewRequest("GET", URL.String(), nil)
	// if err != nil {
	// 	return balance, err
	// }
	// if api.Verbose {
	// 	fmt.Println(URL.String())
	// }
	// req.Header.Set("Content-Type", "application/json")
	// resp, err := ctxhttp.Do(ctx, api.defaultHTTPClient(), req)
	// if err != nil {
	// 	return balance, err
	// }
	// defer resp.Body.Close()
	// byt, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return balance, err
	// }
	// err = json.Unmarshal(byt, &balance)
	// if err != nil {
	// 	return balance, err
	// }
	return resp, nil
}
