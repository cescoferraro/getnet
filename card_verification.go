package getnet

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// CardVerificationPayload sdfkjn
type CardVerificationPayload struct {
	NumberToken     string `json:"number_token"`
	Brand           string `json:"brand"`
	CardholderName  string `json:"cardholder_name"`
	ExpirationMonth string `json:"expiration_month"`
	ExpirationYear  string `json:"expiration_year"`
	SecurityCode    string `json:"security_code"`
}

// CardVerificationResponse dfjkn
type CardVerificationResponse struct {
	Status            string `json:"status"`
	VerificationID    string `json:"verification_id"`
	AuthorizationCode string `json:"authorization_code"`
}

// IsVerified sdfkjn
func (cardverification CardVerificationResponse) IsVerified() bool {
	if cardverification.Status == "VERIFIED" {
		return true
	}
	return false
}

// CardVerification dsjkfndsfkfjng
func (api *API) CardVerification(
	ctx context.Context,
	auth AuthResponse,
	payload CardVerificationPayload,
) (CardVerificationResponse, error) {
	result := CardVerificationResponse{}
	URL, err := api.getURL()
	if err != nil {
		return result, err
	}
	URL.Path += "/v1/cards/verification"
	payloadBytes, err := json.Marshal(payload)
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
	log.Println(string(byt))
	err = json.Unmarshal(byt, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
