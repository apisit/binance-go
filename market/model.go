package market

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
