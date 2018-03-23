package market

import (
	"encoding/json"
	"fmt"
)

type Params struct {
	Symbol string `json:"symbol"`
}

type Price struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type OrderBook struct {
	LastUpdateID int             `json:"lastUpdateId"`
	Bids         [][]interface{} `json:"bids"`
	Asks         [][]interface{} `json:"asks"`
}

type OrderBookParams struct {
	Symbol string `json:"symbol"`
	Limit  int    `json:"limit,omitempty"` //Default 100; max 100.
}

type TwentyFourHourPrice struct {
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FristID            int    `json:"fristId"`
	LastID             int    `json:"lastId"`
	Count              int    `json:"count"`
}

type AllBookTickers []struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

type KLine struct {
	OpenTime              int64
	Open                  string
	High                  string
	Low                   string
	Close                 string
	Volume                string
	CloseTime             int64
	QuoteAssetVol         string
	NoTrades              int
	TakerBuyBaseAssetVol  string
	TakerBuyQuoteAssetVol string
	Ignore                string
}

type KLineParams struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`        //Possible values: 1m,3m,5m,15m,30m,1h,2h,4h,6h,8h,12h,1d,3d,1w,1M
	Limit     int    `json:"limit,omitempty"` //Default 500; max 500.
	StartTime int64  `json:"startTime,omitempty"`
	EndTime   int64  `json:"endTime,omitempty"`
}

func (k *KLine) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&k.OpenTime, &k.Open, &k.High, &k.Low, &k.Close, &k.Volume, &k.CloseTime, &k.QuoteAssetVol, &k.NoTrades, &k.TakerBuyBaseAssetVol, &k.TakerBuyQuoteAssetVol, &k.Ignore}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in KLine: %d != %d", g, e)
	}
	return nil
}
