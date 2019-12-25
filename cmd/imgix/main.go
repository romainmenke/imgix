package main

import (
	"log"
	"os"

	"github.com/romainmenke/imgix/v4/api"
)

func main() {
	client, err := api.New(api.APIKeyOption(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	_ = client
}
