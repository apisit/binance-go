package stream

//Aggregate Trade params
type AggregateTradeParams struct {
	Symbol string `json:"symbol"`
}

//Aggregate Trade stream response
type AggregateTradeStream struct {
	EventType         string  `json:"e"`
	EventTime         int64   `json:"E"`
	Symbol            string  `json:"s"`
	AggregatedTradeID int     `json:"a"`
	Price             float64 `json:"p"`
	Quantity          float64 `json:"q"`
	FirtTradeID       int     `json:"f"`
	LastTradeID       int     `json:"l"`
	TradeTime         int64   `json:"T"`
	IsMaker           bool    `json:"m"`
	M                 string  `json:"M"` //can be ignored
}

//Depth params
type DepthParams struct {
	Symbol string `json:"symbol"`
}

//Depth stream response.
type DepthStream struct {
	EventType string          `json:"e"`
	Time      int64           `json:"E"`
	Symbol    string          `json:"s"`
	UpdateID  int             `json:"u"`
	Bids      [][]interface{} `json:"b"`
	Asks      [][]interface{} `json:"a"`
}

type KlineIntervals string

//Kline intervals.
//m -> minutes; h -> hours; d -> days; w -> weeks; M -> months
const (
	OneMinute      KlineIntervals = "1m"
	ThreeMinutes   KlineIntervals = "3m"
	FiveMinutes    KlineIntervals = "5m"
	FifteenMinutes KlineIntervals = "15m"
	ThirtyMinutes  KlineIntervals = "30m"
	OneHour        KlineIntervals = "1h"
	TwoHours       KlineIntervals = "2h"
	FourHours      KlineIntervals = "4h"
	SixHours       KlineIntervals = "6h"
	EightHours     KlineIntervals = "8h"
	TwelveHours    KlineIntervals = "12h"
	OneDay         KlineIntervals = "1d"
	ThreeDays      KlineIntervals = "3d"
	OneWeek        KlineIntervals = "1w"
	OneMonth       KlineIntervals = "1M"
)

//Kline params
type KlineParams struct {
	Symbol   string         `json:"symbol"`
	Interval KlineIntervals `json:"interval"`
}

//Kline stream response
type KlineStream struct {
	EventType string `json:"e"`
	Time      int64  `json:"E"`
	Symbol    string `json:"s"`
	Kline     struct {
		StartTime              int64  `json:"t"`
		Endtime                int64  `json:"T"`
		Symbol                 string `json:"s"`
		Interval               string `json:"i"`
		FirtTradeID            int    `json:"f"`
		LastTradeID            int    `json:"L"`
		Open                   string `json:"o"`
		Close                  string `json:"c"`
		High                   string `json:"h"`
		Low                    string `json:"l"`
		Volume                 string `json:"v"`
		NumberOfTrades         int    `json:"n"`
		IsFinal                bool   `json:"x"`
		QuoteVolume            string `json:"q"`
		VolumeOfAciveBuy       string `json:"V"`
		QuoteVolumeOfActiveBuy string `json:"Q"`
		B                      string `json:"B"` //can be ignored
	} `json:"k"`
}
