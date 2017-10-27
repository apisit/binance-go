package market

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/apisit/binance-go/client"
)

var httpClient = &http.Client{Timeout: 80 * time.Second}
var baseURL = "https://www.binance.com/api"
var key = ""
var api = client.API{URL: baseURL, Key: key, HTTPClient: httpClient}

func TestMarketPrices(t *testing.T) {
	client := Client{API: api}
	p, err := client.Prices()
	if err != nil {
		t.Error(err)
	}
	log.Printf("%v", p["BNBBTC"])
}

func TestOrderBook(t *testing.T) {
	client := Client{API: api}
	params := OrderBookParams{
		Symbol: "BNBBTC",
	}
	p, err := client.OrderBook(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", p)
}

func TestTwentyFourHourPrice(t *testing.T) {
	client := Client{API: api}
	params := Params{
		Symbol: "BNBBTC",
	}
	p, err := client.TwentyFourHourPrice(params)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%v", p)
}

func TestAllBookTickers(t *testing.T) {
	client := Client{API: api}
	p, err := client.AllBookTickers()
	if err != nil {
		t.Error(err)
	}
	log.Printf("%v", p)
}
