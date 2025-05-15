package models

// PortfolioHolding represents a single holding in the portfolio
type PortfolioHolding struct {
	Symbol          string  `json:"symbol"`
	Quantity        int     `json:"quantity"`
	AveragePrice    float64 `json:"average_price"`
	LastTradedPrice float64 `json:"last_traded_price"`
	PnL             float64 `json:"pnl"`
	ProductType     string  `json:"product_type"`
	Exchange        string  `json:"exchange"`
	ISIN            string  `json:"isin"`
	Name            string  `json:"name"`
	TotalValue      float64 `json:"total_value"`
	FreeQuantity    int     `json:"free_quantity"`
	LockedQuantity  int     `json:"locked_quantity"`
}

// PortfolioHoldingsResponse represents the response for portfolio holdings
type PortfolioHoldingsResponse struct {
	Success []PortfolioHolding `json:"Success"`
	Status  int                `json:"Status"`
	Error   string             `json:"Error"`
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
	ExpiryDate      string  `json:"expiry_date"`
	StrikePrice     float64 `json:"strike_price"`
	OptionType      string  `json:"option_type"`
	TotalValue      float64 `json:"total_value"`
	FreeQuantity    int     `json:"free_quantity"`
	LockedQuantity  int     `json:"locked_quantity"`
}

// PositionsResponse represents the response for positions
type PositionsResponse struct {
	Success []Position `json:"Success"`
	Status  int        `json:"Status"`
	Error   string     `json:"Error"`
}
