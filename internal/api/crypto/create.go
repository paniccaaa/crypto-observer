package api

import (
	"encoding/json"
	"net/http"

	desc "github.com/paniccaaa/crypto-observer/internal/pb"
)

func (i *Implementation) PostCurrencyAdd(w http.ResponseWriter, r *http.Request) {
	var req desc.AddCurrencyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	coin, err := i.cryptoService.Create(req.Coin)
	if err != nil {
		http.Error(w, "failed to add cryptocurrency: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := desc.AddCurrencyResponse{
		Id:        coin.ID,
		Name:      coin.Name,
		CreatedAt: int(coin.CreatedAt.Unix()),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode resp: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
