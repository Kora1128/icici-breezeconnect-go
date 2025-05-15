package breezeconnect

import (
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
}
