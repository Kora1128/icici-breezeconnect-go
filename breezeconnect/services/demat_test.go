package services

import (
	"errors"
	"testing"

	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect"
	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect/mock"
	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect/models"
)

func TestNewDematService(t *testing.T) {
	client := breezeconnect.NewClient("test_api_key", "test_api_secret")
	service := NewDematService(client)

	if service.client != client {
		t.Error("Expected service.client to be the same as the input client")
	}
}

func TestGetDematHoldings(t *testing.T) {
	// Create test response
	testResponse := models.DematHoldingsResponse{
		Success: []models.DematHolding{
			{
				ISIN:            "INE002A01018",
				Symbol:          "RELIANCE",
				Quantity:        100,
				ProductType:     "equity",
				Exchange:        "NSE",
				AveragePrice:    2500.50,
				LastTradedPrice: 2600.00,
				PnL:             9950.00,
				TotalValue:      260000.00,
			},
			{
				ISIN:            "INE848E01016",
				Symbol:          "TCS",
				Quantity:        50,
				ProductType:     "equity",
				Exchange:        "NSE",
				AveragePrice:    3500.75,
				LastTradedPrice: 3600.00,
				PnL:             4962.50,
				TotalValue:      180000.00,
			},
		},
		Status: 200,
	}

	// Create mock client
	mockClient := mock.NewMockClient()
	mockClient.SetResponse("/demat/holdings", testResponse)

	// Create service with mock client
	service := NewDematService(mockClient)

	// Test GetDematHoldings
	holdings, err := service.GetDematHoldings()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(holdings) != 2 {
		t.Errorf("Expected 2 holdings, got %d", len(holdings))
	}

	// Test first holding
	holding1 := holdings[0]
	if holding1.Symbol != "RELIANCE" {
		t.Errorf("Expected symbol to be RELIANCE, got %s", holding1.Symbol)
	}
	if holding1.Quantity != 100 {
		t.Errorf("Expected quantity to be 100, got %d", holding1.Quantity)
	}
	if holding1.ISIN != "INE002A01018" {
		t.Errorf("Expected ISIN to be INE002A01018, got %s", holding1.ISIN)
	}

	// Test second holding
	holding2 := holdings[1]
	if holding2.Symbol != "TCS" {
		t.Errorf("Expected symbol to be TCS, got %s", holding2.Symbol)
	}
	if holding2.Quantity != 50 {
		t.Errorf("Expected quantity to be 50, got %d", holding2.Quantity)
	}
	if holding2.ISIN != "INE848E01016" {
		t.Errorf("Expected ISIN to be INE848E01016, got %s", holding2.ISIN)
	}
}

func TestGetDematHoldingsError(t *testing.T) {
	// Create mock client with error
	mockClient := mock.NewMockClient()
	mockClient.SetError("/demat/holdings", errors.New("API error"))

	// Create service with mock client
	service := NewDematService(mockClient)

	// Test GetDematHoldings with error
	_, err := service.GetDematHoldings()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestGetDematHoldingsEmpty(t *testing.T) {
	// Create test response with empty holdings
	testResponse := models.DematHoldingsResponse{
		Success: []models.DematHolding{},
		Status:  200,
	}

	// Create mock client
	mockClient := mock.NewMockClient()
	mockClient.SetResponse("/demat/holdings", testResponse)

	// Create service with mock client
	service := NewDematService(mockClient)

	// Test GetDematHoldings
	holdings, err := service.GetDematHoldings()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(holdings) != 0 {
		t.Errorf("Expected 0 holdings, got %d", len(holdings))
	}
}
