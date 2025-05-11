package breezeconnect

import (
	"encoding/json"
	"fmt"
)

// GetPortfolioHoldings retrieves the portfolio holdings
func (c *Client) GetPortfolioHoldings() ([]PortfolioHolding, error) {
	response, err := c.makeRequest("GET", "/portfolio/holdings", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Status  int                `json:"status"`
		Message string             `json:"message"`
		Data    []PortfolioHolding `json:"data"`
	}

	if err := json.Unmarshal(response, &result); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return result.Data, nil
}

// GetPositions retrieves the current positions
func (c *Client) GetPositions() ([]Position, error) {
	response, err := c.makeRequest("GET", "/portfolio/positions", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    []Position `json:"data"`
	}

	if err := json.Unmarshal(response, &result); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return result.Data, nil
}
