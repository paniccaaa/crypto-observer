package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetCoinPrice(cryptocurrency string) (float64, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", cryptocurrency)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}()

	var price map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&price); err != nil {
		return 0, err
	}

	return price[cryptocurrency]["usd"], nil
}
