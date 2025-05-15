package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect"
	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect/models"
)

func TestNewPortfolioService(t *testing.T) {
	client := breezeconnect.NewClient("test_api_key", "test_api_secret")
	service := NewPortfolioService(client)

	if service.client != client {
		t.Error("Expected service.client to be the same as the input client")
	}
}

func TestGetPortfolioHoldings(t *testing.T) {
	// Create test response
	testResponse := models.PortfolioHoldingsResponse{
		Success: []models.PortfolioHolding{
			{
				Symbol:          "RELIANCE",
				Quantity:        100,
				AveragePrice:    2500.50,
				LastTradedPrice: 2600.00,
				PnL:             9950.00,
				ProductType:     "equity",
				Exchange:        "NSE",
				ISIN:            "INE002A01018",
				Name:            "Reliance Industries Ltd",
				TotalValue:      260000.00,
				FreeQuantity:    100,
				LockedQuantity:  0,
			},
		},
		Status: 200,
	}

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Test request method
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		// Test request path
		if r.URL.Path != "/portfolio/holdings" {
			t.Errorf("Expected /portfolio/holdings path, got %s", r.URL.Path)
		}

		// Return test response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(testResponse)
	}))
	defer server.Close()

	client := breezeconnect.NewClient("test_api_key", "test_api_secret")
	client.SetSessionKey("test_session")
	service := NewPortfolioService(client)

	holdings, err := service.GetPortfolioHoldings()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(holdings) != 1 {
		t.Errorf("Expected 1 holding, got %d", len(holdings))
	}

	holding := holdings[0]
	if holding.Symbol != "RELIANCE" {
		t.Errorf("Expected symbol to be RELIANCE, got %s", holding.Symbol)
	}

	if holding.Quantity != 100 {
		t.Errorf("Expected quantity to be 100, got %d", holding.Quantity)
	}
}

func TestGetPositions(t *testing.T) {
	// Create test response
	testResponse := models.PositionsResponse{
		Success: []models.Position{
			{
				Symbol:          "NIFTY",
				Quantity:        1,
				AveragePrice:    19500.50,
				LastTradedPrice: 19600.00,
				PnL:             99.50,
				ProductType:     "futures",
				Exchange:        "NFO",
				ExpiryDate:      "2024-03-28",
				TotalValue:      19600.00,
				FreeQuantity:    1,
				LockedQuantity:  0,
			},
		},
		Status: 200,
	}

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Test request method
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		// Test request path
		if r.URL.Path != "/portfolio/positions" {
			t.Errorf("Expected /portfolio/positions path, got %s", r.URL.Path)
		}

		// Return test response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(testResponse)
	}))
	defer server.Close()

	client := breezeconnect.NewClient("test_api_key", "test_api_secret")
	client.SetSessionKey("test_session")
	service := NewPortfolioService(client)

	positions, err := service.GetPositions()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(positions) != 1 {
		t.Errorf("Expected 1 position, got %d", len(positions))
	}

	position := positions[0]
	if position.Symbol != "NIFTY" {
		t.Errorf("Expected symbol to be NIFTY, got %s", position.Symbol)
	}

	if position.Quantity != 1 {
		t.Errorf("Expected quantity to be 1, got %d", position.Quantity)
	}

	if position.ProductType != "futures" {
		t.Errorf("Expected product type to be futures, got %s", position.ProductType)
	}
}
