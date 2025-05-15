package models

// Funds represents the available funds
type Funds struct {
	Success struct {
		AvailableBalance float64 `json:"available_balance"`
	} `json:"Success"`
	Status int    `json:"Status"`
	Error  string `json:"Error"`
}
