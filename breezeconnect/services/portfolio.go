package services

import (
	"encoding/json"
	"fmt"

	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect"
	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect/models"
)

// PortfolioService handles portfolio-related API calls
type PortfolioService struct {
	client *breezeconnect.Client
}

// NewPortfolioService creates a new portfolio service
func NewPortfolioService(client *breezeconnect.Client) *PortfolioService {
	return &PortfolioService{client: client}
}

// GetPortfolioHoldings retrieves the portfolio holdings
func (s *PortfolioService) GetPortfolioHoldings() ([]models.PortfolioHolding, error) {
	response, err := s.client.MakeRequest("GET", "/portfolio/holdings", nil)
	if err != nil {
		return nil, fmt.Errorf("error getting portfolio holdings: %v", err)
	}

	var result models.PortfolioHoldingsResponse
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return result.Success, nil
}

// GetPositions retrieves the current positions
func (s *PortfolioService) GetPositions() ([]models.Position, error) {
	response, err := s.client.MakeRequest("GET", "/portfolio/positions", nil)
	if err != nil {
		return nil, fmt.Errorf("error getting positions: %v", err)
	}

	var result models.PositionsResponse
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return result.Success, nil
}
