package services

import (
	"encoding/json"
	"fmt"

	"github.com/Kora1128/icici-breezeconnect-go/breezeconnect"
	"github.com/Kora1128/icici-breezeconnect-go/breezeconnect/models"
)

// CustomerService handles customer-related API calls
type CustomerService struct {
	client breezeconnect.ClientInterface
}

// NewCustomerService creates a new customer service
func NewCustomerService(client breezeconnect.ClientInterface) *CustomerService {
	return &CustomerService{client: client}
}

// GetCustomerDetails retrieves customer details and session token
func (s *CustomerService) GetCustomerDetails(sessionToken string) (*models.CustomerDetails, error) {
	payload := map[string]string{
		"SessionToken": sessionToken,
		"AppKey":       s.client.GetAPIKey(),
	}

	response, err := s.client.MakeRequest("GET", "/customerdetails", payload)
	if err != nil {
		return nil, fmt.Errorf("error getting customer details: %v", err)
	}

	var result models.CustomerDetails
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	if result.Success.SessionToken != "" {
		s.client.SetSessionKey(result.Success.SessionToken)
	}

	return &result, nil
}
