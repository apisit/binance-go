//Package general provides the binding for Binance Rest APIs market endpoints
package market

import "github.com/apisit/binance-go/client"

type Client struct {
	API client.API
}

//Methods for market endpoints
type MarketInterface interface {
	Prices() (Prices, error)
	OrderBook(params OrderBookParams) (OrderBook, error)
	TwentyFourHourPrice(params Params) (TwentyFourHourPrice, error)
	AllBookTickers() (AllBookTickers, error)
}

var _ MarketInterface = (*Client)(nil)

type Prices map[string]Price

//Latest price for all symbols.
func (c *Client) Prices() (Prices, error) {
	priceList := []Price{}
	err := c.API.Request("GET", "/v1/ticker/allPrices", nil, &priceList)

	prices := Prices{}
	for _, p := range priceList {
		prices[p.Symbol] = p
	}

	return prices, err
}

//Order book.
func (c *Client) OrderBook(params OrderBookParams) (OrderBook, error) {
	orderbook := OrderBook{}
	err := c.API.Request("GET", "/v1/depth", params, &orderbook)
	return orderbook, err
}

//24 hour price change statistics.
func (c *Client) TwentyFourHourPrice(params Params) (TwentyFourHourPrice, error) {
	out := TwentyFourHourPrice{}
	err := c.API.Request("GET", "/v1/ticker/24hr", params, &out)
	return out, err
}

//Best price/qty on the order book for all symbols.
func (c *Client) AllBookTickers() (AllBookTickers, error) {
	out := AllBookTickers{}
	err := c.API.Request("GET", "/v1/ticker/allBookTickers", nil, &out)
	return out, err
}
