package breezeconnect

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	apiKey := "test_api_key"
	apiSecret := "test_api_secret"

	client := NewClient(apiKey, apiSecret)

	if client.apiKey != apiKey {
		t.Errorf("Expected apiKey to be %s, got %s", apiKey, client.apiKey)
	}

	if client.apiSecret != apiSecret {
		t.Errorf("Expected apiSecret to be %s, got %s", apiSecret, client.apiSecret)
	}

	if client.sessionKey != "" {
		t.Errorf("Expected sessionKey to be empty, got %s", client.sessionKey)
	}

	if client.httpClient == nil {
		t.Error("Expected httpClient to be initialized")
	}
}

func TestGenerateChecksum(t *testing.T) {
	client := NewClient("test_api_key", "test_api_secret")
	timestamp := "2024-01-01T00:00:00.000Z"
	jsonData := "test_data"

	checksum := client.generateChecksum(timestamp, jsonData)

	if checksum == "" {
		t.Error("Expected non-empty checksum")
	}

	// Test that same input produces same checksum
	checksum2 := client.generateChecksum(timestamp, jsonData)
	if checksum != checksum2 {
		t.Error("Expected same checksum for same input")
	}

	// Test that different input produces different checksum
	checksum3 := client.generateChecksum(timestamp, "different_data")
	if checksum == checksum3 {
		t.Error("Expected different checksum for different input")
	}
}

func TestMakeRequest(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Test headers
		if r.Header.Get("Content-Type") != "application/json" {
			t.Error("Expected Content-Type header to be application/json")
		}

		// Test authentication headers when session key is set
		if r.Header.Get("X-SessionToken") != "" {
			if r.Header.Get("X-Checksum") == "" {
				t.Error("Expected X-Checksum header when session token is present")
			}
			if r.Header.Get("X-Timestamp") == "" {
				t.Error("Expected X-Timestamp header when session token is present")
			}
			if r.Header.Get("X-AppKey") == "" {
				t.Error("Expected X-AppKey header when session token is present")
			}
		}

		// Return test response
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"Success": "test", "Status": 200}`))
	}))
	defer server.Close()

	client := NewClient("test_api_key", "test_api_secret")

	// Test request without session key
	resp, err := client.MakeRequest("GET", "/test", nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if string(resp) != `{"Success": "test", "Status": 200}` {
		t.Errorf("Expected test response, got %s", string(resp))
	}

	// Test request with session key
	client.SetSessionKey("test_session")
	resp, err = client.MakeRequest("GET", "/test", nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if string(resp) != `{"Success": "test", "Status": 200}` {
		t.Errorf("Expected test response, got %s", string(resp))
	}
}
