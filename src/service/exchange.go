package service

import (
	"fmt"
)

type ExchangeIdentifier string

const (
	ExchangeIdentifierMocked      ExchangeIdentifier = "mock"
	ExchangeIdentifierKraken      ExchangeIdentifier = "kraken"
	ExchangeIdentifierCoinbasePro ExchangeIdentifier = "coinbasepro"
	ExchangeIdentifierBinance     ExchangeIdentifier = "binance"
)

var exchanges map[ExchangeIdentifier]Exchange = make(map[ExchangeIdentifier]Exchange)

func init() {
	exchanges[ExchangeIdentifierMocked] = &ExchangeMocked{}
	exchanges[ExchangeIdentifierKraken] = &ExchangeKraken{}
}

type SupportedAsset struct {
	Asset Asset
}

// An interface representing a generic exchange.
type Exchange interface {
	CreateOrder(*Portfolio, string, float32) (CreatedOrder, error)
	Holdings(*Portfolio) (map[string]Holding, error)
	SupportedAssets(*Portfolio) (map[string]bool, error)
	SupportsAsset(*Portfolio, Asset) bool
	ValidateConnection(*Portfolio) ValidateExchangeConnectionResult
	HoldingSummary(*Portfolio) (HoldingSummary, error)
}

type ValidateExchangeConnectionResult struct {
	// Whether we're able to connect to the exchange at all.
	Success bool

	// Any error message that we may want to provide to the user about the connection.
	Issue string
}

type MockSupportedAssets struct {
}

type CreatedOrder struct {
	OrderIdentifier string
}

type Holding struct {
	Asset   Asset
	Balance float64
}

type HoldingSummary struct {
	// How much the user holds in total.
	TotalBalanceValuation float64
}

type ExchangeMocked struct {
	MockSupportedAssets
}

func (mockSupportedAssets *MockSupportedAssets) SupportedAssets(portfolio *Portfolio) (map[string]bool, error) {
	return map[string]bool{
		"BTC": true,
		"ETH": true,
		"XMR": true,
	}, nil
}

// Gets the Exchange object for a given Exchange Connection, which is where the API call logic is.
func (portfolio *Portfolio) Exchange() (Exchange, error) {
	exchange, exists := exchanges[portfolio.ExchangeIdentifier]
	if !exists {
		return nil, fmt.Errorf("exchange %q is not implemented", portfolio.ExchangeIdentifier)
	}

	return exchange, nil
}
