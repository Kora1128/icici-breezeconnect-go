package mock

import (
	"encoding/json"
	"fmt"
)

// MockClient is a mock implementation of the Breeze API client
type MockClient struct {
	responses map[string]interface{}
	errors    map[string]error
}

// NewMockClient creates a new mock client
func NewMockClient() *MockClient {
	return &MockClient{
		responses: make(map[string]interface{}),
		errors:    make(map[string]error),
	}
}

// SetResponse sets a mock response for a given endpoint
func (m *MockClient) SetResponse(endpoint string, response interface{}) {
	m.responses[endpoint] = response
}

// SetError sets a mock error for a given endpoint
func (m *MockClient) SetError(endpoint string, err error) {
	m.errors[endpoint] = err
}

// MakeRequest implements the client interface for testing
func (m *MockClient) MakeRequest(method, endpoint string, payload interface{}) ([]byte, error) {
	// Check if we have an error for this endpoint
	if err, exists := m.errors[endpoint]; exists {
		return nil, err
	}

	// Check if we have a response for this endpoint
	if response, exists := m.responses[endpoint]; exists {
		jsonData, err := json.Marshal(response)
		if err != nil {
			return nil, fmt.Errorf("error marshaling mock response: %v", err)
		}
		return jsonData, nil
	}

	return nil, fmt.Errorf("no mock response set for endpoint: %s", endpoint)
}

// GetAPIKey returns a mock API key
func (m *MockClient) GetAPIKey() string {
	return "mock_api_key"
}

// SetSessionKey is a no-op for the mock client
func (m *MockClient) SetSessionKey(sessionKey string) {
	// No-op
}
