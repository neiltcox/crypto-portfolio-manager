package service

// A Schedule is a summary of what changes need to happen to a user's account.
type Schedule struct {
	Items []ScheduleItem
}

type ScheduleItem struct {
	// Asset ticker affected.
	Asset string

	// The amount of the asset should change.
	// Positive = buy, negative = sell.
	Amount float32
}
