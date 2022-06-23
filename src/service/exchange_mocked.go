package service

func (exchangeMocked *ExchangeMocked) CreateOrder(portfolio *Portfolio, asset string, amount float32) (CreatedOrder, error) {
	return CreatedOrder{
		OrderIdentifier: "123456",
	}, nil
}

func (exchangeMocked *ExchangeMocked) Holdings(portfolio *Portfolio) (map[string]Holding, error) {
	return map[string]Holding{
		"BTC": {Asset: FindAssetBySymbol("BTC"), Balance: 0.23},
		"ETH": {Asset: FindAssetBySymbol("ETH"), Balance: 2.3},
		"XMR": {Asset: FindAssetBySymbol("XMR"), Balance: 43.145},
		"BNB": {Asset: FindAssetBySymbol("BNB"), Balance: 0.033},
		"ADA": {Asset: FindAssetBySymbol("ADA"), Balance: 50.2},
	}, nil
}

func (exchangeMocked *ExchangeMocked) SupportsAsset(portfolio *Portfolio, asset Asset) bool {
	return true
}

func (exchangeMocked *ExchangeMocked) ValidateConnection(portfolio *Portfolio) ValidateExchangeConnectionResult {
	return ValidateExchangeConnectionResult{
		Success: true,
		Issue:   "",
	}
}

func (exchangeMocked *ExchangeMocked) HoldingSummary(portfolio *Portfolio) (HoldingSummary, error) {
	return HoldingSummary{
		TotalBalanceValuation: 1234.56,
	}, nil
}
