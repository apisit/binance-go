//Package binance provides the binding for Binance Rest APIs
package binance

import (
	"net/http"
	"time"

	"github.com/apisit/binance-go/account"
	"github.com/apisit/binance-go/client"
	"github.com/apisit/binance-go/general"
	"github.com/apisit/binance-go/market"
)

//constant used for API client
const (
	APIURL             = "https://www.binance.com/api"
	clientVersion      = "0.1"
	UserAgent          = "github.com/apisit/binance-go version/" + clientVersion
	defaultHTTPTimeout = 80 * time.Second
)

//Default HTTP client
var httpClient = &http.Client{Timeout: defaultHTTPTimeout}

//Binance API Key
var APIKey string
var SecretKey string

//Use Market endpoint without client
func Market() *market.Client {
	api := client.New(APIURL, APIKey, SecretKey, httpClient, UserAgent)
	return &market.Client{*api}
}

//Use General endpoint without client
func General() *general.Client {
	api := client.New(APIURL, APIKey, SecretKey, httpClient, UserAgent)
	return &general.Client{*api}
}

//Use Account endpoint without client
func Account() *account.Client {
	api := client.New(APIURL, APIKey, SecretKey, httpClient, UserAgent)
	return &account.Client{*api}
}

//Client is the Binance client. It contains all resources available.
type Client struct {
	Market  market.Client
	General general.Client
	Account account.Client
}

//Init initializes the Binance client with given API key, secret key.
func (c *Client) Init(apiKey, secretKey string) {
	api := client.New(APIURL, apiKey, secretKey, httpClient, UserAgent)
	c.Market = market.Client{*api}
	c.General = general.Client{*api}
	c.Account = account.Client{*api}
}
