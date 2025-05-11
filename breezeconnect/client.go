package breezeconnect

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseURL = "https://api.icicidirect.com/breezeapi/api/v1"
)

// Client represents the Breeze API client
type Client struct {
	apiKey     string
	apiSecret  string
	sessionKey string
	httpClient *http.Client
}

// NewClient creates a new Breeze API client
func NewClient(apiKey, apiSecret string) *Client {
	return &Client{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// GenerateSession generates a new session token
func (c *Client) GenerateSession() error {
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z")
	checksum := c.generateChecksum(timestamp)

	payload := map[string]string{
		"api_key":   c.apiKey,
		"timestamp": timestamp,
		"checksum":  checksum,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err)
	}

	req, err := http.NewRequest("POST", baseURL+"/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	if sessionToken, ok := result["session_token"].(string); ok {
		c.sessionKey = sessionToken
		return nil
	}

	return fmt.Errorf("failed to get session token")
}

// generateChecksum generates the checksum for authentication
func (c *Client) generateChecksum(timestamp string) string {
	data := c.apiKey + timestamp
	h := hmac.New(sha256.New, []byte(c.apiSecret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// makeRequest makes an authenticated request to the API
func (c *Client) makeRequest(method, endpoint string, payload interface{}) ([]byte, error) {
	var body io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("error marshaling payload: %v", err)
		}
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, baseURL+endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.sessionKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", string(responseBody))
	}

	return responseBody, nil
}
