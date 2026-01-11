# Binance Alpha Connector Go

This is the Go client library for the Binance Alpha REST API.

## Installation

```bash
go get github.com/binance/binance-connector-go/clients/alpha
```

## Usage

### REST API

```go
package main

import (
    "context"
    "encoding/json"
    "log"

    client "github.com/binance/binance-connector-go/clients/alpha"
    "github.com/binance/binance-connector-go/common/common"
)

func main() {
    configuration := common.NewConfigurationRestAPI(
        common.WithBasePath("https://www.binance.com"),
    )
    apiClient := client.NewBinanceAlphaClient(
        client.WithRestAPI(configuration),
    )

    // Get 24hr Ticker Price Change Statistics
    resp, err := apiClient.RestApi.AlphaAPI.GetTicker(context.Background()).Execute()
    if err != nil {
        log.Println(err)
        return
    }

    dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
    log.Printf("Response: %s\n", string(dataValue))

    // Get Klines
    klinesResp, err := apiClient.RestApi.AlphaAPI.GetKlines(context.Background()).
        Symbol("ALPHA_175USDT").
        Interval("1h").
        Execute()
    if err != nil {
        log.Println(err)
        return
    }

    klinesData, _ := json.MarshalIndent(klinesResp.Data, "", "  ")
    log.Printf("Klines: %s\n", string(klinesData))

    // Get Token List
    tokensResp, err := apiClient.RestApi.AlphaAPI.GetTokens(context.Background()).Execute()
    if err != nil {
        log.Println(err)
        return
    }

    tokensData, _ := json.MarshalIndent(tokensResp.Data, "", "  ")
    log.Printf("Tokens: %s\n", string(tokensData))
}
```

## Available APIs

### AlphaAPI

| Method | Description |
|--------|-------------|
| `GetTicker` | Get 24hr ticker price change statistics |
| `GetKlines` | Get Kline/candlestick data |
| `GetTokens` | Get all available Alpha tokens |

## License

MIT
