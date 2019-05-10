package getnet

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// BuyCreditCardPayload sdkfjn
type BuyCreditCardPayload struct {
	SellerID                    string                        `json:"seller_id"`
	Amount                      int                           `json:"amount"`
	Currency                    string                        `json:"currency"`
	Order                       BuyOrder                      `json:"order"`
	Customer                    Customer                      `json:"customer"`
	Credit                      BuyCredit                     `json:"credit"`
	MaketPlaceSubSellerPayments []MaketPlaceSubSellerPayments `json:"marketplace_subseller_payments,omitempty"`
}

type BuyCreditCardResponseFull struct {
	Code     int                   `json:"code"`
	Response BuyCreditCardResponse `json:"response"`
}

// BuyCreditCardResponse sdkfjn
type BuyCreditCardResponse struct {
	Status    string    `json:"status"`
	SellerID  string    `json:"seller_id"`
	PaymentID string    `json:"payment_id"`
	Amount    int       `json:"amount"`
	Currency  string    `json:"currency"`
	Credit    BuyCredit `json:"credit"`
}

// BuyCreditCard dsjkfndsfkfjng
func (api *API) BuyCreditCard(ctx context.Context, auth AuthResponse, payload BuyCreditCardPayload) (BuyCreditCardResponseFull, error) {
	result := BuyCreditCardResponseFull{Code: api.DefaultCode()}
	URL, err := api.getURL()
	if err != nil {
		return result, err
	}
	URL.Path += "/v1/payments/credit"
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
	response := BuyCreditCardResponse{}
	err = json.Unmarshal(byt, &response)
	if err != nil {
		return result, err
	}
	result.Response = response
	return result, nil
}
