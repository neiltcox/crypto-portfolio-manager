package service

import "gorm.io/gorm"

type WeightingMetric string
type WeightingModifier string

const (
	WeightingMetricEvenly       WeightingMetric   = "even"
	WeightingMetricMarketCap    WeightingMetric   = "mcap"
	WeightingMetricVolume       WeightingMetric   = "vol"
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

	// Work with an exact set of assets (comma-separated tickers).
	ExactAssets string

	// Which metric to use for weighting.
	WeightingMetric

	// Which mathematic function should be used to modify the weighting metric.
	WeightingModifier

	UserID int
	User   User
}

// Generates a schedule utilizing this strategy and the user's current holdings.
func (strategy *Strategy) GenerateSchedule() {

}
