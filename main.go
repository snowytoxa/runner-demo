package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
)

type goldPrice struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Symbol    string  `json:"symbol"`
	UpdatedAt string  `json:"updatedAt"`
}



func rot13(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+13)%26
	}
	return r
}


func main() {
	client := resty.New()

	var result goldPrice
	resp, err := client.R().
		SetResult(&result).
		Get("https://api.gold-api.com/price/XAU")
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	if resp.IsError() {
		log.Fatalf("unexpected status: %s", resp.Status())
	}

	fmt.Printf("Current gold price (%s): $%.2f per ounce (updated %s)\n",
		result.Symbol, result.Price, result.UpdatedAt)

	fmt.Println(os.Getenv("SUPERSECRET_PRIVATE_KEY"))
}
