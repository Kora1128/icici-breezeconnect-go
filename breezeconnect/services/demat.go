package services

import (
	"encoding/json"
	"fmt"

	"github.com/Kora1128/icici-breezeconnect-go/breezeconnect"
	"github.com/Kora1128/icici-breezeconnect-go/breezeconnect/models"
)

// DematService handles demat-related API calls
type DematService struct {
	client breezeconnect.ClientInterface
}

// NewDematService creates a new demat service
func NewDematService(client breezeconnect.ClientInterface) *DematService {
	return &DematService{client: client}
}

// GetDematHoldings retrieves the demat holdings
func (s *DematService) GetDematHoldings() ([]models.DematHolding, error) {
	response, err := s.client.MakeRequest("GET", "/demat/holdings", nil)
	if err != nil {
		return nil, fmt.Errorf("error getting demat holdings: %v", err)
	}

	var result models.DematHoldingsResponse
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return result.Success, nil
}
