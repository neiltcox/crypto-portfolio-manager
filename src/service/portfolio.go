package service

import (
	"fmt"
	"log"
	"time"

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
	TotalValuation     float64
	LastRefresh        time.Time
}

// TODO: add pay-tiers to refresh interval.
// How often to refresh the portfolio.
const PortfolioRefreshInterval time.Duration = 8 * time.Hour

// How many portfolios to refresh per batch.
const PortfolioRefreshBatchCount int = 3

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

func refreshStalePortfolios() {
	stalePortfolios := []Portfolio{}
	database.Handle().Where("last_refresh <= ?", time.Time.Add(time.Now(), -PortfolioRefreshInterval)).Limit(PortfolioRefreshBatchCount).Find(&stalePortfolios)

	if len(stalePortfolios) > 0 {
		log.Printf("Refreshing %d portfolios", len(stalePortfolios))
	}

	for _, stalePortfolio := range stalePortfolios {
		stalePortfolio.Refresh()
	}
}

// Refreshes valuation information for the portfolio.
func (portfolio *Portfolio) Refresh() error {
	exchange, err := portfolio.Exchange()
	if err != nil {
		portfolio.MarkDisconnected()
		return err
	}

	holdingSummary, err := exchange.HoldingSummary(portfolio)
	if err != nil {
		portfolio.MarkDisconnected()
		return err
	}

	portfolio.TotalValuation = holdingSummary.TotalBalanceValuation
	portfolio.Connected = true
	portfolio.LastRefresh = time.Now()

	database.Handle().Save(portfolio)

	return nil
}

func (portfolio *Portfolio) MarkDisconnected() {
	database.Handle().Model(&portfolio).Update("connected", false)
}

func PortfolioRefresher(ticks *time.Ticker, stop chan bool) {
	for {
		select {
		case <-stop:
			return
		case <-ticks.C:
			refreshStalePortfolios()
		}
	}
}
