package mockapi

import (
	"errors"
	"os"

	"rsc.io/quote"

	_ "github.com/joho/godotenv/autoload"
)

func Request(apiKey string) (string, error) {
	if err := auth(apiKey); err != nil {
		return "", err
	}

	return quote.Hello(), nil
}

func auth(apiKey string) error {
	secret := os.Getenv("API_KEY")
	if apiKey != secret {
		return errors.New("invalid API key")
	}

	return nil
}
