package service

import (
	"fmt"

	"gorm.io/gorm"
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

// Represents a configured connection to an exchange.
type ExchangeConnection struct {
	gorm.Model
	ApiKey    string
	ApiSecret string
	ExchangeIdentifier
	UserID int
	User   User
}

// Gets the Exchange object for a given Exchange Connection, which is where the API call logic is.
func (exchangeConnection *ExchangeConnection) Exchange() (Exchange, error) {
	exchange, exists := exchanges[exchangeConnection.ExchangeIdentifier]
	if !exists {
		return nil, fmt.Errorf("exchange %q is not implemented", exchangeConnection.ExchangeIdentifier)
	}

	return exchange, nil
}

// An interface representing a generic exchange.
type Exchange interface {
	CreateOrder(*ExchangeConnection, string, float32) (CreatedOrder, error)
	Holdings(*ExchangeConnection) ([]Holding, error)
}

type CreatedOrder struct {
	OrderIdentifier string
}

type Holding struct {
	Asset   string
	Balance float32
}

type ExchangeKraken struct {
}

type ExchangeMocked struct {
}

func (exchangeMocked *ExchangeMocked) CreateOrder(exchangeConnection *ExchangeConnection, asset string, amount float32) (CreatedOrder, error) {
	return CreatedOrder{
		OrderIdentifier: "123456",
	}, nil
}

func (exchangeMocked *ExchangeMocked) Holdings(exchangeConnection *ExchangeConnection) ([]Holding, error) {
	return []Holding{
		{Asset: "BTC", Balance: 0.23},
		{Asset: "ETH", Balance: 2.3},
		{Asset: "XMR", Balance: 43.145},
		{Asset: "ADA", Balance: 0.033},
	}, nil
}

func (exchangeKraken *ExchangeKraken) CreateOrder(exchangeConnection *ExchangeConnection, asset string, amount float32) (CreatedOrder, error) {
	return CreatedOrder{
		OrderIdentifier: "123456",
	}, nil
}

func (exchangeKraken *ExchangeKraken) Holdings(exchangeConnection *ExchangeConnection) ([]Holding, error) {
	return []Holding{}, nil
}
