package account

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/apisit/binance-go/client"
)

var httpClient = &http.Client{Timeout: 80 * time.Second}
var baseURL = "https://www.binance.com"
var key = ""
var secret = ""
var api = client.API{URL: baseURL, Key: key, SecretKey: secret, HTTPClient: httpClient}

func TestAccountInfo(t *testing.T) {
	client := Client{API: api}
	p, err := client.Info()
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}

func TestAllOrders(t *testing.T) {
	client := Client{API: api}
	params := AllOrdersParams{
		Params: Params{
			Symbol: "BNBBTC",
		},
	}
	p, err := client.AllOrders(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}

func TestOpenOrders(t *testing.T) {
	client := Client{API: api}
	params := Params{
		Symbol: "BNBBTC",
	}

	p, err := client.OpenOrders(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}

func TestOrderStatus(t *testing.T) {
	client := Client{API: api}
	params := OrderStatusParams{
		Params: Params{
			Symbol: "NEOBTC",
		},
		OrderID: "6793518",
	}

	p, err := client.OrderStatus(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}

func TestCancelOrder(t *testing.T) {
	client := Client{API: api}
	params := CancelOrderParams{
		Params: Params{
			Symbol: "BNBBTC",
		},
		OrderID: "7897771",
	}

	p, err := client.CancelOrder(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}

func TestMarketBuy(t *testing.T) {
	client := Client{API: api}
	params := MarketBuyParams{
		Params: Params{
			Symbol: "BNBBTC",
		},
		Quantity: 14,
	}

	p, err := client.MarketBuy(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}

func TestMarketSell(t *testing.T) {
	client := Client{API: api}
	params := MarketSellParams{
		Params: Params{
			Symbol: "BNBBTC",
		},
		Quantity: 13,
	}

	p, err := client.MarketSell(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}

func TestBuy(t *testing.T) {
	client := Client{API: api}
	params := BuyParams{
		Params: Params{
			Symbol: "BNBBTC",
		},
		Price:    0.00021304,
		Quantity: 14,
	}

	p, err := client.Buy(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}

func TestSell(t *testing.T) {
	client := Client{API: api}
	params := SellParams{
		Params: Params{
			Symbol: "BNBBTC",
		},
		Price:    0.00025304,
		Quantity: 13,
	}

	p, err := client.Sell(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}

func TestTradeList(t *testing.T) {
	client := Client{API: api}
	params := Params{
		Symbol: "BNBBTC",
	}

	p, err := client.TradeList(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}

func TestDepositHistory(t *testing.T) {
	client := Client{API: api}
	p, err := client.DepositHistory()
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}

func TestWithdrawHistory(t *testing.T) {
	client := Client{API: api}
	p, err := client.WithdrawHistory()
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}

func TestWithDraw(t *testing.T) {
	client := Client{API: api}
	params := WithdrawParams{
		Asset:   "NEO",
		Address: "AKcm7eABuW1Pjb5HsTwiq7iARSatim9tQ6",
		Amount:  1,
	}
	p, err := client.SubmitWithdrawRequest(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}
