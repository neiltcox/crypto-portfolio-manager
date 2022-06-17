package service

import (
	"strings"
	"time"

	"github.com/neiltcox/coinbake/database"
	"gorm.io/gorm"
)

// An asset's price difference from the average on a particular exchange must be within this threshold.
// Protects against a bad API call, or an exchange pricing bug.
const AssetPriceSanityCheckDiff float32 = 0.05

// TODO: shorten to something like 2 min in production.
// How long the market data lasts for an asset.
const AssetMarketDataExpiry time.Duration = 15 * time.Minute

type Asset struct {
	gorm.Model
	Symbol    string
	MarketCap uint64
	Volume    uint64

	// An approximate price used for informational purposes only.
	// Do NOT use for creating Schedules.
	// Could also be used to sanity-check exchange prices of assets.
	ApproxPrice float64

	// Last time we've updated market data for this asset.
	LastRefreshed time.Time
}

func FindAssetsByMarketCap(limit int) []Asset {
	assets := []Asset{}
	database.Handle().Order("market_cap DESC, symbol ASC").Limit(limit).Find(&assets)
	return assets
}

func FindAssetsByVolume(limit int) []Asset {
	assets := []Asset{}
	database.Handle().Order("volume DESC, symbol ASC").Limit(limit).Find(&assets)
	return assets
}

func FindAssetBySymbol(symbol string) Asset {
	symbol = strings.ToUpper(symbol)

	assetsWithSymbol := []Asset{}
	database.Handle().Where("symbol = ?", symbol).Find(&assetsWithSymbol)

	if len(assetsWithSymbol) > 0 {
		return assetsWithSymbol[0]
	}

	asset := Asset{
		Symbol: symbol,
	}
	database.Handle().Create(&asset)

	return asset
}

func (asset *Asset) FreshMarketData() bool {
	return time.Since(asset.LastRefreshed) < AssetMarketDataExpiry
}
