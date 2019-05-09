package getnet_test

import (
	"context"
	"log"
	"testing"

	"github.com/cescoferraro/getnet"
	// . "github.com/smartystreets/goconvey/convey"
)

// TestSpec isk kjsdf
func TestSpec(t *testing.T) {
	api := getnet.New("f92b4308-cf9f-466f-8066-b21da774104a", "33458814-6da5-4348-93a6-3fab5ebb8360")
	auth := getnet.AuthResponse{}
	card := getnet.CardResonse{}
	cardverification := getnet.CardVerificationResponse{}
	t.Run("Auth", func(t *testing.T) {
		response, err := api.Auth(context.Background())
		if err != nil {
			panic(err)
		}
		auth = response
	})
	t.Run("Card", func(t *testing.T) {
		response, err := api.CardToken(context.Background(), auth, "4242424242424242")
		if err != nil {
			panic(err)
		}
		card = response
	})
	t.Run("Verify Card", func(t *testing.T) {
		response, err := api.CardVerification(context.Background(), auth,
			getnet.CardVerificationPayload{
				NumberToken:     card.NumberToken,
				Brand:           "VISA",
				CardholderName:  "Francesco Ferraro",
				ExpirationMonth: "03",
				ExpirationYear:  "19",
				SecurityCode:    "699",
			})
		if err != nil {
			panic(err)
		}
		cardverification = response
	})
	log.Println(auth.AccessToken)
	log.Println(card.NumberToken)
	log.Println(cardverification.Status)
	log.Println(cardverification.IsVerified())

}
