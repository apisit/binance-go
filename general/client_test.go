package general

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

func TestPing(t *testing.T) {

	client := Client{API: api}
	err := client.Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestServerTime(t *testing.T) {
	client := Client{API: api}
	time := client.ServerTime()
	log.Printf("server time = %v", time)
}
