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

// GetAPIKey returns the API key
func (c *Client) GetAPIKey() string {
	return c.apiKey
}

// SetSessionKey sets the session key
func (c *Client) SetSessionKey(sessionKey string) {
	c.sessionKey = sessionKey
}

// MakeRequest makes an authenticated request to the API
func (c *Client) MakeRequest(method, endpoint string, payload interface{}) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z")

	var jsonData []byte
	var err error
	if payload != nil {
		jsonData, err = json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("error marshaling payload: %v", err)
		}
	}

	req, err := http.NewRequest(method, baseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Only add authentication headers if we have a session key
	// This ensures GetCustomerDetails works without session token
	if c.sessionKey != "" {
		checksum := c.generateChecksum(timestamp, string(jsonData))
		req.Header.Set("X-Checksum", "token "+checksum)
		req.Header.Set("X-Timestamp", timestamp)
		req.Header.Set("X-AppKey", c.apiKey)
		req.Header.Set("X-SessionToken", c.sessionKey)
	}

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

// generateChecksum generates the checksum for authentication
func (c *Client) generateChecksum(timestamp, jsonData string) string {
	data := timestamp + jsonData + c.apiSecret
	h := hmac.New(sha256.New, []byte(c.apiSecret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
