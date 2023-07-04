package futures

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/nntaoli-project/goex/v2/logger"
	"github.com/nntaoli-project/goex/v2/model"
	"github.com/nntaoli-project/goex/v2/okx/common"
	"github.com/nntaoli-project/goex/v2/options"
	"github.com/nntaoli-project/goex/v2/util"
)

type Swap struct {
	*common.OKxV5
	currencyPairM map[string]model.CurrencyPair
}

func NewSwap() *Swap {
	var currencyPairM = make(map[string]model.CurrencyPair, 64)
	return &Swap{
		OKxV5:         common.New(),
		currencyPairM: currencyPairM}
}

func (f *Swap) GetExchangeInfo() (map[string]model.CurrencyPair, []byte, error) {
	m, b, er := f.OKxV5.GetExchangeInfo("SWAP")
	f.currencyPairM = m
	return m, b, er
}

func (f *Swap) NewCurrencyPair(baseSym, quoteSym string, opts ...model.OptionParameter) (model.CurrencyPair, error) {
	currencyPair := f.currencyPairM[baseSym+quoteSym]
	if currencyPair.Symbol == "" {
		return currencyPair, errors.New("not found currency pair")
	}
	return currencyPair, nil
}

func (f *Swap) NewPrvApi(apiOpts ...options.ApiOption) *PrvApi {
	return NewPrvApi(f.OKxV5, apiOpts...)
}

func (f *PrvApi) ClosePosition(pair model.CurrencyPair, opts ...model.OptionParameter) ([]model.ClosePostion, []byte, error) {
	reqUrl := fmt.Sprintf("%s%s", f.UriOpts.Endpoint, "/v5/trade/close-position")
	params := url.Values{}

	params.Set("instId", pair.Symbol)
	params.Set("mgnMode", "isolated")
	params.Set("autoCxl", "true")
	//params.Set("posSide", "true")

	util.MergeOptionParams(&params, opts...)

	data, responseBody, err := f.DoAuthRequest(http.MethodPost, reqUrl, &params, nil)
	if err != nil {
		logger.Errorf("[ClosePosition] err=%s, response=%s", err.Error(), string(data))
		return nil, responseBody, err
	}

	pos, err := common.UnmarshalClosePositionsResponse(data)
	if err != nil {
		return nil, responseBody, err
	}

	return pos, responseBody, err
}
