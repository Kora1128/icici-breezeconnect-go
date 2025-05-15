package services

import (
	"errors"
	"testing"

	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect"
	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect/mock"
	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect/models"
)

func TestNewFundsService(t *testing.T) {
	client := breezeconnect.NewClient("test_api_key", "test_api_secret")
	service := NewFundsService(client)

	if service.client != client {
		t.Error("Expected service.client to be the same as the input client")
	}
}

func TestGetFunds(t *testing.T) {
	// Create test response
	testResponse := models.Funds{
		Success: struct {
			AvailableBalance float64 `json:"available_balance"`
		}{
			AvailableBalance: 100000.50,
		},
		Status: 200,
	}

	// Create mock client
	mockClient := mock.NewMockClient()
	mockClient.SetResponse("/funds", testResponse)

	// Create service with mock client
	service := NewFundsService(mockClient)

	// Test GetFunds
	funds, err := service.GetFunds()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if funds.Success.AvailableBalance != 100000.50 {
		t.Errorf("Expected available balance to be 100000.50, got %.2f", funds.Success.AvailableBalance)
	}
}

func TestGetFundsError(t *testing.T) {
	// Create mock client with error
	mockClient := mock.NewMockClient()
	mockClient.SetError("/funds", errors.New("API error"))

	// Create service with mock client
	service := NewFundsService(mockClient)

	// Test GetFunds with error
	_, err := service.GetFunds()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
