package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect"
	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect/models"
)

func TestNewCustomerService(t *testing.T) {
	client := breezeconnect.NewClient("test_api_key", "test_api_secret")
	service := NewCustomerService(client)

	if service.client != client {
		t.Error("Expected service.client to be the same as the input client")
	}
}

func TestGetCustomerDetails(t *testing.T) {
	// Create test response
	testResponse := models.CustomerDetails{
		Success: struct {
			ExgTradeDate struct {
				NSE string `json:"NSE"`
				BSE string `json:"BSE"`
				FNO string `json:"FNO"`
				NDX string `json:"NDX"`
			} `json:"exg_trade_date"`
			ExgStatus struct {
				NSE string `json:"NSE"`
				BSE string `json:"BSE"`
				FNO string `json:"FNO"`
				NDX string `json:"NDX"`
			} `json:"exg_status"`
			SegmentsAllowed struct {
				Trading     string `json:"Trading"`
				Equity      string `json:"Equity"`
				Derivatives string `json:"Derivatives"`
				Currency    string `json:"Currency"`
			} `json:"segments_allowed"`
			IDirectUserID        string `json:"idirect_userid"`
			IDirectUserName      string `json:"idirect_user_name"`
			IDirectORDTYP        string `json:"idirect_ORD_TYP"`
			IDirectLastLoginTime string `json:"idirect_lastlogin_time"`
			SessionToken         string `json:"session_token"`
		}{
			IDirectUserID:   "test_user",
			IDirectUserName: "Test User",
			SessionToken:    "test_session_token",
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
		if r.URL.Path != "/customerdetails" {
			t.Errorf("Expected /customerdetails path, got %s", r.URL.Path)
		}

		// Return test response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(testResponse)
	}))
	defer server.Close()

	client := breezeconnect.NewClient("test_api_key", "test_api_secret")
	service := NewCustomerService(client)

	details, err := service.GetCustomerDetails("test_session")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if details.Success.IDirectUserID != "test_user" {
		t.Errorf("Expected user ID to be test_user, got %s", details.Success.IDirectUserID)
	}

	if details.Success.IDirectUserName != "Test User" {
		t.Errorf("Expected user name to be Test User, got %s", details.Success.IDirectUserName)
	}

	if details.Success.SessionToken != "test_session_token" {
		t.Errorf("Expected session token to be test_session_token, got %s", details.Success.SessionToken)
	}
}
