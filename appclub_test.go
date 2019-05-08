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
	t.Run("LEITOR READ REDIRECT", func(t *testing.T) {
		log.Println(9999)
		api := getnet.New("f92b4308-cf9f-466f-8066-b21da774104a", "33458814-6da5-4348-93a6-3fab5ebb8360")
		response, err := api.Auth(context.Background())
		if err != nil {
			panic(err)
		}
		log.Println(response)
	})
}
