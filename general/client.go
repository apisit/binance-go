//Package general provides the binding for Binance Rest APIs general endpoints
package general

import (
	"time"

	"github.com/apisit/binance-go/client"
)

type Client struct {
	API client.API
}

//Methods for general endpoints
type Interface interface {
	Ping() error
	ServerTime() time.Time
}

var _ Interface = (*Client)(nil)

//Ping binance server
func (c *Client) Ping() error {
	err := c.API.Request("GET", "/v1/ping", nil, nil)
	return err
}

//Binance server time
func (c *Client) ServerTime() time.Time {
	result := struct {
		Time int64 `json:"serverTime"`
	}{}
	c.API.Request("GET", "/v1/time", nil, &result)
	//time returns from server is milliseconds
	return time.Unix(result.Time/1000, 0)
}
