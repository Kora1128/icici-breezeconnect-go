package main

import (
	"fmt"
	"log"

	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect"
)

func main() {
	// Initialize the client
	client := breezeconnect.NewClient("your_api_key", "your_api_secret")

	// Generate session
	if err := client.GenerateSession(); err != nil {
		log.Fatalf("Error generating session: %v", err)
	}

	// Get portfolio holdings
	holdings, err := client.GetPortfolioHoldings()
	if err != nil {
		log.Fatalf("Error getting portfolio holdings: %v", err)
	}

	fmt.Println("Portfolio Holdings:")
	for _, holding := range holdings {
		fmt.Printf("Symbol: %s, Quantity: %d, Average Price: %.2f\n",
			holding.Symbol, holding.Quantity, holding.AveragePrice)
	}

	// Get positions
	positions, err := client.GetPositions()
	if err != nil {
		log.Fatalf("Error getting positions: %v", err)
	}

	fmt.Println("\nPositions:")
	for _, position := range positions {
		fmt.Printf("Symbol: %s, Quantity: %d, PnL: %.2f\n",
			position.Symbol, position.Quantity, position.PnL)
	}
}
