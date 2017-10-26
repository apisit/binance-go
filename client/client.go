//Package client contains methods to make request to Binance API server.
package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

//API is a Binance API client.
type API struct {
	URL        string
	Key        string
	SecretKey  string
	HTTPClient *http.Client
	UserAgent  string
}

//New initializes API with given URL, api key and secret key. it also provides a way to overwrite *http.Client
func New(url, key, secretKey string, httpClient *http.Client, userAgent string) *API {
	return &API{
		URL:        url,
		Key:        key,
		SecretKey:  secretKey,
		HTTPClient: httpClient,
		UserAgent:  userAgent,
	}
}

//Making a public request to Binance API server.
func (a *API) Request(method, endpoint string, params interface{}, out interface{}) error {
	url, _ := url.ParseRequestURI(a.URL)
	url.Path = url.Path + endpoint

	if method == "GET" {
		//parse params to query string
		b, _ := json.Marshal(params)
		m := map[string]interface{}{}
		json.Unmarshal(b, &m)
		q := url.Query()
		for k, v := range m {
			q.Set(k, fmt.Sprintf("%v", v))
		}
		url.RawQuery = q.Encode()
	}
	log.Printf("%v %v", method, url.String())
	req, _ := http.NewRequest(method, url.String(), nil)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-MBX-APIKEY", a.Key)
	req.Header.Add("UserAgent", a.UserAgent)
	res, err := a.HTTPClient.Do(req)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		type binanceError struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}
		e := binanceError{}
		err = json.NewDecoder(res.Body).Decode(&e)
		return errors.New(e.Msg)
	}

	if out != nil {
		err = json.NewDecoder(res.Body).Decode(&out)
	}

	return err
}

//Making a signed request to Binance API server.
func (a *API) SignedRequest(method, endpoint string, params interface{}, out interface{}) error {
	url, _ := url.ParseRequestURI(a.URL)
	url.Path = url.Path + endpoint

	//parse params to query string
	b, _ := json.Marshal(params)
	m := map[string]interface{}{}
	json.Unmarshal(b, &m)

	q := url.Query()
	for k, v := range m {
		q.Set(k, fmt.Sprintf("%v", v))
	}

	//timestamp is mandatory in signed request
	q.Add("timestamp", fmt.Sprintf("%v", time.Now().Unix()*1000))

	mac := hmac.New(sha256.New, []byte(a.SecretKey))
	mac.Write([]byte(q.Encode()))
	expectedMAC := mac.Sum(nil)
	signed := hex.EncodeToString(expectedMAC)
	//signature needs to be at the last param
	url.RawQuery = q.Encode() + "&signature=" + signed

	log.Printf("%v %v", method, url.String())

	req, _ := http.NewRequest(method, url.String(), nil)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-MBX-APIKEY", a.Key)
	req.Header.Add("UserAgent", a.UserAgent)
	res, err := a.HTTPClient.Do(req)

	defer res.Body.Close()
	if res.StatusCode != 200 {
		type binanceError struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}
		e := binanceError{}
		err = json.NewDecoder(res.Body).Decode(&e)
		return errors.New(e.Msg)
	}
	defer res.Body.Close()
	if out != nil {
		err = json.NewDecoder(res.Body).Decode(&out)
	}
	return err
}
