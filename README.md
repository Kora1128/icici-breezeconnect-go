# ICICI Direct BreezeConnect Go SDK

A Go SDK for interacting with the ICICI Direct BreezeConnect API. This SDK provides a simple and type-safe way to access various trading and market data endpoints.

## Features

- Authentication and session management
- Funds and balance information
- Demat holdings
- Portfolio holdings and positions
- Customer details
- Type-safe request/response models
- Comprehensive test coverage

## Installation

```bash
go get github.com/Kora1128/icici-breezeconnect-go
```

## Usage

### Initialization

```go
import "github.com/Kora1128/icici-breezeconnect-go/breezeconnect"

// Create a new client
client := breezeconnect.NewClient("your_api_key", "your_api_secret")

// Get customer details and session token
customerService := services.NewCustomerService(client)
details, err := customerService.GetCustomerDetails("your_session_token")
if err != nil {
    log.Fatal(err)
}

// The session token is automatically set in the client
```

### Available Services

#### Funds Service

Get available balance and funds information:

```go
fundsService := services.NewFundsService(client)
funds, err := fundsService.GetFunds()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Available Balance: %.2f\n", funds.Success.AvailableBalance)
```

#### Demat Service

Get demat holdings:

```go
dematService := services.NewDematService(client)
holdings, err := dematService.GetDematHoldings()
if err != nil {
    log.Fatal(err)
}

for _, holding := range holdings {
    fmt.Printf("Symbol: %s, Quantity: %d, Value: %.2f\n",
        holding.Symbol,
        holding.Quantity,
        holding.TotalValue)
}
```

#### Portfolio Service

Get portfolio holdings and positions:

```go
portfolioService := services.NewPortfolioService(client)

// Get portfolio holdings
holdings, err := portfolioService.GetPortfolioHoldings()
if err != nil {
    log.Fatal(err)
}

// Get current positions
positions, err := portfolioService.GetPositions()
if err != nil {
    log.Fatal(err)
}
```

#### Customer Service

Get customer details and manage session:

```go
customerService := services.NewCustomerService(client)
details, err := customerService.GetCustomerDetails("your_session_token")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("User ID: %s\n", details.Success.IDirectUserID)
fmt.Printf("User Name: %s\n", details.Success.IDirectUserName)
```

## Response Models

The SDK provides type-safe models for all API responses:

### Funds
```go
type Funds struct {
    Success struct {
        AvailableBalance float64 `json:"available_balance"`
    } `json:"Success"`
    Status int    `json:"Status"`
    Error  string `json:"Error"`
}
```

### Demat Holdings
```go
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
```

### Portfolio Holdings
```go
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
```

### Positions
```go
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
```

## Error Handling

The SDK provides detailed error messages for API errors. All service methods return errors that can be checked:

```go
funds, err := fundsService.GetFunds()
if err != nil {
    // Handle error
    fmt.Printf("Error getting funds: %v\n", err)
    return
}
```

## Testing

The SDK includes comprehensive tests using a mock client. To run the tests:

```bash
go test -v ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.