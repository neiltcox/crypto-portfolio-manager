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

// Only submit orders with a change in amount of at least this much.
// For example, if an asset balance should change by less than epsilon value 0.001 (0.1%), simply ignore it.
const DustEpsilonRate float64 = 0.0001
const MaxStrategyAssetCount int = 100

// Generates a schedule utilizing this strategy and the user's current holdings.
func (strategy *Strategy) GenerateSchedule(portfolio *Portfolio) (Schedule, error) {
	exchange, err := portfolio.Exchange()
	if err != nil {
		return Schedule{}, fmt.Errorf("could not get exchange from exchange connection: %s", err)
	}

	holdings, err := exchange.Holdings(portfolio)
	if err != nil {
		return Schedule{}, fmt.Errorf("could not fetch holdings: %s", err)
	}

	supportedAssets, err := exchange.SupportedAssets(portfolio)
	if err != nil {
		return Schedule{}, fmt.Errorf("could not get supported exchange assets: %s", err)
	}

	// Which assets to use
	var eligibleAssets []Asset
	assets := make([]Asset, 0)

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
			return Schedule{}, fmt.Errorf("weighting metric %q not yet implemented", strategy.WeightingMetric)
		}
	} else {
		return Schedule{}, fmt.Errorf("asset selection method not yet implemented")
	}

	// At this point we have a sorted list of assets that should be included.

	// Prepare the resulting schedule struct.
	schedule := Schedule{
		Items:             make([]ScheduleItem, 0),
		UnsupportedAssets: make([]Asset, 0),
	}

	// Build the list of unsupported assets.
	// We must remove the unsupported assets from the weight calculations.
	// This is because an unsupported asset should not effectively take up any weight.
	for _, asset := range eligibleAssets {
		supported, exists := supportedAssets[asset.Symbol]
		if !exists || !supported {
			// Not a supported asset.
			// Add the asset to the schedule's list of unsupported assets.
			schedule.UnsupportedAssets = append(schedule.UnsupportedAssets, asset)
		} else {
			// Add the eligible asset to the list of assets that we'll include in the calculations.
			assets = append(assets, asset)
		}
	}

	// Total weight of the metric we're weighing by.
	var totalWeight uint64
	for _, asset := range assets {
		var weight uint64

		switch strategy.WeightingMetric {
		case WeightingMetricMarketCap:
			weight = asset.MarketCap
		case WeightingMetricVolume:
			weight = asset.Volume
		}

		totalWeight += weightAfterModifier(weight, strategy.WeightingModifier)
	}

	return schedule, nil
}

func weightAfterModifier(weight uint64, modifier WeightingModifier) uint64 {
	switch modifier {
	case WeightingModifierNone:
		return weight
	case WeightingModifierSquareRoot:
		return uint64(math.Pow(float64(weight), 0.5))
	case WeightingModifierCubeRoot:
		return uint64(math.Pow(float64(weight), 0.33333))
	default:
		return 1
	}
}

func FindStrategyByExchangeConnectionId(exchangeConnectionId uint) *Strategy {
	strategy := Strategy{}
	database.Handle().Where("exchange_connection_id = ?", exchangeConnectionId).First(&strategy)
	return &strategy
}
