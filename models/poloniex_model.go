// Appropriate labels for these data are, in order: currencyPair, last, lowestAsk, highestBid, percentChange, baseVolume, quoteVolume, isFrozen, 24hrHigh, 24hrLow

package models

type PoloniexData struct {
	CurrencyPair  string `json:"currencyPair"`
	Last          int    `json:"last"`
	LowestAsk     int    `json:"lowestAsk"`
	HighestBid    int    `json:"highestBid"`
	PercentChange int    `json:"percentChange"`
	BaseVolume    int    `json:"baseVolume"`
	QuoteVolume   int    `json:"quoteVolume"`
	IsFrozen      bool   `json:"isFrozen"`
	DayHigh       int    `json:"24hrHigh"`
	DayLow        int    `json:"24hrLow"`
}
