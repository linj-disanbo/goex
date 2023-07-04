package common

import (
	"encoding/json"

	. "github.com/nntaoli-project/goex/v2/options"
)

type OKxV5 struct {
	UriOpts       UriOptions
	UnmarshalOpts UnmarshalerOptions
}

type BaseResp struct {
	Code int             `json:"code,string"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

func New() *OKxV5 {
	unmarshaler := new(RespUnmarshaler)

	f := &OKxV5{
		UriOpts: UriOptions{
			//Endpoint: "https://www.xungcloud.com/priapi",
			//      "https://www.hibuyapp.com", //https://www.okx.com",
			Endpoint:            "https://www.okx.com/api",
			KlineUri:            "/v5/market/candles",
			TickerUri:           "/v5/market/ticker",
			DepthUri:            "/v5/market/books",
			NewOrderUri:         "/v5/trade/order",
			GetOrderUri:         "/v5/trade/order",
			GetPendingOrdersUri: "/v5/trade/orders-pending",
			CancelOrderUri:      "/v5/trade/cancel-order",
			GetAccountUri:       "/v5/account/balance",
			GetPositionsUri:     "/v5/account/positions",
			GetExchangeInfoUri:  "/v5/public/instruments",
		},
		UnmarshalOpts: UnmarshalerOptions{
			ResponseUnmarshaler:                  unmarshaler.UnmarshalResponse,
			KlineUnmarshaler:                     unmarshaler.UnmarshalGetKlineResponse,
			TickerUnmarshaler:                    unmarshaler.UnmarshalTicker,
			DepthUnmarshaler:                     unmarshaler.UnmarshalDepth,
			CreateOrderResponseUnmarshaler:       unmarshaler.UnmarshalCreateOrderResponse,
			GetPendingOrdersResponseUnmarshaler:  unmarshaler.UnmarshalGetPendingOrdersResponse,
			CancelOrderResponseUnmarshaler:       unmarshaler.UnmarshalCancelOrderResponse,
			GetOrderInfoResponseUnmarshaler:      unmarshaler.UnmarshalGetOrderInfoResponse,
			GetAccountResponseUnmarshaler:        unmarshaler.UnmarshalGetAccountResponse,
			GetPositionsResponseUnmarshaler:      unmarshaler.UnmarshalGetPositionsResponse,
			GetFuturesAccountResponseUnmarshaler: unmarshaler.UnmarshalGetFuturesAccountResponse,
			GetExchangeInfoResponseUnmarshaler:   unmarshaler.UnmarshalGetExchangeInfoResponse,
		},
	}

	return f
}

func (okx *OKxV5) WithUriOption(opts ...UriOption) *OKxV5 {
	for _, opt := range opts {
		opt(&okx.UriOpts)
	}
	return okx
}

func (okx *OKxV5) WithUnmarshalOption(opts ...UnmarshalerOption) *OKxV5 {
	for _, opt := range opts {
		opt(&okx.UnmarshalOpts)
	}
	return okx
}

func (okx *OKxV5) NewPrvApi(opts ...ApiOption) *Prv {
	api := NewPrvApi(opts...)
	api.OKxV5 = okx
	return api
}
