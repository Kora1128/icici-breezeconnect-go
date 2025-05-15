package main

import (
	"fmt"
	"log"

	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect"
	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect/services"
)

func main() {
	// Initialize the client
	client := breezeconnect.NewClient("your_api_key", "your_api_secret")

	// Initialize services
	customerService := services.NewCustomerService(client)
	portfolioService := services.NewPortfolioService(client)

	// Get customer details and session token
	customerDetails, err := customerService.GetCustomerDetails("your_session_token")
	if err != nil {
		log.Fatalf("Error getting customer details: %v", err)
	}

	fmt.Printf("Customer Details:\n")
	fmt.Printf("User ID: %s\n", customerDetails.Success.IDirectUserID)
	fmt.Printf("User Name: %s\n", customerDetails.Success.IDirectUserName)

	// Get portfolio holdings
	holdings, err := portfolioService.GetPortfolioHoldings()
	if err != nil {
		log.Fatalf("Error getting portfolio holdings: %v", err)
	}

	fmt.Println("\nPortfolio Holdings:")
	for _, holding := range holdings {
		fmt.Printf("Symbol: %s, Quantity: %d, Average Price: %.2f, Total Value: %.2f\n",
			holding.Symbol, holding.Quantity, holding.AveragePrice, holding.TotalValue)
	}

	// Get positions
	positions, err := portfolioService.GetPositions()
	if err != nil {
		log.Fatalf("Error getting positions: %v", err)
	}

	fmt.Println("\nPositions:")
	for _, position := range positions {
		fmt.Printf("Symbol: %s, Quantity: %d, PnL: %.2f, Product Type: %s\n",
			position.Symbol, position.Quantity, position.PnL, position.ProductType)
		if position.ProductType == "options" {
			fmt.Printf("  Strike Price: %.2f, Option Type: %s, Expiry: %s\n",
				position.StrikePrice, position.OptionType, position.ExpiryDate)
		}
	}
}
