package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CmcArray struct {
	LastUpdate time.Time
	Data       []CmcObj
}
type CmcObj struct {
	ID                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Symbol            string `json:"symbol,omitempty"`
	Rank              string `json:"rank,omitempty"`
	USDPrice          string `json:"price_usd,omitempty"`
	BTCPrice          string `json:"price_btc,omitempty"`
	DayVolume         string `json:"24h_volume_usd,omitempty"`
	MarketCapUSD      string `json:"market_cap_usd,omitempty"`
	AvaliableSupply   string `json:"available_supply,omitempty"`
	TotalSupply       string `json:"total_supply,omitempty"`
	PercentChangeHour string `json:"percent_change_1h,omitempty"`
	PercentChangeDay  string `json:"percent_change_24h,omitempty"`
	PercentChangeWeek string `json:"percent_change_7d,omitempty"`
	LastUpdated       string `json:"last_updated,omitempty"`
}

// func CallCMC() []CmcObj {
func CallCMC(t time.Time) CmcArray {
	rs, err := http.Get("https://api.coinmarketcap.com/v1/ticker/?limit=200")

	if err != nil {
		fmt.Println(err)
		log.Fatal("Error Connecting to CoinMarketCap\n", t)
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	result := CmcArray{}
	var prices []CmcObj

	parseErr := json.Unmarshal(bodyBytes, &prices)

	if parseErr != nil {
		fmt.Println(parseErr)
		log.Fatal(parseErr)
	}

	result.LastUpdate = t
	result.Data = prices

	return result
}
