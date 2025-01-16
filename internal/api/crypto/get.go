package api

import (
	"encoding/json"
	"net/http"
	"time"

	desc "github.com/paniccaaa/crypto-observer/internal/pb"
)

func (i *Implementation) GetCurrencyPrice(w http.ResponseWriter, r *http.Request, params desc.GetCurrencyPriceParams) {
	coinPrice, err := i.cryptoService.Get(params.Coin, time.Unix(int64(params.Timestamp), 0))
	if err != nil {
		http.Error(w, "invalid to get coin price: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := desc.PriceResponse{
		Id:        coinPrice.ID,
		Coin:      coinPrice.Name,
		Price:     float32(coinPrice.Price),
		Timestamp: params.Timestamp,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode resp: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
