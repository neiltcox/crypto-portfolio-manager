package service

import (
	"fmt"
	"math"

	"github.com/neiltcox/coinbake/database"
	"gorm.io/gorm"
)

type WeightingMetric string
type WeightingModifier string

const (
	WeightingMetricMarketCap    WeightingMetric   = "mcap"
	WeightingMetricVolume       WeightingMetric   = "vol"
	WeightingModifierEven       WeightingModifier = "even"
	WeightingModifierNone       WeightingModifier = "none"
	WeightingModifierSquareRoot WeightingModifier = "sqrt"
	WeightingModifierCubeRoot   WeightingModifier = "cbrt"
)

// Determines how a Schedule is created for a user's account.
// A strategy could be "evenly distribute among top N coins" or
type Strategy struct {
	gorm.Model

	// Work with the top N assets.
	TopAssetCount int

	// TODO: ignore-list of assets

	// Which metric to use for weighting.
	WeightingMetric

	// Which mathematic function should be used to modify the weighting metric.
	WeightingModifier

	// Which exchange this strategy is associated to.
	PortfolioID int
	Portfolio   Portfolio
}

// A single line item in the strategy calculation.
type RebalanceMovement struct {
	Asset   Asset
	Balance float64
	// The absolute weight value given to this line item.
	WeightValue float64
	// The target weight proportion that this line item should have.
	WeightProportion float64
	TargetValuation  float64
	ValuationDiff    float64
}

type RebalanceMovementSummary struct {
	Movements         []*RebalanceMovement
	UnsupportedAssets []Asset
}

// Only submit orders with a change in amount of at least this much.
// For example, if an asset balance should change by less than epsilon value 0.001 (0.1%), simply ignore it.
const DustEpsilonRate float64 = 0.001

// Maximum number of assets that can be included in a strategy.
const MaxStrategyAssetCount int = 500

func (rebalanceMovement *RebalanceMovement) Valuation() float64 {
	return rebalanceMovement.Balance * rebalanceMovement.Asset.ApproxPrice
}

// Generates a schedule utilizing this strategy and the user's current holdings.
func (strategy *Strategy) RebalanceMovements(portfolio *Portfolio) (RebalanceMovementSummary, error) {
	exchange, err := portfolio.Exchange()
	if err != nil {
		return RebalanceMovementSummary{}, fmt.Errorf("could not get exchange from exchange connection: %s", err)
	}

	// Get the assets we should do calculations with.
	// Also, get the unsupported assets list so that we can communicate to the user that these assets could not be included.
	assets, unsupportedAssets, err := strategy.considerableAssets(exchange, portfolio)
	if err != nil {
		return RebalanceMovementSummary{}, err
	}

	// Get the map of asset symbol -> holding.
	holdings, err := exchange.Holdings(portfolio)
	if err != nil {
		return RebalanceMovementSummary{}, fmt.Errorf("could not fetch holdings: %s", err)
	}

	// Prepare our calculation line items
	rebalanceMovements := make([]*RebalanceMovement, 0)

	for _, asset := range assets {
		// Find the balance, defaulting to zero if there is no balance.
		balance := 0.0
		holding, isHeld := holdings[asset.Symbol]
		if isHeld {
			balance = holding.Balance
		}

		rebalanceMovements = append(
			rebalanceMovements,
			&RebalanceMovement{
				Asset:       asset,
				Balance:     balance,
				WeightValue: calculateWeight(asset, strategy.WeightingMetric, strategy.WeightingModifier),
			},
		)
	}

	// Find our totals for the calc lines
	var totalWeight float64
	var totalValuation float64
	for _, rebalanceMovement := range rebalanceMovements {
		totalWeight += rebalanceMovement.WeightValue
		totalValuation += rebalanceMovement.Valuation()
	}

	for _, rebalanceMovement := range rebalanceMovements {
		rebalanceMovement.WeightProportion = rebalanceMovement.WeightValue / totalWeight
		rebalanceMovement.TargetValuation = rebalanceMovement.WeightProportion * totalValuation
		rebalanceMovement.ValuationDiff = rebalanceMovement.TargetValuation - rebalanceMovement.Valuation()
	}

	return RebalanceMovementSummary{
		Movements:         rebalanceMovements,
		UnsupportedAssets: unsupportedAssets,
	}, nil
}

// Which assets should be considered for calculations for a given strategy.
// Filters out assets unsupported by the exchange.
// First return are considerable assets, second return are unsupported assets.
func (strategy *Strategy) considerableAssets(exchange Exchange, portfolio *Portfolio) ([]Asset, []Asset, error) {
	// Which assets to choose from.
	var eligibleAssets []Asset

	// Respect our maximum strategy asset count
	assetCount := strategy.TopAssetCount
	if assetCount > MaxStrategyAssetCount {
		assetCount = MaxStrategyAssetCount
	}

	if assetCount > 0 {
		// Consider all assets
		if strategy.WeightingMetric == WeightingMetricMarketCap {
			eligibleAssets = FindAssetsByMarketCap(assetCount)
		} else if strategy.WeightingMetric == WeightingMetricVolume {
			eligibleAssets = FindAssetsByVolume(assetCount)
		} else {
			return nil, nil, fmt.Errorf("weighting metric %q not yet implemented", strategy.WeightingMetric)
		}
	} else {
		return nil, nil, fmt.Errorf("asset selection method not yet implemented")
	}

	// Initialize our resulting asset lists.
	considerableAssets := make([]Asset, 0)
	unsupportedAssets := make([]Asset, 0)

	// Build the list of unsupported assets.
	// We must remove the unsupported assets from the weight calculations.
	// This is because an unsupported asset should not effectively take up any weight.
	for _, asset := range eligibleAssets {
		if exchange.SupportsAsset(portfolio, asset) {
			// Add the eligible asset to the list of assets that we'll include in the calculations.
			considerableAssets = append(considerableAssets, asset)
		} else {
			// Not a supported asset.
			// Add the asset to the schedule's list of unsupported assets.
			unsupportedAssets = append(unsupportedAssets, asset)
		}
	}

	return considerableAssets, unsupportedAssets, nil
}

func calculateWeight(asset Asset, metric WeightingMetric, modifier WeightingModifier) float64 {
	var weight float64

	// Figure the weight value based on the metric.
	// Dividing by 1,000 because weights should be expressed in thousands of dollars.
	switch metric {
	case WeightingMetricMarketCap:
		//weight = float64(asset.MarketCap) / 1000.0
		weight = float64(asset.MarketCap)
	case WeightingMetricVolume:
		//weight = float64(asset.Volume) / 1000.0
		weight = float64(asset.Volume)
	}

	switch modifier {
	case WeightingModifierNone:
		return weight
	case WeightingModifierSquareRoot:
		return math.Pow(weight, 0.5)
	case WeightingModifierCubeRoot:
		return math.Pow(weight, 0.33333)
	default:
		return 1
	}
}

func FindStrategyByPortfolioId(exchangeConnectionId uint) *Strategy {
	strategy := Strategy{}
	database.Handle().Where("portfolio_id = ?", exchangeConnectionId).First(&strategy)
	return &strategy
}
