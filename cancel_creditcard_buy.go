package getnet

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// CancelCreditCardBuyPayload sdfjk
type CancelCreditCardBuyPayload struct {
	PaymentID       string `json:"payment_id"`
	CancelAmount    int    `json:"cancel_amount"`
	CancelCustomKey string `json:"cancel_custom_key,omitempty"`
}

// CancelCreditCardBuyResponse
type CancelCreditCardBuyFull struct {
	Code     int                         `json:"code"`
	Response CancelCreditCardBuyResponse `json:"response"`
}

// CancelCreditCardBuyResponse asdfkjn
type CancelCreditCardBuyResponse struct {
	SellerID        string    `json:"seller_id"`
	PaymentID       string    `json:"payment_id"`
	CancelRequestAt time.Time `json:"cancel_request_at"`
	CancelRequestID string    `json:"cancel_request_id"`
	CancelCustomKey string    `json:"cancel_custom_key"`
	Status          string    `json:"status"`
}

// CancelCreditCardBuy dsjkfndsfkfjng
func (api *API) CancelCreditCardBuy(
	ctx context.Context,
	auth AuthResponse,
	all CancelCreditCardBuyPayload,
) (CancelCreditCardBuyFull, error) {
	result := CancelCreditCardBuyFull{Code: api.DefaultCode()}
	URL, err := api.getURL()
	if err != nil {
		return result, err
	}
	URL.Path += "v1/payments/cancel/request"
	payloadBytes, err := json.Marshal(all)
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
	result.Code = resp.StatusCode
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	payload := CancelCreditCardBuyResponse{}
	err = json.Unmarshal(byt, &payload)
	if err != nil {
		return result, err
	}
	result.Response = payload
	return result, nil
}
