package stream

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/apisit/binance-go/client"
)

type DepthHandler func(data DepthStream)
type KlineHandler func(data KlineStream)
type AggregateTradeHandler func(data AggregateTradeStream)

type Client struct {
	API client.API
}

//Methods for stream endpoints
type Interface interface {
	AggregateTrade(params AggregateTradeParams, handler AggregateTradeHandler)
	Depth(params DepthParams, handler DepthHandler)
	Kline(params KlineParams, handler KlineHandler)
}

var _ Interface = (*Client)(nil)

//Stream for aggregate trade endpoint
func (c *Client) AggregateTrade(params AggregateTradeParams, handler AggregateTradeHandler) {
	endpoint := fmt.Sprintf("%s@aggTrade", strings.ToLower(params.Symbol))

	c.API.Stream(endpoint, func(d []byte) {
		out := AggregateTradeStream{}
		json.Unmarshal(d, &out)
		go handler(out)
	})
}

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
