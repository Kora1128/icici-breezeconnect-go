package services

import (
	"testing"

	"github.com/kowshikr/icici-breezeconnect-go/breezeconnect"
)

func TestNewPortfolioService(t *testing.T) {
	client := breezeconnect.NewClient("test_api_key", "test_api_secret")
	service := NewPortfolioService(client)

	if service.client != client {
		t.Error("Expected service.client to be the same as the input client")
	}
}
