package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect"
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

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Test request method
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		// Test request path
		if r.URL.Path != "/funds" {
			t.Errorf("Expected /funds path, got %s", r.URL.Path)
		}

		// Test headers
		if r.Header.Get("X-SessionToken") == "" {
			t.Error("Expected X-SessionToken header to be present")
		}

		// Return test response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(testResponse)
	}))
	defer server.Close()

	client := breezeconnect.NewClient("test_api_key", "test_api_secret")
	client.SetSessionKey("test_session")
	service := NewFundsService(client)

	funds, err := service.GetFunds()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if funds.Success.AvailableBalance != 100000.50 {
		t.Errorf("Expected available balance to be 100000.50, got %.2f", funds.Success.AvailableBalance)
	}
}

func TestGetFundsError(t *testing.T) {
	// Create a test server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Funds{
			Status: 400,
			Error:  "Invalid session token",
		})
	}))
	defer server.Close()

	client := breezeconnect.NewClient("test_api_key", "test_api_secret")
	client.SetSessionKey("invalid_session")
	service := NewFundsService(client)

	_, err := service.GetFunds()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
