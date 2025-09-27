package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"goodbye-world/mockapi"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	quote, err := mockapi.Request(apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(quote)
}
