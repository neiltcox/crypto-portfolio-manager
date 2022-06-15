package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/jaksonkallio/coinbake/config"
)

type Asset struct {
	Ticker    string
	MarketCap uint64
}

var assets map[string]*Asset = make(map[string]*Asset)

// Refreshes all market data.
func RefreshMarketData(cfg config.Config) {
	resp, err := marketDataApi(cfg, "cryptocurrency/listings/latest", url.Values{
		"start":   []string{"1"},
		"limit":   []string{"5000"},
		"convert": []string{"USD"},
	})
	if err != nil {
		log.Printf("Could not reach market data API: %s", err)
	}

	log.Println(string(resp))
}

func marketDataApi(cfg config.Config, endpoint string, query url.Values) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", cfg.MarketData.BaseUrl, endpoint), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", cfg.MarketData.ApiKey)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
