package account

//Account information response.
type AccountInfo struct {
	MakerCommission  int  `json:"makerCommission"`
	TakerCommission  int  `json:"takerCommission"`
	BuyerCommission  int  `json:"buyerCommission"`
	SellerCommission int  `json:"sellerCommission"`
	CanTrade         bool `json:"canTrade"`
	CanWithdraw      bool `json:"canWithdraw"`
	CanDeposit       bool `json:"canDeposit"`
	Balances         []struct {
		Asset  string `json:"asset"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
	} `json:"balances"`
}

//Common properties.
type Params struct {
	Symbol     string `json:"symbol"`
	RecvWindow int64  `json:"recvWindow,omitempty"`
}

//All order params and response.
type (
	AllOrdersParams struct {
		Params
		OrderID string `json:"orderId,omitempty"`
		Limit   int    `json:"limit,omitempty"` //Default 500; max 500.
	}

	AllOrders []struct {
		Symbol        string `json:"symbol"`
		OrderID       int    `json:"orderId"`
		ClientOrderID string `json:"clientOrderId"`
		Price         string `json:"price"`
		OrigQty       string `json:"origQty"`
		ExecutedQty   string `json:"executedQty"`
		Status        string `json:"status"`
		TimeInForce   string `json:"timeInForce"`
		Type          string `json:"type"`
		Side          string `json:"side"`
		StopPrice     string `json:"stopPrice"`
		IcebergQty    string `json:"icebergQty"`
		Time          int64  `json:"time"`
	}
)

//Open order response.
type OpenOrders []struct {
	Symbol        string `json:"symbol"`
	OrderID       int    `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
	Price         string `json:"price"`
	OrigQty       string `json:"origQty"`
	ExecutedQty   string `json:"executedQty"`
	Status        string `json:"status"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	Side          string `json:"side"`
	StopPrice     string `json:"stopPrice"`
	IcebergQty    string `json:"icebergQty"`
	Time          int64  `json:"time"`
}

//Order status params.
type (
	OrderStatusParams struct {
		Params
		OrderID string `json:"orderId,omitempty"`
	}

	OrderStatus struct {
		Symbol        string `json:"symbol"`
		OrderID       int    `json:"orderId"`
		ClientOrderID string `json:"clientOrderId"`
		Price         string `json:"price"`
		OrigQty       string `json:"origQty"`
		ExecutedQty   string `json:"executedQty"`
		Status        string `json:"status"`
		TimeInForce   string `json:"timeInForce"`
		Type          string `json:"type"`
		Side          string `json:"side"`
		StopPrice     string `json:"stopPrice"`
		IcebergQty    string `json:"icebergQty"`
		Time          int64  `json:"time"`
	}
)

//Cancel order params.
type (
	CancelOrderParams struct {
		Params
		OrderID string `json:"orderId,omitempty"`
	}

	CancelOrder struct {
		Symbol            string `json:"symbol"`
		OrigClientOrderID string `json:"origClientOrderId"`
		OrderID           int    `json:"orderId"`
		ClientOrderID     string `json:"clientOrderId"`
	}
)

//Order params.
type OrderParams struct {
	Symbol      string      `json:"symbol"`
	Side        OrderSide   `json:"side"`
	Type        OrderType   `json:"type"`
	TimeInForce TimeInForce `json:"timeInForce,omitempty"`
	Quantity    float64     `json:"quantity"`
	Price       float64     `json:"price,omitempty"`
	StopPrice   float64     `json:"stopPrice,omitempty"`
	IcebergQty  float64     `json:"icebergQty,omitempty"`
	RecvWindow  int64       `json:"recvWindow,omitempty"`
}

//Common order response.
type OrderResponse struct {
	Symbol        string `json:"symbol"`
	OrderID       int    `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
	TransactTime  int64  `json:"transactTime"`
}

//Buy params.
type (
	BuyParams struct {
		Params
		Price    float64 `json:"price"`
		Quantity float64 `json:"quantity"`
	}
	SellParams struct {
		Params
		Price    float64 `json:"price"`
		Quantity float64 `json:"quantity"`
	}
)

//Market buy params
type (
	MarketBuyParams struct {
		Params
		Quantity float64 `json:"quantity"`
	}
	MarketSellParams struct {
		Params
		Quantity float64 `json:"quantity"`
	}
)

//Trade list response.
type TradeList []struct {
	ID              int    `json:"id"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}

//Deposit history response.
type DepositHistory struct {
	DepositList []struct {
		InsertTime int64   `json:"insertTime"`
		Amount     float64 `json:"amount"`
		Asset      string  `json:"asset"`
		Status     int     `json:"status"`
	} `json:"depositList"`
	Success bool `json:"success"`
}

type WithdrawHistory struct {
	WithdrawList []struct {
		Amount    int    `json:"amount"`
		Address   string `json:"address"`
		Asset     string `json:"asset"`
		ApplyTime int64  `json:"applyTime"`
		Status    int    `json:"status"`
		TxID      string `json:"txId,omitempty"`
	} `json:"withdrawList"`
	Success bool `json:"success"`
}

//Withdraw params
type WithdrawParams struct {
	Asset   string  `json:"asset"`
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
	Name    float64 `json:"name,omitempty"`
}

type WithdrawResponse struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}
