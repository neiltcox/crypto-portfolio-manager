package service

import (
	"fmt"
	"log"
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
	tradeBalanceRaw, err := api.Query("TradeBalance", map[string]string{
		"asset": "USD",
	})

	if err != nil {
		return ValidateExchangeConnectionResult{
			Success: false,
			Issue:   fmt.Sprintf("could not query Kraken: %s", err),
		}
	}

	log.Printf("%#v", tradeBalanceRaw)

	return ValidateExchangeConnectionResult{
		Success: true,
		Issue:   "",
	}
}

func (exchangeKraken *ExchangeKraken) SupportedAssets(exchangeConnection *Portfolio) (map[string]bool, error) {
	// TODO: implement
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
