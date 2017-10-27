# Go Binance API [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/apisit/binance-go)

binance-go is a go client library for Binance Rest APIs


## Installation

```sh
go get github.com/apisit/binance-go
```

For details on all the functionality in this library, see the [GoDoc](http://godoc.org/github.com/apisit/binance-go) documentation.

### Usage without a Client
If you are dealing with one account. There is no need to create a new client. you can simply call `binance.$resource$()`


```go
import (
	"github.com/apisit/binance-go"
)

// Setup
binance.APIKey = ""
binance.SecretKey = ""
info, err := binance.Account().Info()
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%v", info)

```

### Usage with a Client
If you are dealing with multiple accounts. You can create a new `binance.Client` by the following

```go
binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")
info, err := binanceClient.Account.Info()
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%v", info)
```


### General endpoints

#### Test connectivity
```go
err := binance.General().Ping()
log.Printf("%v", err)
```

#### Check server time
```go
serverTime := binance.General().ServerTime()
log.Printf("%v", serverTime)
```

### Market Data endpoints

#### Order book
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/market"
)

orderBookParams := market.OrderBookParams{
	Symbol: "BNBBTC",
}

p, err := binance.Market().OrderBook(orderBookParams)
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%v", p)
```
#### 24hr ticker price change statistics

```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/market"
)

params := market.Params{
		Symbol: "BNBBTC",
	}
p, err := binance.Market().TwentyFourHourPrice(params)
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%v", p)
```

#### Symbols price ticker
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/market"
)

p, err := binance.Market().Prices()
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%v", p)
```

#### Symbols order book ticker
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/market"
)

p, err := binance.Market().AllBookTickers()
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%v", p)
```

### Account endpoints
Account requires signed endpoint. You must provide API Key and Secret Key to be able to call these methods.

#### Create Market buy order
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/account"
)

binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")
params := account.MarketBuyParams{
	Params: account.Params{
		Symbol: "BNBBTC",
	},
	Quantity: 14,
}

p, err := binanceClient.Market.MarketBuy(params)
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%+v", p)
```

#### Create Market sell order
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/account"
)

binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")
params := account.MarketBuyParams{
	Params: account.Params{
		Symbol: "BNBBTC",
	},
	Quantity: 14,
}

p, err := binanceClient.Market.MarketSell(params)
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%+v", p)
```

#### Create Limit sell order
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/account"
)

binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")
params := account.SellParams{
	Params: account.Params{
		Symbol: "BNBBTC",
	},
	Price:    0.00025304,
	Quantity: 13,
}

p, err := binanceClient.Market.Sell(params)
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%+v", p)
```

#### Create Limit buy order
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/account"
)

binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")
params := account.BuyParams{
	Params: account.Params{
		Symbol: "BNBBTC",
	},
	Price:    0.00025304,
	Quantity: 13,
}

p, err := binanceClient.Market.Buy(params)
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%+v", p)
```

#### Query order
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/account"
)

binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")
params := account.OrderStatusParams{
	Params: account.Params{
		Symbol: "NEOBTC",
	},
	OrderID: "6793518",
}

p, err := binanceClient.Account.OrderStatus(params)
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%+v", p)
```

#### Cancel order
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/account"
)

binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")
params := account.CancelOrderParams{
	Params: account.Params{
		Symbol: "BNBBTC",
	},
	OrderID: "7897771",
}

p, err := binanceClient.Account.CancelOrder(params)
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%+v", p)
```


#### Current open orders
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/account"
)

binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")
params := account.Params{
	Symbol: "BNBBTC",
}

p, err := binanceClient.Account.OpenOrders(params)
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%+v", p)
```

#### All orders
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/account"
)

binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")
params := account.AllOrdersParams{
	Params: account.Params{
		Symbol: "BNBBTC",
	},
}

p, err := binanceClient.Account.AllOrders(params)
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%+v", p)
```

#### Account information
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/account"
)

binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")
p, err := binanceClient.Account.Info()
if err != nil {
	log.Printf("err = %v", err)
	return
}
log.Printf("%v", p)
```
#### Account trade list
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/account"
)

binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")
params := account.Params{
	Symbol: "BNBBTC",
}

p, err := binanceClient.Account.TradeList(params)
if err != nil {
	t.Error(err)
}
log.Printf("%+v", p)
```

#### Submit a withdraw request
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/account"
)

binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")
params := account.WithdrawParams{
	Asset:   "NEO",
	Address: "AKcm7eABuW1Pjb5HsTwiq7iARSatim9tQ6",
	Amount:  1,
}

p, err := binanceClient.Account.SubmitWithdrawRequest(params)
if err != nil {
	t.Error(err)
}
log.Printf("%+v", p)
```

##### Fetch withdraw history
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/account"
)

binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")

p, err := binanceClient.Account.WithdrawHistory()
if err != nil {
	t.Error(err)
}
log.Printf("%+v", p)
```

##### Fetch deposit history
```go
import (
	"github.com/apisit/binance-go"
	"github.com/apisit/binance-go/account"
)

binanceClient := binance.Client{}
binanceClient.Init("YOUR_API_KEY", "YOUR_SECRET_KEY")

p, err := binanceClient.Account.DepositHistory()
if err != nil {
	t.Error(err)
}
log.Printf("%+v", p)
```