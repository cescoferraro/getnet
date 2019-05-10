package getnet

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// CardTokenPayload sadfkjn
type CardTokenPayload struct {
	CardNumber string `json:"card_number"`
}

// CardResponse sdfkjn
type CardFullResponse struct {
	Code     int          `json:"code"`
	Response CardResponse `json:"response"`
}

// CardResponse sdfkjn
type CardResponse struct {
	NumberToken string `json:"number_token"`
}

// CardToken dsjkfndsfkfjng
func (api *API) CardToken(ctx context.Context, auth AuthResponse, cardNumber string) (CardFullResponse, error) {
	result := CardFullResponse{Code: api.DefaultCode()}
	URL, err := api.getURL()
	if err != nil {
		return result, err
	}
	URL.Path += "/v1/tokens/card"
	data := CardTokenPayload{
		CardNumber: cardNumber,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", URL.String(), body)
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", auth.Bearer())
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.defaultHTTPClient().Do(req)
	result.Code = resp.StatusCode
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
	payload := CardResponse{}
	err = json.Unmarshal(byt, &payload)
	if err != nil {
		return result, err
	}
	result.Response = payload
	return result, nil
}
