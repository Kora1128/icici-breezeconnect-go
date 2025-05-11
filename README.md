# ICICI BreezeConnect Go SDK

This is a Go SDK for the ICICI Direct Breeze API, providing a simple interface to interact with ICICI Direct's trading platform.

## Installation

```bash
go get github.com/kowshikr/icici-breezeconnect-go
```

## Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/kowshikr/icici-breezeconnect-go/breezeconnect"
)

func main() {
    // Initialize the client
    client := breezeconnect.NewClient("your_api_key", "your_api_secret")

    // Generate session
    if err := client.GenerateSession(); err != nil {
        log.Fatalf("Error generating session: %v", err)
    }

    // Get portfolio holdings
    holdings, err := client.GetPortfolioHoldings()
    if err != nil {
        log.Fatalf("Error getting portfolio holdings: %v", err)
    }

    // Process holdings...
}
```

## Features

- Authentication and session management
- Portfolio holdings retrieval
- Position tracking
- Error handling and response parsing

## API Documentation

For detailed API documentation, please refer to the [ICICI Direct Breeze API Documentation](https://api.icicidirect.com/breezeapi/documents/index.html).

## License

MIT License