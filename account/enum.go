//ENUM definitions.
package account

type OrderStatusEnum string

//Order status.
const (
	NEW              OrderStatusEnum = "NEW"
	PARTIALLY_FILLED OrderStatusEnum = "PARTIALLY_FILLED"
	FILLED           OrderStatusEnum = "FILLED"
	CANCELED         OrderStatusEnum = "CANCELED"
	PENDING_CANCEL   OrderStatusEnum = "PENDING_CANCEL"
	REJECTED         OrderStatusEnum = "REJECTED"
	EXPIRED          OrderStatusEnum = "EXPIRED"
)

type OrderType string

//Order type.
const (
	LIMIT  OrderType = "LIMIT"
	MARKET OrderType = "MARKET"
)

type OrderSide string

//Order side.
const (
	BUY  OrderSide = "BUY"
	SELL OrderSide = "SELL"
)

type TimeInForce string

//Time in force.
const (
	GTC TimeInForce = "GTC"
	IOC TimeInForce = "IOC"
)

type SymbolType string

//Symbol type.
const (
	SPOT SymbolType = "SPOT"
)

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
