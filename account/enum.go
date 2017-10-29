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
