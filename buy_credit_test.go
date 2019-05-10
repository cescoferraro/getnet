package getnet_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/cescoferraro/getnet"
	"github.com/stretchr/testify/assert"
)

// TestSpec isk kjsdf
func TestSpecBuyCredit(t *testing.T) {
	api := sandboxTestAPI()
	auth, err := api.Auth(context.Background())
	assert.NoError(t, err)
	card, err := api.CardToken(context.Background(), auth.Response, "4242424242424242")
	assert.NoError(t, err)
	response, err := api.BuyCreditCard(context.Background(), auth.Response,
		getnet.BuyCreditCardPayload{
			SellerID: "f359fc2e-f532-4d86-89fc-62c3adbb6405",
			Amount:   100,
			Currency: "BRL",
			Order: getnet.BuyOrder{
				OrderID:     "12345",
				SalesTax:    10,
				ProductType: "service",
			},
			Customer: customer,
			Credit: getnet.BuyCredit{
				Delayed:            false,
				SaveCardData:       false,
				TransactionType:    "FULL",
				NumberInstallments: 1,
				SoftDescriptor:     "LOJA*TESTE",
				DynamicMcc:         1799,
				Card: getnet.Card{
					NumberToken:     card.Response.NumberToken,
					Bin:             "VISA",
					CardholderName:  "Francesco Ferraro",
					ExpirationMonth: "03",
					ExpirationYear:  "27",
					SecurityCode:    "699",
				},
			},
		})
	assert.NoError(t, err)
	log.Println(response.Code)
	responsee, err := api.CancelCreditCardBuy(context.Background(), auth.Response, getnet.CancelCreditCardBuyPayload{
		PaymentID:    response.Response.PaymentID,
		CancelAmount: 100,
	})
	assert.NoError(t, err)
	// // prettyPrint(response)

	// log.Println(responsee)
	log.Println(responsee.Code)
	log.Println(responsee.Code)
	log.Println(responsee.Code)
	log.Println(responsee.Code)
	log.Println(responsee.Code)
	// log.Println(responsee)
	// log.Println(response.Response.Status)
	// log.Println(response.PaymentID)
}

func prettyPrint(obj interface{}) {
	str, _ := json.MarshalIndent(&obj, "    ", "")
	log.Println(string(str))
}

var customer = getnet.Customer{
	CustomerID:   "12345",
	FirstName:    "12345",
	LastName:     "12345",
	Name:         "12345",
	Email:        "francescoaferraro@gmail.com",
	DocumentType: "CPF",
	BillingAddress: getnet.Address{
		Street:     "dkjfn",
		Number:     "dkjfn",
		District:   "dkjfn",
		City:       "dkjfn",
		State:      "dkjfn",
		Country:    "dkjfn",
		PostalCode: "dkjfn",
	},
}
