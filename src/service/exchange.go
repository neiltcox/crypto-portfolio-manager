package service

import (
	"fmt"

	"gorm.io/gorm"
)

type ExchangeIdentifier string

const (
	ExchangeIdentifierKraken      ExchangeIdentifier = "kraken"
	ExchangeIdentifierCoinbasePro ExchangeIdentifier = "coinbasepro"
	ExchangeIdentifierBinance     ExchangeIdentifier = "binance"
)

var exchanges map[ExchangeIdentifier]Exchange

func init() {
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
	Holdings(*ExchangeConnection) (Holdings, error)
}

type CreatedOrder struct {
	OrderIdentifier string
}

type Holdings struct {
	Holdings []Holding
}

type Holding struct {
	Asset  string
	Amount float32
}

type ExchangeKraken struct {
}

func (exchangeKraken *ExchangeKraken) CreateOrder(exchangeConnection *ExchangeConnection, asset string, amount float32) (CreatedOrder, error) {
	// TODO: implement
	return CreatedOrder{}, nil
}

func (exchangeKraken *ExchangeKraken) Holdings(exchangeConnection *ExchangeConnection) (Holdings, error) {
	// TODO: implement
	return Holdings{}, nil
}
