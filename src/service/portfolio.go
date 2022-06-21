package service

import (
	"fmt"

	"github.com/neiltcox/coinbake/database"
	"gorm.io/gorm"
)

// Represents a configured connection to an exchange.
type Portfolio struct {
	gorm.Model
	// TODO: hide field when marshalling to JSON
	ApiKey string
	// TODO: hide field when marshalling to JSON
	ApiSecret          string
	ExchangeIdentifier ExchangeIdentifier
	UserID             int
	User               User
	Name               string
	Connected          bool
	Valuation          float64
}

func FindPortfoliosByUserId(userId uint) []Portfolio {
	portfolios := []Portfolio{}
	database.Handle().Where("user_id = ?", userId).Find(&portfolios)
	return portfolios
}

func FindPortfolioById(portfolioId uint) (*Portfolio, error) {
	portfolio := &Portfolio{}
	result := database.Handle().First(portfolio, portfolioId)
	if result.Error != nil {
		return nil, fmt.Errorf("could not find portfolio: %s", result.Error)
	}

	return portfolio, nil
}
