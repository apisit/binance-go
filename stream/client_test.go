package stream

import (
	"log"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/apisit/binance-go/client"
)

var httpClient = &http.Client{Timeout: 80 * time.Second}
var baseURL = "https://www.binance.com/api"
var key = ""
var api = client.API{URL: baseURL, Key: key, HTTPClient: httpClient}

func TestAggregateTradeStream(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	client := Client{API: api}
	p := AggregateTradeParams{
		Symbol: "BNBBTC",
	}
	client.AggregateTrade(p, func(d AggregateTradeStream) {
		log.Printf("%v", d)
		wg.Done()
	})
	wg.Wait()
}

func TestDepthStream(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	client := Client{API: api}
	p := DepthParams{
		Symbol: "BNBBTC",
	}
	client.Depth(p, func(d DepthStream) {
		log.Printf("%v", d)
		wg.Done()
	})
	wg.Wait()
}

func TestKlineStream(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	client := Client{API: api}
	p := KlineParams{
		Symbol:   "BNBBTC",
		Interval: OneMinute,
	}
	client.Kline(p, func(d KlineStream) {
		log.Printf("%v", d)
		wg.Done()
	})
	wg.Wait()
}
