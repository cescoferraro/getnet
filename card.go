package getnet

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CardTokenPayload sadfkjn
type CardTokenPayload struct {
	CardNumber string `json:"card_number"`
}

// CardResonse sdfkjn
type CardResonse struct {
	NumberToken string `json:"number_token"`
}

// CardToken dsjkfndsfkfjng
func (api *API) CardToken(ctx context.Context, auth AuthResponse, cardNumber string) (CardResonse, error) {
	result := CardResonse{}
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
