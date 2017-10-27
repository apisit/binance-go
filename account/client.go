//Package account provides the binding for Binance Rest APIs account endpoints
package account

import (
	"github.com/apisit/binance-go/client"
)

//Default RecvWindow
//more at https://www.binance.com/restapipub.html#user-content-signed-endpoint-security
//Serious trading is about timing. Networks can be unstable and unreliable,
//which can lead to requests taking varying amounts of time to reach the servers.
//With recvWindow, you can specify that the request must be processed within
//a certain number of milliseconds or be rejected by the server.

const (
	DefaultRecvWindow int64 = 6000000
)

type Client struct {
	API client.API
}

//Methods for account endpoints
type Interface interface {
	Info() (AccountInfo, error)
	AllOrders(params AllOrdersParams) (AllOrders, error)
	OpenOrders(params Params) (OpenOrders, error)
	OrderStatus(params OrderStatusParams) (OrderStatus, error)
	CancelOrder(params CancelOrderParams) (CancelOrder, error)
	Buy(params BuyParams) (OrderResponse, error)
	Sell(params SellParams) (OrderResponse, error)
	MarketBuy(params MarketBuyParams) (OrderResponse, error)
	MarketSell(params MarketSellParams) (OrderResponse, error)
	TradeList(params Params) (TradeList, error)

	DepositHistory() (DepositHistory, error)
	WithdrawHistory() (WithdrawHistory, error)
	SubmitWithdrawRequest(params WithdrawParams) (WithdrawResponse, error)
}

var _ Interface = (*Client)(nil)

//Get current account information.
func (c *Client) Info() (AccountInfo, error) {
	out := AccountInfo{}
	err := c.API.SignedRequest("GET", "/api/v3/account", nil, &out)
	return out, err
}

//Get all account orders; active, canceled, or filled.
func (c *Client) AllOrders(params AllOrdersParams) (AllOrders, error) {
	out := AllOrders{}
	err := c.API.SignedRequest("GET", "/api/v3/allOrders", params, &out)
	return out, err
}

//Get all open orders on a symbol.
func (c *Client) OpenOrders(params Params) (OpenOrders, error) {
	out := OpenOrders{}
	err := c.API.SignedRequest("GET", "/api/v3/openOrders", params, &out)
	return out, err
}

//Check an order's status.
func (c *Client) OrderStatus(params OrderStatusParams) (OrderStatus, error) {
	out := OrderStatus{}
	err := c.API.SignedRequest("GET", "/api/v3/order", params, &out)
	return out, err
}

//Cancel an active order.
func (c *Client) CancelOrder(params CancelOrderParams) (CancelOrder, error) {
	out := CancelOrder{}
	err := c.API.SignedRequest("DELETE", "/api/v3/order", params, &out)
	return out, err
}

//Send in a new market buy order
func (c *Client) MarketBuy(params MarketBuyParams) (OrderResponse, error) {
	out := OrderResponse{}
	orderParams := OrderParams{
		Symbol:   params.Symbol,
		Quantity: params.Quantity,
		Type:     MARKET,
		Side:     BUY,
	}
	err := c.API.SignedRequest("POST", "/api/v3/order", orderParams, &out)
	return out, err
}

//Send in a new market sell order
func (c *Client) MarketSell(params MarketSellParams) (OrderResponse, error) {
	out := OrderResponse{}
	orderParams := OrderParams{
		Symbol:   params.Symbol,
		Quantity: params.Quantity,
		Type:     MARKET,
		Side:     SELL,
	}
	err := c.API.SignedRequest("POST", "/api/v3/order", orderParams, &out)
	return out, err
}

//Send in a new limit buy order
func (c *Client) Buy(params BuyParams) (OrderResponse, error) {
	out := OrderResponse{}
	orderParams := OrderParams{
		Symbol:      params.Symbol,
		Quantity:    params.Quantity,
		Price:       params.Price,
		Type:        LIMIT,
		Side:        BUY,
		TimeInForce: GTC,
	}
	err := c.API.SignedRequest("POST", "/api/v3/order", orderParams, &out)
	return out, err
}

//Send in a new limit sell order
func (c *Client) Sell(params SellParams) (OrderResponse, error) {
	out := OrderResponse{}
	orderParams := OrderParams{
		Symbol:      params.Symbol,
		Quantity:    params.Quantity,
		Price:       params.Price,
		Type:        LIMIT,
		Side:        SELL,
		TimeInForce: GTC,
	}
	err := c.API.SignedRequest("POST", "/api/v3/order", orderParams, &out)
	return out, err
}

//Get trades for a specific account and symbol.
func (c *Client) TradeList(params Params) (TradeList, error) {
	out := TradeList{}
	err := c.API.SignedRequest("GET", "/api/v3/myTrades", params, &out)
	return out, err
}

//Fetch deposit history.
func (c *Client) DepositHistory() (DepositHistory, error) {
	out := DepositHistory{}
	err := c.API.SignedRequest("POST", "/wapi/v1/getDepositHistory.html", nil, &out)
	return out, err
}

//Fetch withdraw history.
func (c *Client) WithdrawHistory() (WithdrawHistory, error) {
	out := WithdrawHistory{}
	err := c.API.SignedRequest("POST", "/wapi/v1/getWithdrawHistory.html", nil, &out)
	return out, err
}

//Submit a withdraw request.
func (c *Client) SubmitWithdrawRequest(params WithdrawParams) (WithdrawResponse, error) {
	out := WithdrawResponse{}
	err := c.API.SignedRequest("POST", "/wapi/v1/withdraw.html", params, &out)
	return out, err
}
