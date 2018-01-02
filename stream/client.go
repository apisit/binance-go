package stream

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/apisit/binance-go/client"
)

type DepthHandler func(data DepthStream)
type KlineHandler func(data KlineStream)
type TradesHandler func(data TradesStream)

type Client struct {
	API client.API
}

//Methods for stream endpoints
type Interface interface {
	Depth(params DepthParams, handler DepthHandler)
	Kline(params KlineParams, handler KlineHandler)
	Trades(params TradesParams, handler TradesHandler)
}

var _ Interface = (*Client)(nil)

//Stream for depth endpoint
func (c *Client) Depth(params DepthParams, handler DepthHandler) {
	endpoint := fmt.Sprintf("%s@depth", strings.ToLower(params.Symbol))

	c.API.Stream(endpoint, func(d []byte) {
		out := DepthStream{}
		json.Unmarshal(d, &out)
		go handler(out)
	})
}

//Stream for kline endpoint
func (c *Client) Kline(params KlineParams, handler KlineHandler) {
	endpoint := fmt.Sprintf("%s@kline_%s", strings.ToLower(params.Symbol), params.Interval)
	c.API.Stream(endpoint, func(d []byte) {
		out := KlineStream{}
		json.Unmarshal(d, &out)
		go handler(out)
	})
}

//Stream for trades endpoint
func (c *Client) Trades(params TradesParams, handler TradesHandler) {
	endpoint := fmt.Sprintf("%s@aggTrade", strings.ToLower(params.Symbol))

	c.API.Stream(endpoint, func(d []byte) {
		out := TradesStream{}
		json.Unmarshal(d, &out)
		go handler(out)
	})
}
