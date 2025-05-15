package models

// DematHolding represents a single demat holding
type DematHolding struct {
	ISIN            string  `json:"isin"`
	Symbol          string  `json:"symbol"`
	Quantity        int     `json:"quantity"`
	ProductType     string  `json:"product_type"`
	Exchange        string  `json:"exchange"`
	AveragePrice    float64 `json:"average_price"`
	LastTradedPrice float64 `json:"last_traded_price"`
	PnL             float64 `json:"pnl"`
	TotalValue      float64 `json:"total_value"`
}

// DematHoldingsResponse represents the response for demat holdings
type DematHoldingsResponse struct {
	Success []DematHolding `json:"Success"`
	Status  int            `json:"Status"`
	Error   string         `json:"Error"`
}
