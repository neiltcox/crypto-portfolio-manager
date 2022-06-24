package service

import (
	"fmt"
	"strconv"

	krakenapi "github.com/beldur/kraken-go-api-client"
)

type ExchangeKraken struct {
}

func (exchangeKraken *ExchangeKraken) CreateOrder(exchangeConnection *Portfolio, asset string, amount float32) (CreatedOrder, error) {
	return CreatedOrder{
		OrderIdentifier: "123456",
	}, nil
}

func (exchangeKraken *ExchangeKraken) Holdings(exchangeConnection *Portfolio) (map[string]Holding, error) {
	return map[string]Holding{}, nil
}

func (exchangeKraken *ExchangeKraken) SupportsAsset(exchangeConnection *Portfolio, asset Asset) bool {
	return true
}

func (exchangeKraken *ExchangeKraken) ValidateConnection(portfolio *Portfolio) ValidateExchangeConnectionResult {
	api := krakenapi.New(portfolio.ApiKey, portfolio.ApiSecret)
	_, err := api.Query("TradeBalance", map[string]string{
		"asset": "USD",
	})

	if err != nil {
		return ValidateExchangeConnectionResult{
			Success: false,
			Issue:   fmt.Sprintf("could not query Kraken: %s", err),
		}
	}

	return ValidateExchangeConnectionResult{
		Success: true,
		Issue:   "",
	}
}

func (exchangeKraken *ExchangeKraken) SupportedAssets(portfolio *Portfolio) (map[string]bool, error) {
	api := krakenapi.New(portfolio.ApiKey, portfolio.ApiSecret)
	_, err := api.Query("Assets", map[string]string{})
	if err != nil {
		return make(map[string]bool, 0), err
	}

	/*
		assetsResponse := assetsResponseRaw.([]map[string]interface{})
		supportedAssets := make(map[string]bool, 0)

		for _, asset := range assetsResponse {
			//supportedAssets[asset["symbol"]]
		}*/

	return map[string]bool{}, nil
}

func (exchangeKraken *ExchangeKraken) HoldingSummary(portfolio *Portfolio) (HoldingSummary, error) {
	api := krakenapi.New(portfolio.ApiKey, portfolio.ApiSecret)
	tradeBalancesResponseRaw, err := api.Query("TradeBalance", map[string]string{
		"asset": "USD",
	})
	if err != nil {
		return HoldingSummary{}, err
	}

	tradeBalanceResponse := tradeBalancesResponseRaw.(map[string]interface{})

	totalBalance, err := strconv.ParseFloat(tradeBalanceResponse["eb"].(string), 64)

	return HoldingSummary{
		TotalBalanceValuation: totalBalance,
	}, nil
}
