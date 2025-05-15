package services

import (
	"encoding/json"
	"fmt"

	"github.com/Kora1128/icici-breezeconnect-go/breezeconnect"
	"github.com/Kora1128/icici-breezeconnect-go/breezeconnect/models"
)

// FundsService handles funds-related API calls
type FundsService struct {
	client breezeconnect.ClientInterface
}

// NewFundsService creates a new funds service
func NewFundsService(client breezeconnect.ClientInterface) *FundsService {
	return &FundsService{client: client}
}

// GetFunds retrieves the available funds
func (s *FundsService) GetFunds() (*models.Funds, error) {
	response, err := s.client.MakeRequest("GET", "/funds", nil)
	if err != nil {
		return nil, fmt.Errorf("error getting funds: %v", err)
	}

	var result models.Funds
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &result, nil
}
