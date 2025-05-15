package breezeconnect

import (
	"errors"
	"testing"

	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect/mock"
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
	// Create mock client
	mockClient := mock.NewMockClient()

	// Test successful request
	testResponse := map[string]interface{}{
		"Success": "test",
		"Status":  200,
	}
	mockClient.SetResponse("/test", testResponse)

	// Test request without session key
	resp, err := mockClient.MakeRequest("GET", "/test", nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if string(resp) != `{"Status":200,"Success":"test"}` {
		t.Errorf("Expected test response, got %s", string(resp))
	}

	// Test request with error
	mockClient.SetError("/test", errors.New("API error"))
	_, err = mockClient.MakeRequest("GET", "/test", nil)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Test request with no mock response
	_, err = mockClient.MakeRequest("GET", "/nonexistent", nil)
	if err == nil {
		t.Error("Expected error for nonexistent endpoint, got nil")
	}
}
