package getnet_test

import (
	"context"
	"log"
	"testing"

	"github.com/cescoferraro/getnet"
	"github.com/stretchr/testify/assert"
)

func badURLSandboxTestAPI() *getnet.API {
	return getnet.NewAPI(getnet.API{
		URL:          "22",
		ClientID:     "f92b4308cf9f-466f-8066-b21da774104a",
		ClientSecret: "33458814-6da5-4348-93a6-3fab5ebb8360",
		Verbose:      true,
	})
}
func badSandboxTestAPI() *getnet.API {
	return getnet.NewAPI(getnet.API{
		URL:          "https://api-sandbox.getnet.com.br",
		ClientID:     "f92b4308cf9f-466f-8066-b21da774104a",
		ClientSecret: "33458814-6da5-4348-93a6-3fab5ebb8360",
		Verbose:      true,
	})
}

func sandboxTestAPI() *getnet.API {
	return getnet.NewAPI(getnet.API{
		URL:          "https://api-sandbox.getnet.com.br",
		ClientID:     "f92b4308-cf9f-466f-8066-b21da774104a",
		ClientSecret: "33458814-6da5-4348-93a6-3fab5ebb8360",
		Verbose:      true,
	})
}

// TestSpecVerifyCard isk kjsdf
func TestSpecVerifyCard(t *testing.T) {
	badurltestapi := badURLSandboxTestAPI()
	_, errr := badurltestapi.Auth(context.Background())
	assert.Error(t, errr)
	badtestapi := badSandboxTestAPI()
	_, err := badtestapi.Auth(context.Background())
	assert.Error(t, err)
	testapi := sandboxTestAPI()
	auth, err := testapi.Auth(context.Background())
	assert.NoError(t, err)
	log.Println(auth)
	card, err := testapi.CardToken(context.Background(), auth.Response, "4242424242424242")
	assert.NoError(t, err)
	cardverification, err := testapi.CardVerification(context.Background(), auth.Response,
		getnet.CardVerificationPayload{
			NumberToken:     card.Response.NumberToken,
			Brand:           "VISA",
			CardholderName:  "Francesco Ferraro",
			ExpirationMonth: "03",
			ExpirationYear:  "19",
			SecurityCode:    "699",
		})
	assert.NoError(t, err)
	// prettyPrint(cardverification)
	log.Println(cardverification)

}
