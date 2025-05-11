package breezeconnect

// PortfolioHolding represents a single holding in the portfolio
type PortfolioHolding struct {
	Symbol          string  `json:"symbol"`
	Quantity        int     `json:"quantity"`
	AveragePrice    float64 `json:"average_price"`
	LastTradedPrice float64 `json:"last_traded_price"`
	PnL             float64 `json:"pnl"`
	ProductType     string  `json:"product_type"`
	Exchange        string  `json:"exchange"`
}

// Position represents a single position
type Position struct {
	Symbol          string  `json:"symbol"`
	Quantity        int     `json:"quantity"`
	AveragePrice    float64 `json:"average_price"`
	LastTradedPrice float64 `json:"last_traded_price"`
	PnL             float64 `json:"pnl"`
	ProductType     string  `json:"product_type"`
	Exchange        string  `json:"exchange"`
	ExpiryDate      string  `json:"expiry_date,omitempty"`
	StrikePrice     float64 `json:"strike_price,omitempty"`
	OptionType      string  `json:"option_type,omitempty"`
}

// APIResponse represents the standard API response structure
type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
